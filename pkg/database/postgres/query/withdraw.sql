-- Create Withdraw
-- name: CreateWithdraw :one
INSERT INTO
    withdraws (
        card_number,
        withdraw_amount,
        withdraw_time,
        created_at,
        updated_at
    )
VALUES (
        $1,
        $2,
        $3,
        current_timestamp,
        current_timestamp
    ) RETURNING *;

-- Get Withdraw by ID
-- name: GetWithdrawByID :one
SELECT *
FROM withdraws
WHERE
    withdraw_id = $1
    AND deleted_at IS NULL;

-- Get All Active Withdraws
-- name: GetActiveWithdraws :many
SELECT *
FROM withdraws
WHERE
    deleted_at IS NULL
ORDER BY withdraw_time DESC;

-- Get Trashed Withdraws
-- name: GetTrashedWithdraws :many
SELECT *
FROM withdraws
WHERE
    deleted_at IS NOT NULL
ORDER BY withdraw_time DESC;

-- Search Withdraws with Pagination
-- name: SearchWithdraws :many
SELECT *
FROM withdraws
WHERE deleted_at IS NULL
  AND ($1::TEXT IS NULL OR card_number ILIKE '%' || $1 || '%')
ORDER BY withdraw_time DESC
LIMIT $2 OFFSET $3;

-- Count Active Withdraws by Date
-- name: CountActiveWithdrawsByDate :one
SELECT COUNT(*)
FROM withdraws
WHERE deleted_at IS NULL AND withdraw_time::DATE = $1;

-- Trash Withdraw (Soft Delete)
-- name: TrashWithdraw :exec
UPDATE withdraws
SET
    deleted_at = current_timestamp
WHERE
    withdraw_id = $1
    AND deleted_at IS NULL;

-- Restore Withdraw (Undelete)
-- name: RestoreWithdraw :exec
UPDATE withdraws
SET
    deleted_at = NULL
WHERE
    withdraw_id = $1
    AND deleted_at IS NOT NULL;

-- Update Withdraw
-- name: UpdateWithdraw :exec
UPDATE withdraws
SET
    card_number = $2,
    withdraw_amount = $3,
    withdraw_time = $4,
    updated_at = current_timestamp
WHERE
    withdraw_id = $1
    AND deleted_at IS NULL;

-- Delete Withdraw Permanently
-- name: DeleteWithdrawPermanently :exec
DELETE FROM withdraws WHERE withdraw_id = $1;

-- Search Withdraw by Card Number
-- name: SearchWithdrawByCardNumber :many
SELECT *
FROM withdraws
WHERE
    deleted_at IS NULL
    AND card_number ILIKE '%' || $1 || '%'
ORDER BY withdraw_time DESC;

-- Get Trashed By Withdraw ID
-- name: GetTrashedWithdrawByID :one
SELECT *
FROM withdraws
WHERE
    withdraw_id = $1
    AND deleted_at IS NOT NULL;