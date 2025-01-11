-- Search Saldos with Pagination and Total Count
-- name: GetSaldos :many
SELECT
    *,
    COUNT(*) OVER() AS total_count
FROM saldos
WHERE deleted_at IS NULL
  AND ($1::TEXT IS NULL OR card_number ILIKE '%' || $1 || '%')
ORDER BY saldo_id
LIMIT $2 OFFSET $3;


-- Get Saldo by ID
-- name: GetSaldoByID :one
SELECT * FROM saldos WHERE saldo_id = $1 AND deleted_at IS NULL;


-- Get All Active Saldos with Pagination, Search, and Total Count
-- name: GetActiveSaldos :many
SELECT
    *,
    COUNT(*) OVER() AS total_count
FROM saldos
WHERE deleted_at IS NULL
  AND ($1::TEXT IS NULL OR card_number ILIKE '%' || $1 || '%')
ORDER BY saldo_id
LIMIT $2 OFFSET $3;

-- Get Trashed Saldos with Pagination, Search, and Total Count
-- name: GetTrashedSaldos :many
SELECT
    *,
    COUNT(*) OVER() AS total_count
FROM saldos
WHERE deleted_at IS NOT NULL
  AND ($1::TEXT IS NULL OR card_number ILIKE '%' || $1 || '%')
ORDER BY saldo_id
LIMIT $2 OFFSET $3;



-- Get Trashed By Saldo ID
-- name: GetTrashedSaldoByID :one
SELECT * FROM saldos WHERE saldo_id = $1 AND deleted_at IS NOT NULL;

-- name: GetMonthlyTotalBalance :many
SELECT
    TO_CHAR(s.created_at, 'Mon') AS month,
    SUM(s.total_balance) AS total_balance
FROM
    saldos s
WHERE
    s.deleted_at IS NULL
    AND EXTRACT(YEAR FROM s.created_at) = $1
GROUP BY
    TO_CHAR(s.created_at, 'Mon'),
    EXTRACT(MONTH FROM s.created_at)
ORDER BY
    EXTRACT(MONTH FROM s.created_at);

-- name: GetYearlyTotalBalance :many
SELECT
    EXTRACT(YEAR FROM s.created_at) AS year,
    SUM(s.total_balance) AS total_balance
FROM
    saldos s
WHERE
    s.deleted_at IS NULL
GROUP BY
    EXTRACT(YEAR FROM s.created_at)
ORDER BY
    year;

-- Get Saldo by Card Number
-- name: GetSaldoByCardNumber :one
SELECT * FROM saldos WHERE card_number = $1 AND deleted_at IS NULL;


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


-- Delete Saldo Permanently
-- name: DeleteSaldoPermanently :exec
DELETE FROM saldos WHERE saldo_id = $1 AND deleted_at IS NOT NULL;


-- Restore All Trashed Saldos
-- name: RestoreAllSaldos :exec
UPDATE saldos
SET
    deleted_at = NULL
WHERE
    deleted_at IS NOT NULL;


-- Delete All Trashed Saldos Permanently
-- name: DeleteAllPermanentSaldos :exec
DELETE FROM saldos
WHERE
    deleted_at IS NOT NULL;


-- name: CountSaldos :one
SELECT COUNT(*)
FROM saldos
WHERE deleted_at IS NULL
  AND ($1::TEXT IS NULL OR card_number ILIKE '%' || $1 || '%');

-- name: CountAllSaldos :one
SELECT COUNT(*)
FROM saldos
WHERE deleted_at IS NULL;
