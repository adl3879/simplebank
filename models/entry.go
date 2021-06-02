package models

import "time"

type Entry struct {
	ID        int64
	AccountID int64 `db:"account_id"`
	Amount    int64
	CreatedAt time.Time `db:"created_at"`
}

type CreateEntryParams struct {
	AccountID int64 `db:"account_id"`
	Amount    int64
}

type UpdateEntryParams struct {
	ID     int64
	Amount int64
}

type ListEntriesParams struct {
	AccountID int64 `json:"account_id"`
	Limit     int32 `json:"limit"`
	Offset    int32 `json:"offset"`
}
