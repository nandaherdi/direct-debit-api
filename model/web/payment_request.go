package web

type PaymentRequest struct {
	PartnerReferenceNo string  `validate:"required"`
	Amount             float64 `validate:"required"`
	Currency           string  `validate:"required"`
}
