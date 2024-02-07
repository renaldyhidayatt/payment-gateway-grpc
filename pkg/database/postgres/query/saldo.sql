-- name: CreateSaldo :one
INSERT INTO saldo (user_id, total_balance, withdraw_amount, withdraw_time)VALUES ($1, $2, $3, $4)RETURNING *;

-- name: GetAllSaldo :many
SELECT * FROM saldo;

-- name: GetSaldoById :one
SELECT * FROM saldo WHERE saldo_id = $1 LIMIT 1;

-- name: GetSaldoByUserId :one
SELECT * FROM saldo WHERE user_id = $1 LIMIT 1;

-- name: UpdateSaldoBalance :one
UPDATE saldo SET total_balance = $1 WHERE user_id = $2 RETURNING *;

-- name: GetSaldoByUsers :many
SELECT * FROM saldo WHERE user_id = $1;

-- name: UpdateSaldo :one
UPDATE saldo SET total_balance = $1, withdraw_amount = $2, withdraw_time = $3 WHERE user_id = $4 RETURNING *;

-- name: DeleteSaldo :exec
DELETE FROM saldo WHERE user_id = $1 RETURNING *;
