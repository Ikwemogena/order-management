package payments

import (
	"context"
	"errors"
	"log"

	"github.com/ikwemogena/order-management/payments/db"
	pb "github.com/ikwemogena/order-management/payments/proto"
)

// PaymentServer implements the gRPC PaymentService interface
type PaymentServer struct {
	pb.UnimplementedPaymentServiceServer
	DB *db.Database
	PaymentGateway *PaymentGateway
}

// NewPaymentServer creates a new PaymentServer
func NewPaymentServer(paymentProvider *PaymentGateway) *PaymentServer {
	return &PaymentServer{PaymentGateway: paymentProvider}
}

func (s *PaymentServer) InitializePayment(ctx context.Context, req *pb.InitializePaymentRequest) (*pb.InitializePaymentResponse, error) {
	// Fetch the correct payment provider from the PaymentGateway
	provider, err := s.PaymentGateway.GetProvider(req.Provider)
	if err != nil {
		log.Printf("Payment provider error: %v", err)
		return nil, errors.New("unsupported payment provider")
	}

	paymentReq := PaymentRequest{
		Email:     req.Email,
		Amount:    req.Amount,
		Currency:  req.Currency,
	}

	payResp, err := provider.InitializePayment(paymentReq)
	if err != nil {
		log.Printf("Error initializing payment: %v", err)
		return nil, err
	}

	return &pb.InitializePaymentResponse{
		Status:          payResp.Status,
		AuthorizationUrl: payResp.Data.AuthorizationURL,
		Reference:       payResp.Data.Reference,
		Message:         payResp.Message,
		AccessCode: 	payResp.Data.AccessCode,
	}, nil
}

func (s *PaymentServer) VerifyPayment(ctx context.Context, req *pb.VerifyPaymentRequest) (*pb.VerifyPaymentResponse, error) {
	// Fetch the correct payment provider from the PaymentGateway
	provider, err := s.PaymentGateway.GetProvider("paystack")
	if err != nil {
		log.Printf("Payment provider error: %v", err)
		return nil, errors.New("unsupported payment provider")
	}

	payResp, err := provider.VerifyPayment(req.Reference)
	if err != nil {
		log.Printf("Error verifying payment: %v", err)
		return nil, err
	}

	return &pb.VerifyPaymentResponse{
		Status:   payResp.Status,
		Message:  payResp.Message,
	}, nil
}