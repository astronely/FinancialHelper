-- name: CreateWallet :one
INSERT INTO wallets (owner,
                     name,
                     balance,
                     currency)
VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetWalletsByOwner :many
SELECT *
FROM wallets
WHERE owner = $1;

-- name: GetWallet :one
SELECT *
FROM wallets
WHERE owner = $1 AND name = $2 LIMIT 1;


-- name: ListWallets :many
SELECT *
FROM wallets
ORDER BY id LIMIT $1
OFFSET $2;

-- name: UpdateWallet :one
UPDATE wallets
SET balance = $3
WHERE owner = $1 AND name = $2 RETURNING *;

-- name: DecreaseWalletBalance :one
UPDATE wallets
set balance = ROUND(CAST(balance - $3 AS numeric), 2)
WHERE owner = $1 AND name = $2 RETURNING *;

-- name: AddWalletBalance :one
UPDATE wallets
SET balance = ROUND(CAST(balance + $3 AS numeric), 2)
WHERE owner = $1 AND name = $2 RETURNING *;

-- name: DeleteWallet :exec
DELETE
FROM wallets
WHERE owner = $1 AND name = $2;

