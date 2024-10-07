package main

import (
	"fmt"
	"log"
	"net"

	"github.com/ikwemogena/order-management/stock/config"
	"github.com/ikwemogena/order-management/stock/db"
	pb "github.com/ikwemogena/order-management/stock/proto"
	"github.com/ikwemogena/order-management/stock/service"
	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	database, err := db.Init(cfg.DatabaseURL)
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
    defer database.Close()

	dbInstance := &db.Database{DB: database}
    if err := dbInstance.InitTables(); err != nil {
        log.Fatalf("Failed to initialize tables: %v", err)
    }


	lis, err := net.Listen("tcp", cfg.ServerAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterStockServiceServer(s, &service.Server{DB: dbInstance})
	fmt.Println("Stock service is running on port", cfg.ServerAddress)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
