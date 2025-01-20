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


-- name: GetMonthlyTotalSaldoBalance :many
WITH monthly_data AS (
    SELECT
        EXTRACT(YEAR FROM s.created_at)::integer AS year,
        EXTRACT(MONTH FROM s.created_at)::integer AS month,
        SUM(s.total_balance) AS total_balance
    FROM
        saldos s
    WHERE
        s.deleted_at IS NULL
        AND (
            (s.created_at >= $1::timestamp AND s.created_at <= $2::timestamp)
            OR (s.created_at >= $3::timestamp AND s.created_at <= $4::timestamp)
        )
    GROUP BY
        EXTRACT(YEAR FROM s.created_at),
        EXTRACT(MONTH FROM s.created_at)
), formatted_data AS (
    SELECT
        year::text,
        TO_CHAR(TO_DATE(month::text, 'MM'), 'Mon') AS month,
        total_balance
    FROM
        monthly_data

    UNION ALL

    SELECT
        EXTRACT(YEAR FROM $1::timestamp)::text AS year,
        TO_CHAR($1::timestamp, 'Mon') AS month,
        0 AS total_balance
    WHERE NOT EXISTS (
        SELECT 1
        FROM monthly_data
        WHERE year = EXTRACT(YEAR FROM $1::timestamp)::integer
        AND month = EXTRACT(MONTH FROM $1::timestamp)::integer
    )

    UNION ALL

    SELECT
        EXTRACT(YEAR FROM $3::timestamp)::text AS year,
        TO_CHAR($3::timestamp, 'Mon') AS month,
        0 AS total_balance
    WHERE NOT EXISTS (
        SELECT 1
        FROM monthly_data
        WHERE year = EXTRACT(YEAR FROM $3::timestamp)::integer
        AND month = EXTRACT(MONTH FROM $3::timestamp)::integer
    )
)
SELECT * FROM formatted_data
ORDER BY
    year DESC,
    TO_DATE(month, 'Mon') DESC;



-- name: GetYearlyTotalSaldoBalances :many
WITH yearly_data AS (
    SELECT
        EXTRACT(YEAR FROM s.created_at)::integer AS year,
        COALESCE(SUM(s.total_balance), 0)::integer AS total_balance
    FROM
        saldos s
    WHERE
        s.deleted_at IS NULL
        AND (
            EXTRACT(YEAR FROM s.created_at) = $1::integer
            OR EXTRACT(YEAR FROM s.created_at) = $1::integer - 1
        )
    GROUP BY
        EXTRACT(YEAR FROM s.created_at)
), formatted_data AS (
    SELECT
        year::text,
        total_balance::integer
    FROM
        yearly_data

    UNION ALL

    SELECT
        $1::text AS year,
        0::integer AS total_balance
    WHERE NOT EXISTS (
        SELECT 1
        FROM yearly_data
        WHERE year = $1::integer
    )

    UNION ALL

    SELECT
        ($1::integer - 1)::text AS year,
        0::integer AS total_balance
    WHERE NOT EXISTS (
        SELECT 1
        FROM yearly_data
        WHERE year = $1::integer - 1
    )
)
SELECT * FROM formatted_data
ORDER BY
    year DESC;




-- name: GetMonthlySaldoBalances :many
WITH months AS (
    SELECT generate_series(
        date_trunc('year', $1::timestamp),
        date_trunc('year', $1::timestamp) + interval '1 year' - interval '1 day',
        interval '1 month'
    ) AS month
)
SELECT
    TO_CHAR(m.month, 'Mon') AS month,
    COALESCE(SUM(s.total_balance), 0)::int AS total_balance
FROM
    months m
LEFT JOIN
    saldos s ON EXTRACT(MONTH FROM s.created_at) = EXTRACT(MONTH FROM m.month)
    AND EXTRACT(YEAR FROM s.created_at) = EXTRACT(YEAR FROM m.month)
    AND s.deleted_at IS NULL
GROUP BY
    m.month
ORDER BY
    m.month;


-- name: GetYearlySaldoBalances :many
WITH last_five_years AS (
    SELECT
        EXTRACT(YEAR FROM s.created_at) AS year,
        SUM(s.total_balance) AS total_balance
    FROM
        saldos s
    WHERE
        s.deleted_at IS NULL
        AND EXTRACT(YEAR FROM s.created_at) >= $1 - 4
        AND EXTRACT(YEAR FROM s.created_at) <= $1
    GROUP BY
        EXTRACT(YEAR FROM s.created_at)
)
SELECT
    year,
    total_balance
FROM
    last_five_years
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
