-- name: CreateTopup :one
INSERT INTO topups (user_id, topup_no, topup_amount, topup_method, topup_time) VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: GetAllTopups :many
SELECT * FROM topups;

-- name: GetTopupByUsers :many
SELECT * FROM topups WHERE user_id = $1;

-- name: GetTopupByUserId :one
SELECT * FROM topups WHERE topup_id = $1 LIMIT 1;

-- name: GetTopupById :one
SELECT * FROM topups WHERE topup_id = $1;

-- name: UpdateTopup :one
UPDATE topups SET topup_amount = $1, topup_method = $2, topup_time = $3 WHERE topup_id = $4 RETURNING *;

-- name: DeleteTopup :exec
DELETE FROM topups WHERE user_id = $1 RETURNING *;
