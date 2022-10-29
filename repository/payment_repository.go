package repository

import (
	"context"
	"database/sql"
	"direct-debit-api/model/domain"
)

type PaymentRepository interface {
	Payment(ctx context.Context, tx *sql.Tx, category domain.Payment) domain.Payment
}
