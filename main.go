package main

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main()  {
	logger:=log.NewLogfmtLogger(os.Stderr)
	db := dbConn()

	r := mux.NewRouter()

	var svc OrderService
	svc = orderService{}
	{
		repository,err := newRepo(db,logger)
		if err != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(-1)
		}
		svc = NewService(repository,logger)
	}

	CreateOrderHandler := httptransport.NewServer(
		makeCreateOrderEndpoint(svc),
		decodeCreateOrderRequest,
		encodeResponse,
		)

	GetAllOrderHandler := httptransport.NewServer(
        makeGetAllOrderEndpoint(svc),
        decodeGetAllOrderRquest,
        encodeResponse,
		)

	http.Handle("/",r)
	http.Handle("/order", CreateOrderHandler)
	http.Handle("/orders", GetAllOrderHandler)

	logger.Log("msg", "HTTP", "addr", ":8001")
	logger.Log("err", http.ListenAndServe(":8001", nil))



}