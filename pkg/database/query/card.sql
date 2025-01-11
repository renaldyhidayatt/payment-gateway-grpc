-- Search Cards with Pagination and Total Count
-- name: GetCards :many
SELECT
    *,
    COUNT(*) OVER() AS total_count
FROM cards
WHERE deleted_at IS NULL
  AND ($1::TEXT IS NULL OR card_number ILIKE '%' || $1 || '%' OR card_type ILIKE '%' || $1 || '%' OR card_provider ILIKE '%' || $1 || '%')
ORDER BY card_id
LIMIT $2 OFFSET $3;


-- Get Card by ID
-- name: GetCardByID :one
SELECT * FROM cards WHERE card_id = $1 AND deleted_at IS NULL;


-- name: GetActiveCardsWithCount :many
SELECT
    *,
    COUNT(*) OVER() AS total_count
FROM cards
WHERE deleted_at IS NULL
  AND ($1::TEXT IS NULL OR card_number ILIKE '%' || $1 || '%' OR card_type ILIKE '%' || $1 || '%' OR card_provider ILIKE '%' || $1 || '%')
ORDER BY card_id
LIMIT $2 OFFSET $3;


-- name: GetTrashedCardsWithCount :many
SELECT
    *,
    COUNT(*) OVER() AS total_count
FROM cards
WHERE deleted_at IS NOT NULL
  AND ($1::TEXT IS NULL OR card_number ILIKE '%' || $1 || '%' OR card_type ILIKE '%' || $1 || '%' OR card_provider ILIKE '%' || $1 || '%')
ORDER BY card_id
LIMIT $2 OFFSET $3;


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


-- name: GetMonthlyBalances :many
SELECT
    TO_CHAR(s.created_at, 'Mon') AS month,
    SUM(s.total_balance) AS total_balance
FROM
    saldos s
JOIN
    cards c ON s.card_number = c.card_number
WHERE
    s.deleted_at IS NULL AND c.deleted_at IS NULL
    AND EXTRACT(YEAR FROM s.created_at) = $1
GROUP BY
    TO_CHAR(s.created_at, 'Mon'), EXTRACT(MONTH FROM s.created_at)
ORDER BY
    EXTRACT(MONTH FROM s.created_at);


-- name: GetYearlyBalances :many
SELECT
    EXTRACT(YEAR FROM s.created_at) AS year,
    SUM(s.total_balance) AS total_balance
FROM
    saldos s
JOIN
    cards c ON s.card_number = c.card_number
WHERE
    s.deleted_at IS NULL AND c.deleted_at IS NULL
GROUP BY
    EXTRACT(YEAR FROM s.created_at)
ORDER BY
    year;

-- name: GetAllBalances :many
SELECT
    c.card_number,
    SUM(s.total_balance) AS total_balance
FROM
    saldos s
JOIN
    cards c ON s.card_number = c.card_number
WHERE
    s.deleted_at IS NULL AND c.deleted_at IS NULL
GROUP BY
    c.card_number
ORDER BY
    total_balance DESC;


-- name: GetAllWithdrawAmount :many
SELECT
    c.card_number,
    SUM(s.withdraw_amount) AS total_withdraw_amount
FROM
    saldos s
JOIN
    cards c ON s.card_number = c.card_number
WHERE
    s.deleted_at IS NULL AND c.deleted_at IS NULL
GROUP BY
    c.card_number
ORDER BY
    total_withdraw_amount DESC;

-- name: GetAllTransactionAmount :many
SELECT
    t.card_number,
    SUM(t.amount) AS total_transaction_amount
FROM
    transactions t
JOIN
    cards c ON t.card_number = c.card_number
WHERE
    t.deleted_at IS NULL AND c.deleted_at IS NULL
GROUP BY
    t.card_number
ORDER BY
    total_transaction_amount DESC;


-- name: GetAllTransferAmount :many
SELECT
    card_number,
    SUM(transfer_amount) AS total_transfer_amount
FROM (
    SELECT
        transfer_from AS card_number,
        transfer_amount
    FROM
        transfers
    WHERE
        deleted_at IS NULL
    UNION ALL
    SELECT
        transfer_to AS card_number,
        transfer_amount
    FROM
        transfers
    WHERE
        deleted_at IS NULL
) AS transfer_data
GROUP BY
    card_number
ORDER BY
    total_transfer_amount DESC;


-- name: GetAllTopupAmount :many
SELECT
    t.card_number,
    SUM(t.topup_amount) AS total_topup_amount
FROM
    topups t
JOIN
    cards c ON t.card_number = c.card_number
WHERE
    t.deleted_at IS NULL AND c.deleted_at IS NULL
GROUP BY
    t.card_number
ORDER BY
    total_topup_amount DESC;


-- name: GetBalanceByCardNumber :one
SELECT
    c.card_number,
    SUM(s.total_balance) AS total_balance
FROM
    saldos s
JOIN
    cards c ON s.card_number = c.card_number
WHERE
    s.deleted_at IS NULL AND c.deleted_at IS NULL AND c.card_number = $1
GROUP BY
    c.card_number;

-- name: GetWithdrawAmountByCardNumber :one
SELECT
    c.card_number,
    SUM(s.withdraw_amount) AS total_withdraw_amount
FROM
    saldos s
JOIN
    cards c ON s.card_number = c.card_number
WHERE
    s.deleted_at IS NULL AND c.deleted_at IS NULL AND c.card_number = $1
GROUP BY
    c.card_number;

-- name: GetTransactionAmountByCardNumber :one
SELECT
    t.card_number,
    SUM(t.amount) AS total_transaction_amount
FROM
    transactions t
JOIN
    cards c ON t.card_number = c.card_number
WHERE
    t.deleted_at IS NULL AND c.deleted_at IS NULL AND t.card_number = $1
GROUP BY
    t.card_number;

-- name: GetTransferAmountByCardNumber :one
SELECT
    card_number,
    SUM(transfer_amount) AS total_transfer_amount
FROM (
    SELECT
        transfer_from AS card_number,
        transfer_amount
    FROM
        transfers
    WHERE
        deleted_at IS NULL AND transfer_from = $1
    UNION ALL
    SELECT
        transfer_to AS card_number,
        transfer_amount
    FROM
        transfers
    WHERE
        deleted_at IS NULL AND transfer_to = $1
) AS transfer_data
GROUP BY
    card_number;

-- name: GetTopupAmountByCardNumber :one
SELECT
    t.card_number,
    SUM(t.topup_amount) AS total_topup_amount
FROM
    topups t
JOIN
    cards c ON t.card_number = c.card_number
WHERE
    t.deleted_at IS NULL AND c.deleted_at IS NULL AND t.card_number = $1
GROUP BY
    t.card_number;

-- name: GetMonthlyBalancesByCardNumber :many
SELECT
    TO_CHAR(s.created_at, 'Mon') AS month,
    SUM(s.total_balance) AS total_balance
FROM
    saldos s
JOIN
    cards c ON s.card_number = c.card_number
WHERE
    s.deleted_at IS NULL
    AND c.deleted_at IS NULL
    AND s.card_number = $1
    AND EXTRACT(YEAR FROM s.created_at) = $2
GROUP BY
    TO_CHAR(s.created_at, 'Mon'),
    EXTRACT(MONTH FROM s.created_at)
ORDER BY
    EXTRACT(MONTH FROM s.created_at);


-- name: GetYearlyBalancesByCardNUmber :many
SELECT
    EXTRACT(YEAR FROM s.created_at) AS year,
    SUM(s.total_balance) AS total_balance
FROM
    saldos s
JOIN
    cards c ON s.card_number = c.card_number
WHERE
    s.deleted_at IS NULL
    AND c.deleted_at IS NULL
    AND s.card_number = $1
GROUP BY
    EXTRACT(YEAR FROM s.created_at)
ORDER BY
    year;




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

-- Delete Card Permanently
-- name: DeleteCardPermanently :exec
DELETE FROM cards WHERE card_id = $1 AND deleted_at IS NOT NULL;


-- Restore All Trashed Cards
-- name: RestoreAllCards :exec
UPDATE cards
SET
    deleted_at = NULL
WHERE
    deleted_at IS NOT NULL;


-- Delete All Trashed Cards Permanently
-- name: DeleteAllPermanentCards :exec
DELETE FROM cards
WHERE
    deleted_at IS NOT NULL;
