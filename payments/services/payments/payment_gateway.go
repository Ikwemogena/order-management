package payments

import (
	"errors"

	"github.com/ikwemogena/order-management/payments/config"
)

type PaymentProvider interface {
    InitializePayment(req PaymentRequest) (*PaymentResponse, error)
    VerifyPayment(reference string) (*VerifyResponse, error)
}

type PaymentRequest struct {
    Email     string `json:"email"`
    Amount    int32 `json:"amount"`
    Currency  string `json:"currency"`
}

type PaymentResponse struct {
    Status  bool   `json:"status"`
    Message string `json:"message"`
    Data    struct {
        AuthorizationURL string `json:"authorization_url"`
        AccessCode       string `json:"access_code"`
        Reference        string `json:"reference"`
    } `json:"data"`
}


type VerifyResponse struct {
    Status  bool
    Message string
    Data    struct {
        Amount    float64
        Currency  string
        Reference string
        Status    string
    }
}

type PaymentGateway struct {
    providers map[string]PaymentProvider
}

func NewPaymentGateway(paystackConfig config.PaystackConfig) *PaymentGateway {
    return &PaymentGateway{
        providers: map[string]PaymentProvider{
            "paystack": NewPaystackService(paystackConfig),
        },
    }
}

func (pg *PaymentGateway) GetProvider(providerName string) (PaymentProvider, error) {
    provider, exists := pg.providers[providerName]
    if !exists {
        return nil, errors.New("payment provider not found")
    }
    return provider, nil
}
