package main

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type Order struct {

	Id int `json:"id"`
	Status string `json:"status"`
	Total int `json:"total"`

}

type Repository interface {

	CreateOrder(ctx context.Context, order Order) error
	GetOrder(ctx context.Context,id int)(interface{},error)
	GetAllOrder(ctx context.Context) (interface{},error)
	UpdateOrder(ctx context.Context, order Order) (int,error)
	DeleteOrder(ctx context.Context,id int)(int, error)
}

type orderService struct {
	repository Repository
	logger log.Logger
}

type OrderService interface {
	CreateOrder(ctx context.Context, order Order) (int,error)
	GetOrder(ctx context.Context,id int)(interface{},error)
	GetAllOrder(ctx context.Context) (interface{},error)
	UpdateOrder(ctx context.Context, order Order) (int,error)
	DeleteOrder(ctx context.Context,id int)(int, error)
}

func (s orderService) CreateOrder(ctx context.Context, order Order) (int,error)  {
	logger := log.With(s.logger,"method", "Create")
	var msg = 1
	orderIn := Order{
		Id: order.Id,
		Status: order.Status,
		Total:  order.Total,
	}
	if err := s.repository.CreateOrder(ctx,orderIn); err != nil {
		level.Error(logger).Log("error from repo ", err)
		return 0, err
	}
	return msg, nil

}

func (s orderService) GetOrder(ctx context.Context,id int)(interface{},error)  {
	logger := log.With(s.logger,"method", "GetOrder")
	var order interface{}
	var empty interface{}
	order, err := s.repository.GetOrder(ctx,id)
	if err != nil {
		level.Error(logger).Log("error",err)
		return empty,err
	}
	return order,nil
}

func (s orderService) GetAllOrder(ctx context.Context)(interface{},error)  {
	logger := log.With(s.logger,"method", "GetAllOrder")
	var order interface{}
	var empty interface{}
	order, err := s.repository.GetAllOrder(ctx)
	if err != nil {
		level.Error(logger).Log("error en bd get all",err)
		return empty,err
	}
	return order,nil
}

func (s orderService) UpdateOrder(ctx context.Context,order Order)(int,error)  {
	logger := log.With(s.logger,"method", "GetOrder")
	var msg = 1
	orderIn := Order{
		Id: order.Id,
		Status: order.Status,
		Total:  order.Total,
	}
	msg, err := s.repository.UpdateOrder(ctx,orderIn)

	if  err != nil {
		level.Error(logger).Log("error from repo ", err)
		return 0, err
	}
	return msg, nil
}

func (s orderService) DeleteOrder(ctx context.Context,id int)(int,error)  {
	logger := log.With(s.logger,"method", "GetOrder")
	var msg = 1

	msg, err := s.repository.DeleteOrder(ctx,id)

	if  err != nil {
		level.Error(logger).Log("error from repo ", err)
		return 0, err
	}
	return msg, nil
}

func NewService(rep Repository, logger log.Logger) OrderService  {
	return &orderService{
		repository: rep,
		logger: logger,
	}
}