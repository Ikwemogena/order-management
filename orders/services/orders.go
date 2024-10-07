package services

import (
	"context"

	pb "github.com/ikwemogena/order-management/orders/proto"
	"github.com/ikwemogena/order-management/stock/db"
)

type Server struct {
    pb.UnimplementedOrderServiceServer
    DB *db.Database
}

func NewServer(db *db.Database) *Server {
	return &Server{DB: db}
}

func (h *Server) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {}
