package models

import "time"

type Account struct {
	ID        int64
	Owner     string
	Balance   int64
	Currency  string
	CreatedAt time.Time `db:"created_at"`
}

type CreateAccountParams struct {
	Owner    string
	Balance  int64
	Currency string
}

type UpdateAccountParams struct {
	ID      int64
	Balance int64
}

type ListAccountsParams struct {
	Owner  string `json:"owner"`
	Limit  int32  `json:"limit"`
	Offset int32  `json:"offset"`
}

type AddAccountBalanceParams struct {
	Amount int64 `json:"amount"`
	ID     int64 `json:"id"`
}
