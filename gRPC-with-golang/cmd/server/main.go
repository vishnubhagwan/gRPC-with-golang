package main

import (
	"gRPC-with-golang/internal"
	"gRPC-with-golang/protogen/golang/orders"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	const addr = "0.0.0.0:50051"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on %s", addr)
	}
	server := grpc.NewServer()
	db := internal.NewDB()
	orderService := internal.NewOrderService(db)
	orders.RegisterOrdersServer(server, &orderService)
	log.Printf("Listening on %v", listener.Addr())
	if err := server.Serve(listener); err != nil {
		log.Fatalf("error occurred while listeninig to server %s", err)
	}
}
