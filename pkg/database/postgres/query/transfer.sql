-- name: CreateTransfer :one
INSERT INTO transfers (transfer_from, transfer_to, transfer_amount, transfer_time) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetTransferById :one
SELECT * FROM transfers WHERE transfer_id = $1;

-- name: GetTransferByUsers :many
SELECT * FROM transfers WHERE transfer_from = $1 OR transfer_to = $1;

-- name: GetTransferByUserId :one
SELECT * FROM transfers WHERE transfer_from = $1 OR transfer_to = $1;

-- name: GetAllTransfers :many
SELECT * FROM transfers;

-- name: UpdateTransfer :one
UPDATE transfers SET transfer_amount = $1, transfer_time = $2 WHERE transfer_id = $3 RETURNING *;

-- name: DeleteTransfer :exec
DELETE FROM transfers WHERE transfer_from = $1 OR transfer_to = $1 RETURNING *;
