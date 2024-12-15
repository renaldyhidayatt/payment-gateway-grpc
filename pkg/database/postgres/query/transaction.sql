-- Create Transaction
-- name: CreateTransaction :one
INSERT INTO
    transactions (
        card_number,
        amount,
        payment_method,
        merchant_id,
        transaction_time,
        created_at,
        updated_at
    )
VALUES (
        $1,
        $2,
        $3,
        $4,
        $5,
        current_timestamp,
        current_timestamp
    ) RETURNING *;

-- Get Transaction by ID
-- name: GetTransactionByID :one
SELECT *
FROM transactions
WHERE
    transaction_id = $1
    AND deleted_at IS NULL;

-- Get All Active Transactions
-- name: GetActiveTransactions :many
SELECT *
FROM transactions
WHERE
    deleted_at IS NULL
ORDER BY transaction_time DESC;

-- Get Trashed Transactions
-- name: GetTrashedTransactions :many
SELECT *
FROM transactions
WHERE
    deleted_at IS NOT NULL
ORDER BY transaction_time DESC;

-- Search Transactions with Pagination
-- name: GetTransactions :many
SELECT *
FROM transactions
WHERE deleted_at IS NULL
  AND ($1::TEXT IS NULL OR card_number ILIKE '%' || $1 || '%' OR payment_method ILIKE '%' || $1 || '%')
ORDER BY transaction_time DESC
LIMIT $2 OFFSET $3;

-- Count Transactions by Date
-- name: CountTransactionsByDate :one
SELECT COUNT(*)
FROM transactions
WHERE deleted_at IS NULL
  AND transaction_time::DATE = $1::DATE;

-- Count All Transactions
-- name: CountAllTransactions :one
SELECT COUNT(*) FROM transactions WHERE deleted_at IS NULL;

-- Trash Transaction
-- name: TrashTransaction :exec
UPDATE transactions
SET
    deleted_at = current_timestamp
WHERE
    transaction_id = $1
    AND deleted_at IS NULL;

-- Restore Trashed Transaction
-- name: RestoreTransaction :exec
UPDATE transactions
SET
    deleted_at = NULL
WHERE
    transaction_id = $1
    AND deleted_at IS NOT NULL;

-- Update Transaction
-- name: UpdateTransaction :exec
UPDATE transactions
SET
    card_number = $2,
    amount = $3,
    payment_method = $4,
    merchant_id = $5,
    transaction_time = $6,
    updated_at = current_timestamp
WHERE
    transaction_id = $1
    AND deleted_at IS NULL;

-- Delete Transaction Permanently
-- name: DeleteTransactionPermanently :exec
DELETE FROM transactions WHERE transaction_id = $1;

-- Get Transactions by Card Number
-- name: GetTransactionsByCardNumber :many
SELECT *
FROM transactions
WHERE
    card_number = $1
    AND deleted_at IS NULL
ORDER BY transaction_time DESC;

-- Get Transactions by Merchant ID
-- name: GetTransactionsByMerchantID :many
SELECT *
FROM transactions
WHERE
    merchant_id = $1
    AND deleted_at IS NULL
ORDER BY transaction_time DESC;

-- Get Trashed By Transaction ID
-- name: GetTrashedTransactionByID :one
SELECT *
FROM transactions
WHERE
    transaction_id = $1
    AND deleted_at IS NOT NULL;