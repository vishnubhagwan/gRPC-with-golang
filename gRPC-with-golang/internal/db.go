package internal

import (
	"fmt"
	"gRPC-with-golang/protogen/golang/orders"
)

type DB struct {
	collection []*orders.Order
}

func NewDB() *DB {
	return &DB{collection: make([]*orders.Order, 0)}
}

func (db *DB) AddOrder(order *orders.Order) error {
	for _, o := range db.collection {
		if o.GetOrderId() == order.GetOrderId() {
			return fmt.Errorf("order already exists with order id %d", order.GetOrderId())
		}
	}
	db.collection = append(db.collection, order)
	return nil
}
