package internal

import (
	"context"
	"errors"
	"gRPC-with-golang/protogen/golang/orders"
	"log"
)

type OrderService struct {
	db *DB
	orders.UnimplementedOrdersServer
}

func NewOrderService(db *DB) OrderService {
	return OrderService{db: db}
}

func (o *OrderService) AddOrder(_ context.Context, req *orders.PayloadWithSingleOrder) (*orders.Empty, error) {
	log.Printf("Adding order with payload: %v", req)

	err := o.db.AddOrder(req.GetOrder())

	return &orders.Empty{}, err
}

func (o *OrderService) GetOrder(_ context.Context, req *orders.PayloadWithOrderId) (*orders.PayloadWithSingleOrder, error) {
	log.Printf("Getting order with payload: %v", req)
	return o.db.GetOrder(req.GetOrderId())
}

func (o *OrderService) UpdateOrder(_ context.Context, req *orders.PayloadWithSingleOrder) (*orders.Empty, error) {
	log.Printf("Updating order with payload: %v", req)
	isUpdated := o.db.UpdateOrder(req.GetOrder())
	if isUpdated {
		return &orders.Empty{}, nil
	}
	return &orders.Empty{}, errors.New("order not found")
}

func (o *OrderService) RemoveOrder(_ context.Context, req *orders.PayloadWithOrderId) (*orders.Empty, error) {
	log.Printf("Deleting order with id: %d", req)
	o.db.RemoveOrder(req.GetOrderId())
	return &orders.Empty{}, nil
}
