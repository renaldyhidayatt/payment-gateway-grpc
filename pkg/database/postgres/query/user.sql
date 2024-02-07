-- name: CreateUser :one
INSERT INTO users (firstname, lastname,email, password, noc_transfer) VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: GetAllUsers :many
SELECT * FROM users;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: GetUserById :one
SELECT * FROM users WHERE user_id = $1;

-- name: UpdateUser :one
UPDATE users SET noc_transfer = $1, firstname = $2, lastname = $3, email = $4, password = $5 WHERE user_id = $6 RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users WHERE user_id = $1 RETURNING *;
