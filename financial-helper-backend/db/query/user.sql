-- name: GetUser :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY username;

-- name: CreateUser :one
INSERT INTO users (
     username,
     hashed_password,
     full_name,
     email,
     password_change_at,
     created_at

) VALUES (
             $1, $2, $3, $4, $5, $6
         )
    RETURNING *;

-- name: UpdateUser :exec
UPDATE users
set username = $2,
    hashed_password = $3,
    full_name = $4,
    email = $5,
    password_change_at = $6,
    created_at = $7
WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;