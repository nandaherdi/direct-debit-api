package repository

import (
	"context"
	"database/sql"
	"direct-debit-api/helper"
	"direct-debit-api/model/domain"
)

type PaymentRepositoryImpl struct {
}

func NewPaymentRepository() *PaymentRepositoryImpl {
	return &PaymentRepositoryImpl{}
}

func (repository *PaymentRepositoryImpl) Payment(ctx context.Context, tx *sql.Tx, payment domain.Payment) domain.Payment {
	SQL := "insert into payments(partner_reference_no, amount, currency) values (?, ?, ?)"
	_, err := tx.ExecContext(ctx, SQL, payment.PartnerReferenceNo, payment.Amount, payment.Currency)
	helper.PanicIfError(err)

	//id, err := result.LastInsertId()
	helper.PanicIfError(err)

	//payment.PartnerReferenceNo = string(id)
	return payment
}
