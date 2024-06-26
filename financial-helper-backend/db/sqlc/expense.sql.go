// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: expense.sql

package db

import (
	"context"
	"time"
)

const createExpense = `-- name: CreateExpense :one
INSERT INTO expenses (owner,
                      wallet,
                      wallet_name,
                      currency,
                      value,
                      name,
                      category,
                      date)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id, owner, wallet, wallet_name, currency, value, name, category, date, created_at
`

type CreateExpenseParams struct {
	Owner      int64
	Wallet     int64
	WalletName string
	Currency   string
	Value      float64
	Name       string
	Category   ExpenseCategory
	Date       time.Time
}

func (q *Queries) CreateExpense(ctx context.Context, arg CreateExpenseParams) (Expense, error) {
	row := q.db.QueryRowContext(ctx, createExpense,
		arg.Owner,
		arg.Wallet,
		arg.WalletName,
		arg.Currency,
		arg.Value,
		arg.Name,
		arg.Category,
		arg.Date,
	)
	var i Expense
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Wallet,
		&i.WalletName,
		&i.Currency,
		&i.Value,
		&i.Name,
		&i.Category,
		&i.Date,
		&i.CreatedAt,
	)
	return i, err
}

const deleteExpense = `-- name: DeleteExpense :exec
DELETE
FROM expenses
WHERE id = $1
`

func (q *Queries) DeleteExpense(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteExpense, id)
	return err
}

const getExpensesByCategory = `-- name: GetExpensesByCategory :many
SELECT id, owner, wallet, wallet_name, currency, value, name, category, date, created_at
FROM expenses
WHERE owner = $1
  AND category = $2
`

type GetExpensesByCategoryParams struct {
	Owner    int64
	Category ExpenseCategory
}

func (q *Queries) GetExpensesByCategory(ctx context.Context, arg GetExpensesByCategoryParams) ([]Expense, error) {
	rows, err := q.db.QueryContext(ctx, getExpensesByCategory, arg.Owner, arg.Category)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Expense
	for rows.Next() {
		var i Expense
		if err := rows.Scan(
			&i.ID,
			&i.Owner,
			&i.Wallet,
			&i.WalletName,
			&i.Currency,
			&i.Value,
			&i.Name,
			&i.Category,
			&i.Date,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getExpensesByDate = `-- name: GetExpensesByDate :many
SELECT id, owner, wallet, wallet_name, currency, value, name, category, date, created_at
FROM expenses
WHERE owner = $1 AND date = $2
`

type GetExpensesByDateParams struct {
	Owner int64
	Date  time.Time
}

func (q *Queries) GetExpensesByDate(ctx context.Context, arg GetExpensesByDateParams) ([]Expense, error) {
	rows, err := q.db.QueryContext(ctx, getExpensesByDate, arg.Owner, arg.Date)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Expense
	for rows.Next() {
		var i Expense
		if err := rows.Scan(
			&i.ID,
			&i.Owner,
			&i.Wallet,
			&i.WalletName,
			&i.Currency,
			&i.Value,
			&i.Name,
			&i.Category,
			&i.Date,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getExpensesByName = `-- name: GetExpensesByName :many
SELECT id, owner, wallet, wallet_name, currency, value, name, category, date, created_at
FROM expenses
WHERE owner = $1
  and name = $2
`

type GetExpensesByNameParams struct {
	Owner int64
	Name  string
}

func (q *Queries) GetExpensesByName(ctx context.Context, arg GetExpensesByNameParams) ([]Expense, error) {
	rows, err := q.db.QueryContext(ctx, getExpensesByName, arg.Owner, arg.Name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Expense
	for rows.Next() {
		var i Expense
		if err := rows.Scan(
			&i.ID,
			&i.Owner,
			&i.Wallet,
			&i.WalletName,
			&i.Currency,
			&i.Value,
			&i.Name,
			&i.Category,
			&i.Date,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getExpensesByOwner = `-- name: GetExpensesByOwner :many
SELECT id, owner, wallet, wallet_name, currency, value, name, category, date, created_at
FROM expenses
WHERE owner = $1
`

func (q *Queries) GetExpensesByOwner(ctx context.Context, owner int64) ([]Expense, error) {
	rows, err := q.db.QueryContext(ctx, getExpensesByOwner, owner)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Expense
	for rows.Next() {
		var i Expense
		if err := rows.Scan(
			&i.ID,
			&i.Owner,
			&i.Wallet,
			&i.WalletName,
			&i.Currency,
			&i.Value,
			&i.Name,
			&i.Category,
			&i.Date,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getExpensesByWallet = `-- name: GetExpensesByWallet :many
SELECT id, owner, wallet, wallet_name, currency, value, name, category, date, created_at
FROM expenses
WHERE owner = $1
  AND wallet = $2
`

type GetExpensesByWalletParams struct {
	Owner  int64
	Wallet int64
}

func (q *Queries) GetExpensesByWallet(ctx context.Context, arg GetExpensesByWalletParams) ([]Expense, error) {
	rows, err := q.db.QueryContext(ctx, getExpensesByWallet, arg.Owner, arg.Wallet)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Expense
	for rows.Next() {
		var i Expense
		if err := rows.Scan(
			&i.ID,
			&i.Owner,
			&i.Wallet,
			&i.WalletName,
			&i.Currency,
			&i.Value,
			&i.Name,
			&i.Category,
			&i.Date,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
