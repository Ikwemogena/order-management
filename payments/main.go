package main

import (
	"fmt"
	"log"
	"net"

	"github.com/ikwemogena/order-management/payments/config"
	"github.com/ikwemogena/order-management/payments/db"
	pb "github.com/ikwemogena/order-management/payments/proto"
	"github.com/ikwemogena/order-management/payments/services/payments"
	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	paymentGateway := payments.NewPaymentGateway(cfg.Paystack)

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
	pb.RegisterPaymentServiceServer(s, &payments.PaymentServer{
		DB: dbInstance,
		PaymentGateway: paymentGateway,
	})
	fmt.Println("payments service is running on port", cfg.ServerAddress)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
