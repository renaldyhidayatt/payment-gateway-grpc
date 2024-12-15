-- Create Card
-- name: CreateCard :one
INSERT INTO
    cards (
        user_id,
        card_number,
        card_type,
        expire_date,
        cvv,
        card_provider,
        created_at,
        updated_at
    )
VALUES (
        $1,
        $2,
        $3,
        $4,
        $5,
        $6,
        current_timestamp,
        current_timestamp
    ) RETURNING *;

-- Get Card by ID
-- name: GetCardByID :one
SELECT * FROM cards WHERE card_id = $1 AND deleted_at IS NULL;

-- Get Active Cards
-- name: GetActiveCards :many
SELECT * FROM cards WHERE deleted_at IS NULL ORDER BY card_id;

-- Get Trashed Cards
-- name: GetTrashedCards :many
SELECT * FROM cards WHERE deleted_at IS NOT NULL ORDER BY card_id;

-- Search Cards with Pagination
-- name: GetCards :many
SELECT *
FROM cards
WHERE deleted_at IS NULL
  AND ($1::TEXT IS NULL OR card_number ILIKE '%' || $1 || '%' OR card_type ILIKE '%' || $1 || '%' OR card_provider ILIKE '%' || $1 || '%')
ORDER BY card_id
LIMIT $2 OFFSET $3;

-- Trash Card
-- name: TrashCard :exec
UPDATE cards
SET
    deleted_at = current_timestamp
WHERE
    card_id = $1
    AND deleted_at IS NULL;

-- Restore Trashed Card
-- name: RestoreCard :exec
UPDATE cards
SET
    deleted_at = NULL
WHERE
    card_id = $1
    AND deleted_at IS NOT NULL;

-- Update Card
-- name: UpdateCard :exec
UPDATE cards
SET
    card_type = $2,
    expire_date = $3,
    cvv = $4,
    card_provider = $5,
    updated_at = current_timestamp
WHERE
    card_id = $1
    AND deleted_at IS NULL;

-- Delete Card Permanently
-- name: DeleteCardPermanently :exec
DELETE FROM cards WHERE card_id = $1;

-- Get a single Card by User ID
-- name: GetCardByUserID :one
SELECT *
FROM cards
WHERE
    user_id = $1
    AND deleted_at IS NULL
LIMIT 1;

-- Get Card by Card Number
-- name: GetCardByCardNumber :one
SELECT * FROM cards WHERE card_number = $1 AND deleted_at IS NULL;

-- Get Trashed By Card ID
-- name: GetTrashedCardByID :one
SELECT * FROM cards WHERE card_id = $1 AND deleted_at IS NOT NULL;