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

-- Get All Active Merchants
-- name: GetActiveMerchants :many
SELECT *
FROM merchants
WHERE
    deleted_at IS NULL
ORDER BY merchant_id;

-- Get Trashed Merchants
-- name: GetTrashedMerchants :many
SELECT *
FROM merchants
WHERE
    deleted_at IS NOT NULL
ORDER BY merchant_id;

-- Search Merchants with Pagination
-- name: GetMerchants :many
SELECT *
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