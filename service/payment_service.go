package service

import (
	"context"
	"direct-debit-api/model/web"
)

type PaymentService interface {
	Payment(ctx context.Context, request web.PaymentRequest) web.PaymentResponse
}
