package helper

import (
	"direct-debit-api/model/domain"
	"direct-debit-api/model/web"
)

func ToPaymentResponse(category domain.Payment) web.PaymentResponse {
	return web.PaymentResponse{
		PartnerReferenceNo: category.PartnerReferenceNo,
		Amount:             category.Amount,
		Currency:           category.Currency,
	}
}

func ToPaymentResponses(payments []domain.Payment) []web.PaymentResponse {
	var paymentResponses []web.PaymentResponse
	for _, payment := range payments {
		paymentResponses = append(paymentResponses, ToPaymentResponse(payment))
	}
	return paymentResponses
}
