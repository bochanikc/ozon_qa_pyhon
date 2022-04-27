package main

import (
	"context"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"gitlab.ozon.dev/qa/teachers/qa-route-256/item-api/internal/app"
	"gitlab.ozon.dev/qa/teachers/qa-route-256/item-api/internal/database"
	pb "gitlab.ozon.dev/qa/teachers/qa-route-256/item-api/pkg/api"
	"google.golang.org/grpc"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to Listen: %v", err)
	}
	server := grpc.NewServer()
	db, err := database.NewDatabase("postgresql://test:test@localhost:5432/db?sslmode=disable&binary_parameters=yes")
	if err != nil {
		log.Fatalf("failed to create database: %v", err)
	}
	app.RegisterNewItemService(server, db)
	go func() {
		if err := server.Serve(lis); err != nil {
			log.Fatalf("failed to Serve: %v", err)
		}
	}()

	go func() {
		log.Println(http.ListenAndServe("localhost:50053", nil))
	}()

	mux := runtime.NewServeMux()
	err = pb.RegisterItemHandlerServer(ctx, mux, app.NewItemService(db))
	if err != nil {
		log.Fatalf("failed to RegisterItemHandlerServer: %v", err)
	}
	if err := http.ListenAndServe(":50052", mux); err != nil {
		log.Fatalf("failed to ListenAndServe: %v", err)
	}
}
