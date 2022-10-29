package service

import (
	"context"
	"database/sql"
	"direct-debit-api/helper"
	"direct-debit-api/model/domain"
	"direct-debit-api/model/web"
	"direct-debit-api/repository"
	"github.com/go-playground/validator/v10"
)

type PaymentServiceImpl struct {
	PaymentRepository repository.PaymentRepository
	DB                *sql.DB
	Validate          *validator.Validate
}

func NewPaymentService(paymentRepository *repository.PaymentRepositoryImpl, DB *sql.DB, validate *validator.Validate) *PaymentServiceImpl {
	return &PaymentServiceImpl{
		PaymentRepository: paymentRepository,
		DB:                DB,
		Validate:          validate,
	}
}

func (service *PaymentServiceImpl) Payment(ctx context.Context, request web.PaymentRequest) web.PaymentResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	payment := domain.Payment{
		PartnerReferenceNo: request.PartnerReferenceNo,
		Amount:             request.Amount,
		Currency:           request.Currency,
	}

	payment = service.PaymentRepository.Payment(ctx, tx, payment)

	return helper.ToPaymentResponse(payment)
}
