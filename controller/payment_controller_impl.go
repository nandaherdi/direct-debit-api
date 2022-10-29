package controller

import (
	"direct-debit-api/helper"
	"direct-debit-api/model/web"
	"direct-debit-api/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type PaymentControllerImpl struct {
	PaymentService service.PaymentService
}

func NewPaymentController(paymentService service.PaymentService) *PaymentControllerImpl {
	return &PaymentControllerImpl{
		PaymentService: paymentService,
	}
}

func (controller *PaymentControllerImpl) Payment(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	paymentCreateRequest := web.PaymentRequest{}
	helper.ReadFromRequestBody(request, &paymentCreateRequest)

	paymentResponse := controller.PaymentService.Payment(request.Context(), paymentCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "Success",
		Data:   paymentResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
