package internal

import (
	"context"
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
func (o *OrderService) AddOrder(ctx context.Context, req *orders.PayloadWithSingleOrder) (*orders.Empty, error) {
	log.Printf("Adding order with payload: %v", req)

	err := o.db.AddOrder(req.GetOrder())

	return &orders.Empty{}, err
}
