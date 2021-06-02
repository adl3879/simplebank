package models

import "time"

type Transfer struct {
	ID            int64
	FromAccountID int64 `db:"from_account_id"`
	ToAccountID   int64 `db:"to_account_id"`
	Amount        int64
	CreatedAt     time.Time `db:"created_at"`
}

type CreateTransferParams struct {
	FromAccountID int64 `db:"from_account_id"`
	ToAccountID   int64 `db:"to_account_id"`
	Amount        int64
}

type ListTransfersParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Limit         int32 `json:"limit"`
	Offset        int32 `json:"offset"`
}
