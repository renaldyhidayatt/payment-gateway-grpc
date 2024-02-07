-- name: CreateWithdraw :one
INSERT INTO withdraws (user_id, withdraw_amount, withdraw_time) VALUES ($1, $2, $3) RETURNING *;

-- name: GetWithdrawByUsers :many
SELECT * FROM withdraws WHERE user_id = $1;

-- name: GetWithdrawByUserId :one
SELECT * FROM withdraws WHERE user_id = $1;

-- name: GetAllWithdraws :many
SELECT * FROM withdraws;

-- name: GetWithdrawById :one
SELECT * FROM withdraws WHERE withdraw_id = $1;

-- name: UpdateWithdraw :one
UPDATE withdraws SET withdraw_amount = $1, withdraw_time = $2 WHERE withdraw_id = $3 RETURNING *;

-- name: DeleteWithdraw :exec
DELETE FROM withdraws WHERE user_id = $1 RETURNING *;
