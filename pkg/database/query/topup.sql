-- Create Topup
-- name: CreateTopup :one
INSERT INTO
    topups (
        card_number,
        topup_no,
        topup_amount,
        topup_method,
        topup_time,
        created_at,
        updated_at
    )
VALUES (
        $1,
        $2,
        $3,
        $4,
        current_timestamp,
        current_timestamp,
        current_timestamp
    ) RETURNING *;

-- Get Topup by ID
-- name: GetTopupByID :one
SELECT * FROM topups WHERE topup_id = $1 AND deleted_at IS NULL;


-- Get All Active Topups with Pagination and Search
-- name: GetActiveTopups :many
SELECT
    *,
    COUNT(*) OVER() AS total_count
FROM
    topups
WHERE
    deleted_at IS NULL
    AND ($1::TEXT IS NULL OR card_number ILIKE '%' || $1 || '%' OR topup_no ILIKE '%' || $1 || '%' OR topup_method ILIKE '%' || $1 || '%')
ORDER BY
    topup_time DESC
LIMIT $2 OFFSET $3;

-- Get Trashed Topups with Pagination and Search
-- name: GetTrashedTopups :many
SELECT
    *,
    COUNT(*) OVER() AS total_count
FROM
    topups
WHERE
    deleted_at IS NOT NULL
    AND ($1::TEXT IS NULL OR card_number ILIKE '%' || $1 || '%' OR topup_no ILIKE '%' || $1 || '%' OR topup_method ILIKE '%' || $1 || '%')
ORDER BY
    topup_time DESC
LIMIT $2 OFFSET $3;


-- Search Topups with Pagination
-- name: GetTopups :many
SELECT
    *,
    COUNT(*) OVER() AS total_count
FROM
    topups
WHERE
    deleted_at IS NULL
    AND ($1::TEXT IS NULL OR card_number ILIKE '%' || $1 || '%' OR topup_no ILIKE '%' || $1 || '%' OR topup_method ILIKE '%' || $1 || '%')
ORDER BY
    topup_time DESC
LIMIT $2 OFFSET $3;


-- Count Topups by Date
-- name: CountTopupsByDate :one
SELECT COUNT(*)
FROM topups
WHERE deleted_at IS NULL
  AND topup_time::DATE = $1::DATE;

-- Count All Topups
-- name: CountAllTopups :one
SELECT COUNT(*) FROM topups WHERE deleted_at IS NULL;

-- Trash Topup
-- name: TrashTopup :exec
UPDATE topups
SET
    deleted_at = current_timestamp
WHERE
    topup_id = $1
    AND deleted_at IS NULL;

-- Restore Trashed Topup
-- name: RestoreTopup :exec
UPDATE topups
SET
    deleted_at = NULL
WHERE
    topup_id = $1
    AND deleted_at IS NOT NULL;

-- Update Topup
-- name: UpdateTopup :exec
UPDATE topups
SET
    card_number = $2,
    topup_amount = $3,
    topup_method = $4,
    topup_time = current_timestamp,
    updated_at = current_timestamp
WHERE
    topup_id = $1
    AND deleted_at IS NULL;

-- Update Topup Amount
-- name: UpdateTopupAmount :exec
UPDATE topups
SET
    topup_amount = $2,
    updated_at = current_timestamp
WHERE
    topup_id = $1
    AND deleted_at IS NULL;

-- Delete Topup Permanently
-- name: DeleteTopupPermanently :exec
DELETE FROM topups WHERE topup_id = $1;

-- Get Topups by Card Number
-- name: GetTopupsByCardNumber :many
SELECT *
FROM topups
WHERE
    deleted_at IS NULL
    AND card_number = $1
ORDER BY topup_time DESC;

-- Get Trashed By Topup ID
-- name: GetTrashedTopupByID :one
SELECT *
FROM topups
WHERE
    topup_id = $1
    AND deleted_at IS NOT NULL;


-- name: GetMonthlyTopupMethods :many
SELECT
    TO_CHAR(t.topup_time, 'Mon') AS month,
    t.topup_method,
    COUNT(t.topup_id) AS total_topups,
    SUM(t.topup_amount) AS total_amount
FROM
    topups t
WHERE
    t.deleted_at IS NULL
    AND EXTRACT(YEAR FROM t.topup_time) = $1
GROUP BY
    TO_CHAR(t.topup_time, 'Mon'),
    EXTRACT(MONTH FROM t.topup_time),
    t.topup_method
ORDER BY
    EXTRACT(MONTH FROM t.topup_time);


-- name: GetYearlyTopupMethods :many
SELECT
    EXTRACT(YEAR FROM t.topup_time) AS year,
    t.topup_method,
    COUNT(t.topup_id) AS total_topups,
    SUM(t.topup_amount) AS total_amount
FROM
    topups t
WHERE
    t.deleted_at IS NULL
GROUP BY
    EXTRACT(YEAR FROM t.topup_time),
    t.topup_method
ORDER BY
    year;


-- name: GetMonthlyTopupAmounts :many
SELECT
    TO_CHAR(t.topup_time, 'Mon') AS month,
    SUM(t.topup_amount) AS total_amount
FROM
    topups t
WHERE
    t.deleted_at IS NULL
    AND EXTRACT(YEAR FROM t.topup_time) = $1 -- Filter tahun (parameter $1)
GROUP BY
    TO_CHAR(t.topup_time, 'Mon'),
    EXTRACT(MONTH FROM t.topup_time)
ORDER BY
    EXTRACT(MONTH FROM t.topup_time);

-- name: GetYearlyTopupAmounts :many
SELECT
    EXTRACT(YEAR FROM t.topup_time) AS year,
    SUM(t.topup_amount) AS total_amount
FROM
    topups t
WHERE
    t.deleted_at IS NULL
GROUP BY
    EXTRACT(YEAR FROM t.topup_time)
ORDER BY
    year;


-- name: CountTopups :one
SELECT COUNT(*)
FROM topups
WHERE deleted_at IS NULL
    AND ($1::TEXT IS NULL OR
        card_number ILIKE '%' || $1 || '%' OR
        topup_method ILIKE '%' || $1 || '%' OR
        topup_status ILIKE '%' || $1 || '%');



-- name: Topup_CountAll :one
SELECT COUNT(*)
FROM topups
WHERE deleted_at IS NULL;
