package main

import (
	"context"
	"fmt"
	"gRPC-with-golang/protogen/golang/orders"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
)

func main() {
	orderServiceAddr := "localhost:50051"
	conn, err := grpc.NewClient(orderServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()
	mux := runtime.NewServeMux()
	if err := orders.RegisterOrdersHandler(context.Background(), mux, conn); err != nil {
		log.Fatalf("could not register handler: %v", err)
	}
	addr := "0.0.0.0:8080"
	fmt.Println("API gateway server is running on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("gateway server closed abruptly: %s", err)
	}
}
