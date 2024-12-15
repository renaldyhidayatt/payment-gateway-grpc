-- Create Saldo
-- name: CreateSaldo :one
INSERT INTO
    saldos (
        card_number,
        total_balance,
        created_at,
        updated_at
    )
VALUES (
        $1,
        $2,
        current_timestamp,
        current_timestamp
    ) RETURNING *;

-- Get Saldo by ID
-- name: GetSaldoByID :one
SELECT * FROM saldos WHERE saldo_id = $1 AND deleted_at IS NULL;

-- Get All Active Saldos
-- name: GetActiveSaldos :many
SELECT * FROM saldos WHERE deleted_at IS NULL ORDER BY saldo_id;

-- Get Trashed Saldos
-- name: GetTrashedSaldos :many
SELECT * FROM saldos WHERE deleted_at IS NOT NULL ORDER BY saldo_id;

-- Search Saldos with Pagination
-- name: GetSaldos :many
SELECT *
FROM saldos
WHERE deleted_at IS NULL
  AND ($1::TEXT IS NULL OR card_number ILIKE '%' || $1 || '%')
ORDER BY saldo_id
LIMIT $2 OFFSET $3;

-- Trash Saldo
-- name: TrashSaldo :exec
UPDATE saldos
SET
    deleted_at = current_timestamp
WHERE
    saldo_id = $1
    AND deleted_at IS NULL;

-- Restore Trashed Saldo
-- name: RestoreSaldo :exec
UPDATE saldos
SET
    deleted_at = NULL
WHERE
    saldo_id = $1
    AND deleted_at IS NOT NULL;

-- Update Saldo
-- name: UpdateSaldo :exec
UPDATE saldos
SET
    card_number = $2,
    total_balance = $3,
    updated_at = current_timestamp
WHERE
    saldo_id = $1
    AND deleted_at IS NULL;

-- Update Saldo Balance
-- name: UpdateSaldoBalance :exec
UPDATE saldos
SET
    total_balance = $2,
    updated_at = current_timestamp
WHERE
    card_number = $1
    AND deleted_at IS NULL;

-- Update Saldo Withdraw
-- name: UpdateSaldoWithdraw :exec
UPDATE saldos
SET
    withdraw_amount = $2,
    total_balance = total_balance - $2,
    withdraw_time = $3,
    updated_at = current_timestamp
WHERE
    card_number = $1
    AND deleted_at IS NULL
    AND total_balance >= $2;

-- Delete Saldo Permanently
-- name: DeleteSaldoPermanently :exec
DELETE FROM saldos WHERE saldo_id = $1;

-- Get Saldo by Card Number
-- name: GetSaldoByCardNumber :one
SELECT * FROM saldos WHERE card_number = $1 AND deleted_at IS NULL;

-- Get Trashed By Saldo ID
-- name: GetTrashedSaldoByID :one
SELECT * FROM saldos WHERE saldo_id = $1 AND deleted_at IS NOT NULL;