package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

var (
	RepoErr             = errors.New("Unable to handle Repo Request")
	ErrIdNotFound       = errors.New("Id not found")
	ErrPhonenumNotFound = errors.New("Order num is not found")
)

type repo struct {
	db *sql.DB
	logger log.Logger
}

func newRepo(db *sql.DB, logger log.Logger) (Repository, error)  {
	return &repo{
		db: db,
		logger: log.With(logger,"repo","mysql"),
	}, nil
}

func (repo *repo) CreateOrder(ctx context.Context, order Order) error  {
	_, err := repo.db.ExecContext(ctx,
		"insert into ordergo(status, total) values(?,?)", order.Status,order.Total)
	if err!=nil {
		fmt.Println("Error al insertar ")
		return err
	}else {
		fmt.Println("order creada")
	}
	return nil
}

func (repo *repo) GetOrder(ctx context.Context,id int)(interface{},error)  {
	order := Order{}

	err := repo.db.QueryRowContext(ctx,
		"select id, status, total from orderdb.order as o where o.id = ?",id).Scan(&order.Id,&order.Status,&order.Total)

    if err != nil {
		if err != sql.ErrNoRows {
			return order, ErrIdNotFound

		}
		return order, err
	}
	return order, err

}

func (repo *repo) GetAllOrder(ctx context.Context) (interface{},error)  {

	logger := log.With(repo.logger,"method", "GetAllOrder")
	logger.Log("entra al repo a consultar getallorders")
	order := Order{}
	var res []interface{}

	rows , err := repo.db.QueryContext(ctx,"select id, status, total from ordergo")

	if err != nil {
		level.Error(logger).Log("error select all",err)
		if err != sql.ErrNoRows {

			level.Error(logger).Log("error select Errnorows",err)
			return order, ErrIdNotFound

		}
		return order, err
	}

	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&order.Id,&order.Status,&order.Total)
		res = append([]interface{}{order},res...)

	}
	return res,nil

}

func (repo *repo) UpdateOrder(ctx context.Context, order Order) (int,error)  {
	res, err := repo.db.ExecContext(ctx,
		"update orderdb.order as o set o.status , o.total where o.id = ?", order.Status,order.Total,order.Id)
	if err!=nil {
		fmt.Println("Error al actualizar ")
		return 0,err
	}

	rowCnt, err := res.RowsAffected()
	if err!=nil {
		return 0,err
	}
	if rowCnt == 0 {
		return 0,ErrIdNotFound
	}

	return 1,err
}

func (repo *repo) DeleteOrder(ctx context.Context,id int)(int, error)  {

	res, err := repo.db.ExecContext(ctx,
		"delete from orderdb.order as o  where o.id = ?", id)
	if err!=nil {
		fmt.Println("Error al eliminar ")
		return 0,err
	}

	rowCnt, err := res.RowsAffected()
	if err!=nil {
		return 0,err
	}
	if rowCnt == 0 {
		return 0,ErrIdNotFound
	}

	return 1,err
}

