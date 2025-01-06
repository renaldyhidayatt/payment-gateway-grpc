-- Create Merchant
-- name: CreateMerchant :one
INSERT INTO
    merchants (
        name,
        api_key,
        user_id,
        status,
        created_at,
        updated_at
    )
VALUES (
        $1,
        $2,
        $3,
        $4,
        current_timestamp,
        current_timestamp
    ) RETURNING *;

-- Get Merchant by ID
-- name: GetMerchantByID :one
SELECT *
FROM merchants
WHERE
    merchant_id = $1
    AND deleted_at IS NULL;


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

-- name: GetMerchants :many
SELECT
    *,
    COUNT(*) OVER() AS total_count
FROM merchants
WHERE deleted_at IS NULL
    AND ($1::TEXT IS NULL OR name ILIKE '%' || $1 || '%' OR api_key ILIKE '%' || $1 || '%' OR status ILIKE '%' || $1 || '%')
ORDER BY merchant_id
LIMIT $2 OFFSET $3;


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

-- Update Merchant
-- name: UpdateMerchant :exec
UPDATE merchants
SET
    name = $2,
    user_id = $3,
    status = $4,
    updated_at = current_timestamp
WHERE
    merchant_id = $1
    AND deleted_at IS NULL;

-- Delete Merchant Permanently
-- name: DeleteMerchantPermanently :exec
DELETE FROM merchants WHERE merchant_id = $1;

-- Get Merchant by API Key
-- name: GetMerchantByApiKey :one
SELECT * FROM merchants WHERE api_key = $1 AND deleted_at IS NULL;

-- Get Merchant by Name
-- name: GetMerchantByName :one
SELECT * FROM merchants WHERE name = $1 AND deleted_at IS NULL;

-- Get Merchants by User ID
-- name: GetMerchantsByUserID :many
SELECT * FROM merchants WHERE user_id = $1 AND deleted_at IS NULL;

-- Get Trashed By Merchant ID
-- name: GetTrashedMerchantByID :one
SELECT *
FROM merchants
WHERE
    merchant_id = $1
    AND deleted_at IS NOT NULL;


-- name: GetMonthlyPaymentMethodsMerchant :many
SELECT
    TO_CHAR(t.transaction_time, 'Mon') AS month,
    t.payment_method,
    SUM(t.amount) AS total_amount
FROM
    transactions t
JOIN
    merchants m ON t.merchant_id = m.merchant_id
WHERE
    t.deleted_at IS NULL AND m.deleted_at IS NULL
    AND EXTRACT(YEAR FROM t.transaction_time) = $1
GROUP BY
    TO_CHAR(t.transaction_time, 'Mon'),
    EXTRACT(MONTH FROM t.transaction_time),
    t.payment_method
ORDER BY
    EXTRACT(MONTH FROM t.transaction_time);


-- name: GetYearlyPaymentMethodMerchant :many
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
GROUP BY
    EXTRACT(YEAR FROM t.transaction_time),
    t.payment_method
ORDER BY
    year;


-- name: GetMonthlyAmountMerchant :many
SELECT
    TO_CHAR(t.transaction_time, 'Mon') AS month,
    SUM(t.amount) AS total_amount
FROM
    transactions t
JOIN
    merchants m ON t.merchant_id = m.merchant_id
WHERE
    t.deleted_at IS NULL AND m.deleted_at IS NULL
    AND EXTRACT(YEAR FROM t.transaction_time) = $1
GROUP BY
    TO_CHAR(t.transaction_time, 'Mon'),
    EXTRACT(MONTH FROM t.transaction_time)
ORDER BY
    EXTRACT(MONTH FROM t.transaction_time);

-- name: GetYearlyAmountMerchant :many
SELECT
    EXTRACT(YEAR FROM t.transaction_time) AS year,
    SUM(t.amount) AS total_amount
FROM
    transactions t
JOIN
    merchants m ON t.merchant_id = m.merchant_id
WHERE
    t.deleted_at IS NULL AND m.deleted_at IS NULL
GROUP BY
    EXTRACT(YEAR FROM t.transaction_time)
ORDER BY
    year;



-- name: GetMonthlyAmountsByMerchant :many
SELECT
    TO_CHAR(t.transaction_time, 'Mon') AS month,
    SUM(t.amount) AS total_amount
FROM
    transactions t
JOIN
    merchants m ON t.merchant_id = m.merchant_id
WHERE
    t.deleted_at IS NULL
    AND m.deleted_at IS NULL
    AND m.merchant_id = $1
    AND EXTRACT(YEAR FROM t.transaction_time) = $2
GROUP BY
    TO_CHAR(t.transaction_time, 'Mon'),
    EXTRACT(MONTH FROM t.transaction_time)
ORDER BY
    EXTRACT(MONTH FROM t.transaction_time);




-- name: GetYearlyAmountsByMerchant :many
SELECT
    EXTRACT(YEAR FROM t.transaction_time) AS year,
    SUM(t.amount) AS total_amount
FROM
    transactions t
JOIN
    merchants m ON t.merchant_id = m.merchant_id
WHERE
    t.deleted_at IS NULL AND m.deleted_at IS NULL
    AND m.merchant_id = $1 -- Ganti $1 dengan ID merchant
GROUP BY
    EXTRACT(YEAR FROM t.transaction_time)
ORDER BY
    year;


-- name: FindAllTransactionsByMerchantID :many
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
    t.deleted_at
FROM
    transactions t
JOIN
    merchants m ON t.merchant_id = m.merchant_id
WHERE
    (t.merchant_id = $1 OR $1 IS NULL)
    AND t.deleted_at IS NULL
ORDER BY
    t.transaction_time DESC;
