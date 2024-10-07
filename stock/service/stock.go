package service

import (
	"context"

	"github.com/ikwemogena/order-management/stock/db"
	pb "github.com/ikwemogena/order-management/stock/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
    pb.UnimplementedStockServiceServer
    DB *db.Database
}

func NewServer(db *db.Database) *Server {
	return &Server{DB: db}
}

func (s *Server) CheckStock(ctx context.Context, req *pb.CheckStockRequest) (*pb.CheckStockResponse, error) {
    var quantity int32

    err := s.DB.QueryRow("SELECT quantity FROM stock WHERE item_id = $1", req.ItemId).Scan(&quantity)

    if err != nil {
        return nil, status.Errorf(codes.NotFound, "Item %s not found in inventory", req.ItemId)
    }

    available := quantity >= req.Quantity
    return &pb.CheckStockResponse{
        Available:         available,
        AvailableQuantity: quantity,
    }, nil
}

func (s *Server) UpdateStock(ctx context.Context, req *pb.UpdateStockRequest) (*pb.UpdateStockResponse, error) {
    _, err := s.DB.Exec("UPDATE stock SET quantity = $1 WHERE id = $2", req.QuantityChange, req.ItemId)
    if err != nil {
        return nil, status.Errorf(codes.Internal, "Failed to update stock: %v", err)
    }

    return &pb.UpdateStockResponse{}, nil
}

func (s *Server) CreateStock(_ context.Context, in *pb.CreateStockRequest) (*pb.CreateStockResponse, error) {

    _, err := s.DB.Exec("INSERT INTO stock (name, description, quantity) VALUES ($1, $2, $3)", in.ItemName, in.ItemDescription, in.Quantity)

    if err != nil {
        return nil, status.Errorf(codes.Internal, "Failed to create stock: %v", err)
    }

    return &pb.CreateStockResponse{
        Success: true,
    }, nil
}