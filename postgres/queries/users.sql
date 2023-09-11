-- name: GetAllUsers :many
SELECT * FROM users ORDER BY id;

-- name: CreateUser :one
INSERT INTO users (name, age, email, phone, address, city) VALUES ($1, $2, $3, $4, $5, $6) RETURNING *;

-- name: GetUserById :one
SELECT * FROM users WHERE id = $1 LIMIT 1;

-- name: DeleteUserById :exec
DELETE FROM users WHERE id = $1;

-- name: UpdateUser :one
UPDATE users SET name = $2 WHERE id = $1 RETURNING *;
