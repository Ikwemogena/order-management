package payments

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ikwemogena/order-management/payments/config"
)

type PaystackService struct {
	Config config.PaystackConfig
}

func NewPaystackService(config config.PaystackConfig) *PaystackService {
	return &PaystackService{Config: config}
}

func (s *PaystackService) InitializePayment(req PaymentRequest) (*PaymentResponse, error) {
	client := &http.Client{Timeout: time.Second * 10}

	req.Amount = req.Amount * 100

	requestBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	
	url := fmt.Sprintf("%s/transaction/initialize", s.Config.BaseURL)
	httpReq, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Authorization", "Bearer "+s.Config.SecretKey)
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Cache-Control", "no-cache")

	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var paystackResp PaymentResponse
	if err := json.NewDecoder(resp.Body).Decode(&paystackResp); err != nil {
		return nil, err
	}

	if !paystackResp.Status {
		return nil, fmt.Errorf("paystack error: %s", paystackResp.Message)
	}

	return &PaymentResponse{
		Data: paystackResp.Data,
	}, nil
}

func (s *PaystackService) VerifyPayment(reference string) (*VerifyResponse, error) {
	client := &http.Client{Timeout: time.Second * 10}
	
	url := fmt.Sprintf("%s/transaction/verify/%s", s.Config.BaseURL, reference)
	httpReq, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Authorization", "Bearer "+s.Config.SecretKey)
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Cache-Control", "no-cache")

	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var paystackResp VerifyResponse
	if err := json.NewDecoder(resp.Body).Decode(&paystackResp); err != nil {
		return nil, err
	}

	if !paystackResp.Status {
		return nil, fmt.Errorf("paystack verification error: %s", paystackResp.Message)
	}

	return &VerifyResponse{
		Message: paystackResp.Data.Status,
	}, nil
}
