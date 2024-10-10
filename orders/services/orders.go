package services

import (
	"context"
	"fmt"
	"log"

	"github.com/ikwemogena/order-management/orders/db"
	pb "github.com/ikwemogena/order-management/orders/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
    pb.UnimplementedOrderServiceServer
    DB *db.Database
}

func NewServer(db *db.Database) *Server {
	return &Server{DB: db}
}

func (s *Server) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
    tx, err := s.DB.Begin()
    if err != nil {
        return nil, err
    }

    var orderID string
    
    err = tx.QueryRow(
        "INSERT INTO orders (user_id, total_amount, shipping_address, payment_method, status) VALUES ($1, $2, $3, $4, $5) RETURNING id",
        req.UserId, req.Amount, req.ShippingAddress, req.PaymentMethod, "pending",
    ).Scan(&orderID)

    if err != nil {
        tx.Rollback()
        log.Println(err.Error(), "error here")
        return nil, status.Errorf(codes.Internal, "Failed to create order: %v", err)
    }

    for _, product := range req.Products {
        _, err := tx.Exec("INSERT INTO order_items (order_id, product_id, unit_price, quantity, total_price) VALUES ($1, $2, $3, $4, $5)", orderID, product.Id, product.UnitPrice, product.Quantity, product.TotalAmount)
        if err != nil {
            tx.Rollback()
            log.Println(err.Error(), "error")
            return nil, status.Errorf(codes.Internal, "Failed to add order items order: %v", err)
        }
    }

    if err := tx.Commit(); err != nil {
        log.Println(err.Error(), "error")
        return nil, status.Errorf(codes.Internal, "Failed to create order11: %v", err)
    }

    return &pb.CreateOrderResponse{OrderId: fmt.Sprintf("%s", orderID)}, nil
}

func (s *Server) GetOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.GetOrderResponse, error) {
    var userID, status string
    err := s.DB.QueryRow("SELECT user_id, status FROM orders WHERE order_id=$1", req.OrderId).Scan(&userID, &status)
    if err != nil {
        return nil, err
    }

    rows, err := s.DB.Query("SELECT product_id FROM order_items WHERE order_id=$1", req.OrderId)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var productIDs []string
    for rows.Next() {
        var productID string
        if err := rows.Scan(&productID); err != nil {
            return nil, err
        }
        productIDs = append(productIDs, productID)
    }

    return &pb.GetOrderResponse{
        OrderId:    req.OrderId,
        UserId:     userID,
        ProductIds: productIDs,
        Status:     status,
    }, nil
}

func generateOrderID() {
    
}