-- name: CreateExpense :one
INSERT INTO expenses (owner,
                      wallet,
                      wallet_name,
                      currency,
                      value,
                      name,
                      category,
                      date)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING *;

-- name: GetExpensesByOwner :many
SELECT *
FROM expenses
WHERE owner = $1;

-- name: GetExpensesByWallet :many
SELECT *
FROM expenses
WHERE owner = $1
  AND wallet = $2;

-- name: GetExpensesByName :many
SELECT *
FROM expenses
WHERE owner = $1
  and name = $2;

-- name: GetExpensesByCategory :many
SELECT *
FROM expenses
WHERE owner = $1
  AND category = $2;

-- name: GetExpensesByDate :many
SELECT *
FROM expenses
WHERE owner = $1 AND date = $2;

-- name: DeleteExpense :exec
DELETE
FROM expenses
WHERE id = $1;

