package controllers

import (
	"context"

	"github.com/adl3879/simple_bank/models"
)

type Querier interface {
	AddAccountBalance(ctx context.Context, arg models.AddAccountBalanceParams) (models.Account, error)
	CreateAccount(ctx context.Context, arg models.CreateAccountParams) (models.Account, error)
	CreateEntry(ctx context.Context, arg models.CreateEntryParams) (models.Entry, error)
	CreateTransfer(ctx context.Context, arg models.CreateTransferParams) (models.Transfer, error)
	CreateUser(ctx context.Context, arg models.CreateUserParams) (models.User, error)
	DeleteAccount(ctx context.Context, id int64) error
	GetAccount(ctx context.Context, id int64) (models.Account, error)
	GetAccountForUpdate(ctx context.Context, id int64) (models.Account, error)
	GetEntry(ctx context.Context, id int64) (models.Entry, error)
	GetTransfer(ctx context.Context, id int64) (models.Transfer, error)
	GetUser(ctx context.Context, username string) (models.User, error)
	ListAccounts(ctx context.Context, arg models.ListAccountsParams) ([]models.Account, error)
	ListEntries(ctx context.Context, arg models.ListEntriesParams) ([]models.Entry, error)
	ListTransfers(ctx context.Context, arg models.ListTransfersParams) ([]models.Transfer, error)
	UpdateAccount(ctx context.Context, arg models.UpdateAccountParams) (models.Account, error)
}

var _ Querier = (*Queries)(nil)
