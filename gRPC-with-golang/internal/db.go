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
			return fmt.Errorf("duplicate order id: %d", order.GetOrderId())
		}
	}
	db.collection = append(db.collection, order)
	return nil
}

func (db *DB) GetOrder(orderId uint64) (*orders.PayloadWithSingleOrder, error) {
	for _, o := range db.collection {
		if o.GetOrderId() == orderId {
			return &orders.PayloadWithSingleOrder{Order: o}, nil
		}
	}
	return nil, fmt.Errorf("order not found")
}

func (db *DB) UpdateOrder(order *orders.Order) bool {
	for i, o := range db.collection {
		if o.GetOrderId() == order.GetOrderId() {
			db.collection[i] = order
			return true
		}
	}
	return false
}

func (db *DB) RemoveOrder(orderId uint64) {
	for i, o := range db.collection {
		if o.GetOrderId() == orderId {
			db.collection = append(db.collection[:i], db.collection[i+1:]...)
		}
	}
}
