// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"context"
)

type Querier interface {
	AddWalletBalance(ctx context.Context, arg AddWalletBalanceParams) (Wallet, error)
	CreateExpense(ctx context.Context, arg CreateExpenseParams) (Expense, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	CreateWallet(ctx context.Context, arg CreateWalletParams) (Wallet, error)
	DecreaseWalletBalance(ctx context.Context, arg DecreaseWalletBalanceParams) (Wallet, error)
	DeleteExpense(ctx context.Context, id int64) error
	DeleteUser(ctx context.Context, id int64) error
	DeleteWallet(ctx context.Context, arg DeleteWalletParams) error
	GetExpensesByCategory(ctx context.Context, arg GetExpensesByCategoryParams) ([]Expense, error)
	GetExpensesByDate(ctx context.Context, arg GetExpensesByDateParams) ([]Expense, error)
	GetExpensesByName(ctx context.Context, arg GetExpensesByNameParams) ([]Expense, error)
	GetExpensesByOwner(ctx context.Context, owner int64) ([]Expense, error)
	GetExpensesByWallet(ctx context.Context, arg GetExpensesByWalletParams) ([]Expense, error)
	GetUser(ctx context.Context, email string) (User, error)
	GetWallet(ctx context.Context, arg GetWalletParams) (Wallet, error)
	GetWalletsByOwner(ctx context.Context, owner int64) ([]Wallet, error)
	ListUsers(ctx context.Context) ([]User, error)
	ListWallets(ctx context.Context, arg ListWalletsParams) ([]Wallet, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) error
	UpdateWallet(ctx context.Context, arg UpdateWalletParams) (Wallet, error)
}

var _ Querier = (*Queries)(nil)
