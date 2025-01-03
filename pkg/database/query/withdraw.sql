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



-- name: FindAllWithdrawsByCardNumber :many
SELECT
    w.withdraw_id,
    w.card_number,
    w.withdraw_amount,
    w.withdraw_time,
    w.created_at,
    w.updated_at,
    w.deleted_at
FROM
    withdraws w
WHERE
    w.card_number = $1 
    AND w.deleted_at IS NULL
ORDER BY
    w.withdraw_time DESC;


-- name: GetMonthlyWithdrawsByCardNumber :many
SELECT
    TO_CHAR(w.withdraw_time, 'Mon') AS month,
    SUM(w.withdraw_amount) AS total_withdraw_amount
FROM
    withdraws w
WHERE
    w.card_number = $1 
    AND w.deleted_at IS NULL
    AND EXTRACT(YEAR FROM w.withdraw_time) = $2 
GROUP BY
    TO_CHAR(w.withdraw_time, 'Mon'),
    EXTRACT(MONTH FROM w.withdraw_time)
ORDER BY
    EXTRACT(MONTH FROM w.withdraw_time);



-- name: GetMonthlyWithdrawsAll :many
SELECT
    TO_CHAR(w.withdraw_time, 'Mon') AS month,
    SUM(w.withdraw_amount) AS total_withdraw_amount
FROM
    withdraws w
WHERE
    w.deleted_at IS NULL
    AND EXTRACT(YEAR FROM w.withdraw_time) = $1 
GROUP BY
    TO_CHAR(w.withdraw_time, 'Mon'),
    EXTRACT(MONTH FROM w.withdraw_time)
ORDER BY
    EXTRACT(MONTH FROM w.withdraw_time);




-- name: GetYearlyWithdrawsAll :many
SELECT
    EXTRACT(YEAR FROM w.withdraw_time) AS year,
    SUM(w.withdraw_amount) AS total_withdraw_amount
FROM
    withdraws w
WHERE
    w.deleted_at IS NULL
GROUP BY
    EXTRACT(YEAR FROM w.withdraw_time)
ORDER BY
    year;
