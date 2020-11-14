package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/go-kit/kit/endpoint"

)

type (
	CreateOrderRequest struct {
		order Order
	}

	CreateOrderResponse struct {
		Code int `json:"code"`
		Err error `json:"error,omitempty"`
	}

	GetOrderByIdRequest struct {
		Id int `json:"orderid"`
	}

	GetOrderByIdResponse struct {
		Order interface{} `json:"order,omitempty"`
		Err      string      `json:"error,omitempty"`
	}
	GetAllOrdersRequest struct{}

	GetAllOrdersResponse struct {
		Order interface{} `json:"order,omitempty"`
		Err      string    `json:"error,omitempty"`
	}
	DeleteOrderRequest struct {
		Orderid int `json:"orderid"`
	}

	DeleteOrderResponse struct {
		Msg int `json:"response"`
		Err error  `json:"error,omitempty"`
	}
	UpdateOrderRequest struct {
		order Order
	}
	UpdateOrderResponse struct {
		Msg string `json:"status,omitempty"`
		Err error  `json:"error,omitempty"`
	}


)


func makeCreateOrderEndpoint(s OrderService) endpoint.Endpoint  {
	return func(ctx context.Context, request interface{}) (interface{}, error ) {
		req := request.(CreateOrderRequest)
		msg, err := s.CreateOrder(ctx,req.order)
		return CreateOrderResponse{Code: msg,Err: err},nil
	}
}

func makeGetOrderByIdEndpoint(s OrderService) endpoint.Endpoint  {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetOrderByIdRequest)
		orderOut, err := s.GetOrder(ctx,req.Id)
		if err!=nil {
			return GetOrderByIdResponse{Order: orderOut,Err: "error id not found"},nil
		}
		return GetOrderByIdResponse{Order: orderOut,Err: ""},nil

	}
}

func makeGetAllOrderEndpoint(s OrderService) endpoint.Endpoint  {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		orderOut, err := s.GetAllOrder(ctx)
		if err != nil {
			return GetAllOrdersResponse{Order: orderOut,Err: "no data found in get all"}, nil

		}
		return GetAllOrdersResponse{Order: orderOut,Err: ""},nil
	}
}

func makeDeleteOrderEndpoint(s OrderService) endpoint.Endpoint  {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(DeleteOrderRequest)
		msg, err := s.DeleteOrder(ctx,req.Orderid)
		if err!=nil {
			return DeleteOrderResponse{Msg: msg,Err: err},nil
		}
		return DeleteOrderResponse{Msg: msg,Err: nil},nil
	}

}

func decodeGetAllOrderRquest(_ context.Context, r *http.Request) (interface{},error) {
	fmt.Println("------------>>> into decoding get all")
	return r,nil
}

func decodeCreateOrderRequest(_ context.Context, r *http.Request) (interface{},error)  {
	var req CreateOrderRequest
	fmt.Println("------------>>> into decoding create")
	if err := json.NewDecoder(r.Body).Decode(&req.order); err != nil {
		fmt.Println("------------>>> error",err)
		return nil, err
	}
	fmt.Println("------------>>> request ",req)
	return req,nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Println("into Encoding <<<<<<----------------")
	return json.NewEncoder(w).Encode(response)
}