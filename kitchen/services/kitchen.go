package services

import (
	// "context"
	// "log"

	"context"
	"log"

	"github.com/ikwemogena/order-management/kitchen/db"
	pb "github.com/ikwemogena/order-management/kitchen/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	pb.UnimplementedKitchenServiceServer
    DB *db.Database
}

func NewServer(db *db.Database) *Server {
	return &Server{DB: db}
}


func (s *Server) RecieveOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
    tx, err := s.DB.Begin()
    if err != nil {
        return nil, err
    }

    // var orderID string
    
    _, err = tx.Exec(
        "INSERT INTO kitchen_orders (order_id, user_id, status) VALUES ($1, $2, $3)", req.Id, req.UserId, "pending")

	// tx.QueryRow(
    //     "INSERT INTO kitchen_orders (order_id, user_id, status) VALUES ($1, $2, $3) RETURNING order_id",
    //     req.Id, req.UserId, "pending",
    // ).Scan(&orderID)

    if err != nil {
        tx.Rollback()
        return nil, status.Errorf(codes.Internal, "kitchen failed to accept order: %v", err)
    }

    for _, product := range req.Products {
        _, err := tx.Exec("INSERT INTO kitchen_order_items (order_id, product_id, quantity) VALUES ($1, $2, $3)", req.Id, product.Id, product.Quantity)
        if err != nil {
            tx.Rollback()
            return nil, status.Errorf(codes.Internal, "kitchen failed to add order items: %v", err)
        }
    }

    if err := tx.Commit(); err != nil {
        log.Println(err.Error(), "error")
        return nil, status.Errorf(codes.Internal, "Kitchen failed to create order: %v", err)
    }

	return &pb.CreateOrderResponse{
		Status: "PREPARING",
		Message: "Order is being prepared",
    }, nil
}