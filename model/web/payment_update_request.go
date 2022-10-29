package web

type PaymentUpdateRequest struct {
	PartnerReferenceNo string
	Amount             float64
	Currency           string
}
