package service

import (
	"context"
	"database/sql"
	"direct-debit-api/helper"
	"direct-debit-api/model/domain"
	"direct-debit-api/model/web"
	"direct-debit-api/repository"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"net/http"
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

func (service *PaymentServiceImpl) Payment(ctx context.Context, requestClient web.PaymentRequest) web.PaymentResponse {
	err := service.Validate.Struct(requestClient)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	payment := domain.Payment{
		PartnerReferenceNo: requestClient.PartnerReferenceNo,
		Amount:             requestClient.Amount,
		Currency:           requestClient.Currency,
	}

	payment = service.PaymentRepository.Payment(ctx, tx, payment)

	router := httprouter.New()
	router.GET("/images/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		//image := params.ByName("image")
		//text := "Image : " + image

		//fmt.Fprintln(writer, requestClient.PartnerReferenceNo)
		//fmt.Fprintln(writer, requestClient.Amount)
		//fmt.Fprintln(writer, requestClient.Currency)
		bytes, _ := json.Marshal(payment)
		fmt.Fprint(writer, string(bytes))

	})
	//fmt.Println(payment)
	server := http.Server{
		Handler: router,
		Addr:    "localhost:4000",
	}

	server.ListenAndServe()

	return helper.ToPaymentResponse(payment)
}
