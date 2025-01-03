-- Create Transfer
-- name: CreateTransfer :one
INSERT INTO
    transfers (
        transfer_from,
        transfer_to,
        transfer_amount,
        transfer_time,
        created_at,
        updated_at
    )
VALUES (
        $1,
        $2,
        $3,
        current_timestamp,
        current_timestamp,
        current_timestamp
    ) RETURNING *;

-- Get Transfer by ID
-- name: GetTransferByID :one
SELECT *
FROM transfers
WHERE
    transfer_id = $1
    AND deleted_at IS NULL;

-- Get All Active Transfers
-- name: GetActiveTransfers :many
SELECT *
FROM transfers
WHERE
    deleted_at IS NULL
ORDER BY transfer_time DESC;

-- Get Trashed Transfers
-- name: GetTrashedTransfers :many
SELECT *
FROM transfers
WHERE
    deleted_at IS NOT NULL
ORDER BY transfer_time DESC;

-- Search Transfers with Pagination
-- name: GetTransfers :many
SELECT *
FROM transfers
WHERE deleted_at IS NULL
  AND ($1::TEXT IS NULL OR transfer_from ILIKE '%' || $1 || '%' OR transfer_to ILIKE '%' || $1 || '%')
ORDER BY transfer_time DESC
LIMIT $2 OFFSET $3;

-- Count Transfers by Date
-- name: CountTransfersByDate :one
SELECT COUNT(*)
FROM transfers
WHERE deleted_at IS NULL
  AND transfer_time::DATE = $1::DATE;

-- Count All Transfers
-- name: CountAllTransfers :one
SELECT COUNT(*) FROM transfers WHERE deleted_at IS NULL;

-- Trash Transfer
-- name: TrashTransfer :exec
UPDATE transfers
SET
    deleted_at = current_timestamp
WHERE
    transfer_id = $1
    AND deleted_at IS NULL;

-- Restore Trashed Transfer
-- name: RestoreTransfer :exec
UPDATE transfers
SET
    deleted_at = NULL
WHERE
    transfer_id = $1
    AND deleted_at IS NOT NULL;

-- Update Transfer
-- name: UpdateTransfer :exec
UPDATE transfers
SET
    transfer_from = $2,
    transfer_to = $3,
    transfer_amount = $4,
    transfer_time = current_timestamp,
    updated_at = current_timestamp
WHERE
    transfer_id = $1
    AND deleted_at IS NULL;

-- Update Transfer Amount
-- name: UpdateTransferAmount :exec
UPDATE transfers
SET
    transfer_amount = $2,
    transfer_time = current_timestamp,
    updated_at = current_timestamp
WHERE
    transfer_id = $1
    AND deleted_at IS NULL;

-- Delete Transfer Permanently
-- name: DeleteTransferPermanently :exec
DELETE FROM transfers WHERE transfer_id = $1;

-- Get Transfers by Card Number (Source or Destination)
-- name: GetTransfersByCardNumber :many
SELECT *
FROM transfers
WHERE
    deleted_at IS NULL
    AND (
        transfer_from = $1
        OR transfer_to = $1
    )
ORDER BY transfer_time DESC;

-- Get Transfers by Source Card
-- name: GetTransfersBySourceCard :many
SELECT *
FROM transfers
WHERE
    deleted_at IS NULL
    AND transfer_from = $1
ORDER BY transfer_time DESC;

-- Get Transfers by Destination Card
-- name: GetTransfersByDestinationCard :many
SELECT *
FROM transfers
WHERE
    deleted_at IS NULL
    AND transfer_to = $1
ORDER BY transfer_time DESC;

-- Get Trashed By Transfer ID
-- name: GetTrashedTransferByID :one
SELECT *
FROM transfers
WHERE
    transfer_id = $1
    AND deleted_at IS NOT NULL;


-- name: GetMonthlyTransferAmounts :many
SELECT
    TO_CHAR(t.transfer_time, 'Mon') AS month,
    SUM(t.transfer_amount) AS total_transfer_amount
FROM
    transfers t
WHERE
    t.deleted_at IS NULL
    AND EXTRACT(YEAR FROM t.transfer_time) = $1 
GROUP BY
    TO_CHAR(t.transfer_time, 'Mon'),
    EXTRACT(MONTH FROM t.transfer_time)
ORDER BY
    EXTRACT(MONTH FROM t.transfer_time);


-- name: GetYearlyTransferAmounts :many
SELECT
    EXTRACT(YEAR FROM t.transfer_time) AS year,
    SUM(t.transfer_amount) AS total_transfer_amount
FROM
    transfers t
WHERE
    t.deleted_at IS NULL
GROUP BY
    EXTRACT(YEAR FROM t.transfer_time)
ORDER BY
    year;



-- name: FindAllTransfersByCardNumberAsSender :many
SELECT
    t.transfer_id,
    t.transfer_from,
    t.transfer_to,
    t.transfer_amount,
    t.transfer_time,
    t.created_at,
    t.updated_at,
    t.deleted_at
FROM
    transfers t
WHERE
    t.transfer_from = $1 
    AND t.deleted_at IS NULL
ORDER BY
    t.transfer_time DESC;


-- name: FindAllTransfersByCardNumberAsReceiver :many
SELECT
    t.transfer_id,
    t.transfer_from,
    t.transfer_to,
    t.transfer_amount,
    t.transfer_time,
    t.created_at,
    t.updated_at,
    t.deleted_at
FROM
    transfers t
WHERE
    t.transfer_to = $1 
    AND t.deleted_at IS NULL
ORDER BY
    t.transfer_time DESC;


