-- name: GetMerchants :many
SELECT
    *,
    COUNT(*) OVER() AS total_count
FROM merchants
WHERE deleted_at IS NULL
    AND ($1::TEXT IS NULL OR name ILIKE '%' || $1 || '%' OR api_key ILIKE '%' || $1 || '%' OR status ILIKE '%' || $1 || '%')
ORDER BY merchant_id
LIMIT $2 OFFSET $3;


-- Get Merchant by ID
-- name: GetMerchantByID :one
SELECT *
FROM merchants
WHERE
    merchant_id = $1
    AND deleted_at IS NULL;

-- Get Merchant by API Key
-- name: GetMerchantByApiKey :one
SELECT * FROM merchants WHERE api_key = $1 AND deleted_at IS NULL;

-- Get Merchant by Name
-- name: GetMerchantByName :one
SELECT * FROM merchants WHERE name = $1 AND deleted_at IS NULL;

-- Get Merchants by User ID
-- name: GetMerchantsByUserID :many
SELECT * FROM merchants WHERE user_id = $1 AND deleted_at IS NULL;


-- name: GetActiveMerchants :many
SELECT
    *,
    COUNT(*) OVER() AS total_count
FROM merchants
WHERE deleted_at IS NULL
    AND ($1::TEXT IS NULL OR name ILIKE '%' || $1 || '%' OR api_key ILIKE '%' || $1 || '%' OR status ILIKE '%' || $1 || '%')
ORDER BY merchant_id
LIMIT $2 OFFSET $3;

-- name: GetTrashedMerchants :many
SELECT
    *,
    COUNT(*) OVER() AS total_count
FROM merchants
WHERE deleted_at IS NOT NULL
    AND ($1::TEXT IS NULL OR name ILIKE '%' || $1 || '%' OR api_key ILIKE '%' || $1 || '%' OR status ILIKE '%' || $1 || '%')
ORDER BY merchant_id
LIMIT $2 OFFSET $3;


-- Get Trashed By Merchant ID
-- name: GetTrashedMerchantByID :one
SELECT *
FROM merchants
WHERE
    merchant_id = $1
    AND deleted_at IS NOT NULL;



-- name: GetMonthlyPaymentMethodsMerchant :many
WITH months AS (
    SELECT generate_series(
        date_trunc('year', $1::timestamp),
        date_trunc('year', $1::timestamp) + interval '1 year' - interval '1 day',
        interval '1 month'
    ) AS month
),
payment_methods AS (
    SELECT DISTINCT payment_method
    FROM transactions
    WHERE deleted_at IS NULL
)
SELECT
    TO_CHAR(m.month, 'Mon') AS month,
    pm.payment_method,
    COALESCE(SUM(t.amount), 0)::int AS total_amount
FROM
    months m
CROSS JOIN
    payment_methods pm
LEFT JOIN
    transactions t ON EXTRACT(MONTH FROM t.transaction_time) = EXTRACT(MONTH FROM m.month)
    AND EXTRACT(YEAR FROM t.transaction_time) = EXTRACT(YEAR FROM m.month)
    AND t.payment_method = pm.payment_method
    AND t.deleted_at IS NULL
LEFT JOIN
    merchants mch ON t.merchant_id = mch.merchant_id
    AND mch.deleted_at IS NULL
GROUP BY
    m.month,
    pm.payment_method
ORDER BY
    m.month,
    pm.payment_method;



-- name: GetYearlyPaymentMethodMerchant :many
WITH last_five_years AS (
    SELECT
        EXTRACT(YEAR FROM t.transaction_time) AS year,
        t.payment_method,
        SUM(t.amount) AS total_amount
    FROM
        transactions t
    JOIN
        merchants m ON t.merchant_id = m.merchant_id
    WHERE
        t.deleted_at IS NULL AND m.deleted_at IS NULL
        AND EXTRACT(YEAR FROM t.transaction_time) >= $1 - 4
        AND EXTRACT(YEAR FROM t.transaction_time) <= $1
    GROUP BY
        EXTRACT(YEAR FROM t.transaction_time),
        t.payment_method
)
SELECT
    year,
    payment_method,
    total_amount
FROM
    last_five_years
ORDER BY
    year;


-- name: GetMonthlyAmountMerchant :many
WITH months AS (
    SELECT generate_series(
        date_trunc('year', $1::timestamp),
        date_trunc('year', $1::timestamp) + interval '1 year' - interval '1 day',
        interval '1 month'
    ) AS month
)
SELECT
    TO_CHAR(m.month, 'Mon') AS month,
    COALESCE(SUM(t.amount), 0)::int AS total_amount
FROM
    months m
LEFT JOIN
    transactions t ON EXTRACT(MONTH FROM t.transaction_time) = EXTRACT(MONTH FROM m.month)
    AND EXTRACT(YEAR FROM t.transaction_time) = EXTRACT(YEAR FROM m.month)
    AND t.deleted_at IS NULL
LEFT JOIN
    merchants mch ON t.merchant_id = mch.merchant_id
    AND mch.deleted_at IS NULL
GROUP BY
    m.month
ORDER BY
    m.month;


-- name: GetYearlyAmountMerchant :many
WITH last_five_years AS (
    SELECT
        EXTRACT(YEAR FROM t.transaction_time) AS year,
        SUM(t.amount) AS total_amount
    FROM
        transactions t
    JOIN
        merchants m ON t.merchant_id = m.merchant_id
    WHERE
        t.deleted_at IS NULL AND m.deleted_at IS NULL
        AND EXTRACT(YEAR FROM t.transaction_time) >= $1 - 4
        AND EXTRACT(YEAR FROM t.transaction_time) <= $1
    GROUP BY
        EXTRACT(YEAR FROM t.created_at)
)
SELECT
    year,
    total_amount
FROM
    last_five_years
ORDER BY
    year;


-- name: FindAllTransactions :many
SELECT
    t.transaction_id,
    t.card_number,
    t.amount,
    t.payment_method,
    t.merchant_id,
    m.name AS merchant_name,
    t.transaction_time,
    t.created_at,
    t.updated_at,
    t.deleted_at,
    COUNT(*) OVER() AS total_count
FROM
    transactions t
JOIN
    merchants m ON t.merchant_id = m.merchant_id
WHERE
    t.deleted_at IS NULL
ORDER BY
    t.transaction_time DESC;




-- name: GetMonthlyPaymentMethodByMerchants :many
WITH months AS (
    SELECT generate_series(
        date_trunc('year', $1::timestamp),
        date_trunc('year', $1::timestamp) + interval '1 year' - interval '1 day',
        interval '1 month'
    ) AS month
),
payment_methods AS (
    SELECT DISTINCT payment_method
    FROM transactions
    WHERE deleted_at IS NULL
)
SELECT
    TO_CHAR(m.month, 'Mon') AS month,
    pm.payment_method,
    COALESCE(SUM(t.amount), 0)::int AS total_amount
FROM
    months m
CROSS JOIN
    payment_methods pm
LEFT JOIN
    transactions t ON EXTRACT(MONTH FROM t.transaction_time) = EXTRACT(MONTH FROM m.month)
    AND EXTRACT(YEAR FROM t.transaction_time) = EXTRACT(YEAR FROM m.month)
    AND t.payment_method = pm.payment_method
    AND t.deleted_at IS NULL
    AND t.merchant_id = $2
LEFT JOIN
    merchants mch ON t.merchant_id = mch.merchant_id
    AND mch.deleted_at IS NULL
GROUP BY
    m.month,
    pm.payment_method
ORDER BY
    m.month,
    pm.payment_method;


-- name: GetYearlyPaymentMethodByMerchants :many
WITH last_five_years AS (
    SELECT
        EXTRACT(YEAR FROM t.transaction_time) AS year,
        t.payment_method,
        SUM(t.amount) AS total_amount
    FROM
        transactions t
    JOIN
        merchants m ON t.merchant_id = m.merchant_id
    WHERE
        t.deleted_at IS NULL
        AND m.deleted_at IS NULL
        AND t.merchant_id = $1
        AND EXTRACT(YEAR FROM t.transaction_time) >= $2 - 4
        AND EXTRACT(YEAR FROM t.transaction_time) <= $2
    GROUP BY
        EXTRACT(YEAR FROM t.transaction_time),
        t.payment_method
)
SELECT
    year,
    payment_method,
    total_amount
FROM
    last_five_years
ORDER BY
    year;


-- name: GetMonthlyAmountByMerchants :many
WITH months AS (
    SELECT generate_series(
        date_trunc('year', $1::timestamp),
        date_trunc('year', $1::timestamp) + interval '1 year' - interval '1 day',
        interval '1 month'
    ) AS month
)
SELECT
    TO_CHAR(m.month, 'Mon') AS month,
    COALESCE(SUM(t.amount), 0)::int AS total_amount
FROM
    months m
LEFT JOIN
    transactions t ON EXTRACT(MONTH FROM t.transaction_time) = EXTRACT(MONTH FROM m.month)
    AND EXTRACT(YEAR FROM t.transaction_time) = EXTRACT(YEAR FROM m.month)
    AND t.deleted_at IS NULL
    AND t.merchant_id = $2
LEFT JOIN
    merchants mch ON t.merchant_id = mch.merchant_id
    AND mch.deleted_at IS NULL
GROUP BY
    m.month
ORDER BY
    m.month;



-- name: GetYearlyAmountByMerchants :many
WITH last_five_years AS (
    SELECT
        EXTRACT(YEAR FROM t.transaction_time) AS year,
        SUM(t.amount) AS total_amount
    FROM
        transactions t
    JOIN
        merchants m ON t.merchant_id = m.merchant_id
    WHERE
        t.deleted_at IS NULL
        AND m.deleted_at IS NULL
        AND t.merchant_id = $1
        AND EXTRACT(YEAR FROM t.transaction_time) >= $2 - 4
        AND EXTRACT(YEAR FROM t.transaction_time) <= $2
    GROUP BY
        EXTRACT(YEAR FROM t.transaction_time)
)
SELECT
    year,
    total_amount
FROM
    last_five_years
ORDER BY
    year;


-- name: FindAllTransactionsByMerchant :many
SELECT
    t.transaction_id,
    t.card_number,
    t.amount,
    t.payment_method,
    t.merchant_id,
    m.name AS merchant_name,
    t.transaction_time,
    t.created_at,
    t.updated_at,
    t.deleted_at,
    COUNT(*) OVER() AS total_count
FROM
    transactions t
JOIN
    merchants m ON t.merchant_id = m.merchant_id
WHERE
    (t.merchant_id = $1 OR $1 IS NOT NULL)
    AND t.deleted_at IS NULL
ORDER BY
    t.transaction_time DESC;


-- Create Merchant
-- name: CreateMerchant :one
INSERT INTO
    merchants (
        name,
        api_key,
        user_id,
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



-- Update Merchant
-- name: UpdateMerchant :exec
UPDATE merchants
SET
    name = $2,
    user_id = $3,
    updated_at = current_timestamp
WHERE
    merchant_id = $1
    AND deleted_at IS NULL;

-- UpdateMerchantStatus
-- name: UpdateMerchantStatus :exec
UPDATE merchants
SET
    status = $2,
    updated_at = current_timestamp
WHERE
    merchant_id = $1
    AND deleted_at IS NULL;



-- Trash Merchant
-- name: TrashMerchant :exec
UPDATE merchants
SET
    deleted_at = current_timestamp
WHERE
    merchant_id = $1
    AND deleted_at IS NULL;

-- Restore Trashed Merchant
-- name: RestoreMerchant :exec
UPDATE merchants
SET
    deleted_at = NULL
WHERE
    merchant_id = $1
    AND deleted_at IS NOT NULL;


-- Delete Merchant Permanently
-- name: DeleteMerchantPermanently :exec
DELETE FROM merchants WHERE merchant_id = $1 AND deleted_at IS NOT NULL;


-- Restore All Trashed Merchants
-- name: RestoreAllMerchants :exec
UPDATE merchants
SET
    deleted_at = NULL
WHERE
    deleted_at IS NOT NULL;

-- Delete All Trashed Merchants Permanently
-- name: DeleteAllPermanentMerchants :exec
DELETE FROM merchants
WHERE
    deleted_at IS NOT NULL;
