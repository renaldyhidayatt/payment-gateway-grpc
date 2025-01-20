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



-- name: GetTotalBalance :one
SELECT
    SUM(s.total_balance) AS total_balance
FROM
    saldos s
JOIN
    cards c ON s.card_number = c.card_number
WHERE
    s.deleted_at IS NULL AND c.deleted_at IS NULL;


-- name: GetTotalTopupAmount :one
SELECT
    SUM(t.topup_amount) AS total_topup_amount
FROM
    topups t
JOIN
    cards c ON t.card_number = c.card_number
WHERE
    t.deleted_at IS NULL AND c.deleted_at IS NULL;


-- name: GetTotalWithdrawAmount :one
SELECT
    SUM(s.withdraw_amount) AS total_withdraw_amount
FROM
    saldos s
JOIN
    cards c ON s.card_number = c.card_number
WHERE
    s.deleted_at IS NULL AND c.deleted_at IS NULL;


-- name: GetTotalTransactionAmount :one
SELECT
    SUM(t.amount) AS total_transaction_amount
FROM
    transactions t
JOIN
    cards c ON t.card_number = c.card_number
WHERE
    t.deleted_at IS NULL AND c.deleted_at IS NULL;


-- name: GetTotalTransferAmount :one
SELECT
    SUM(transfer_amount) AS total_transfer_amount
FROM (
    SELECT
        transfer_amount
    FROM
        transfers
    WHERE
        deleted_at IS NULL
    UNION ALL
    SELECT
        transfer_amount
    FROM
        transfers
    WHERE
        deleted_at IS NULL
) AS transfer_data;





-- name: GetTotalBalanceByCardNumber :one
SELECT
    SUM(s.total_balance) AS total_balance
FROM
    saldos s
JOIN
    cards c ON s.card_number = c.card_number
WHERE
    s.deleted_at IS NULL AND c.deleted_at IS NULL
    AND c.card_number = $1;


-- name: GetTotalTopupAmountByCardNumber :one
SELECT
    SUM(t.topup_amount) AS total_topup_amount
FROM
    topups t
JOIN
    cards c ON t.card_number = c.card_number
WHERE
    t.deleted_at IS NULL AND c.deleted_at IS NULL
    AND c.card_number = $1;


-- name: GetTotalWithdrawAmountByCardNumber :one
SELECT
    SUM(s.withdraw_amount) AS total_withdraw_amount
FROM
    saldos s
JOIN
    cards c ON s.card_number = c.card_number
WHERE
    s.deleted_at IS NULL AND c.deleted_at IS NULL
    AND c.card_number = $1;


-- name: GetTotalTransactionAmountByCardNumber :one
SELECT
    SUM(t.amount) AS total_transaction_amount
FROM
    transactions t
JOIN
    cards c ON t.card_number = c.card_number
WHERE
    t.deleted_at IS NULL AND c.deleted_at IS NULL
    AND c.card_number = $1;



-- name: GetTotalTransferAmountBySender :one
SELECT
    SUM(transfer_amount) AS total_transfer_amount
FROM
    transfers
WHERE
    transfer_from = $1
    AND deleted_at IS NULL;



-- name: GetTotalTransferAmountByReceiver :one
SELECT
    SUM(transfer_amount) AS total_transfer_amount
FROM
    transfers
WHERE
    transfer_to = $1
    AND deleted_at IS NULL;






-- name: GetMonthlyBalances :many
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
LEFT JOIN
    cards c ON s.card_number = c.card_number
    AND c.deleted_at IS NULL
GROUP BY
    m.month
ORDER BY
    m.month;



-- name: GetYearlyBalances :many
WITH last_five_years AS (
    SELECT
        EXTRACT(YEAR FROM s.created_at) AS year,
        SUM(s.total_balance) AS total_balance
    FROM
        saldos s
    JOIN
        cards c ON s.card_number = c.card_number
    WHERE
        s.deleted_at IS NULL AND c.deleted_at IS NULL
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


-- name: GetMonthlyTopupAmount :many
WITH months AS (
    SELECT generate_series(
        date_trunc('year', $1::timestamp),
        date_trunc('year', $1::timestamp) + interval '1 year' - interval '1 day',
        interval '1 month'
    ) AS month
)
SELECT
    TO_CHAR(m.month, 'Mon') AS month,
    COALESCE(SUM(t.topup_amount), 0)::int AS total_topup_amount
FROM
    months m
LEFT JOIN
    topups t ON EXTRACT(MONTH FROM t.topup_time) = EXTRACT(MONTH FROM m.month)
    AND EXTRACT(YEAR FROM t.topup_time) = EXTRACT(YEAR FROM m.month)
    AND t.deleted_at IS NULL
LEFT JOIN
    cards c ON t.card_number = c.card_number
    AND c.deleted_at IS NULL
GROUP BY
    m.month
ORDER BY
    m.month;


-- name: GetYearlyTopupAmount :many
WITH last_five_years AS (
    SELECT
        EXTRACT(YEAR FROM t.topup_time) AS year,
        SUM(t.topup_amount) AS total_topup_amount
    FROM
        topups t
    JOIN
        cards c ON t.card_number = c.card_number
    WHERE
        t.deleted_at IS NULL
        AND c.deleted_at IS NULL
        AND EXTRACT(YEAR FROM t.topup_time) >= $1 - 4
        AND EXTRACT(YEAR FROM t.topup_time) <= $1
    GROUP BY
        EXTRACT(YEAR FROM t.topup_time)
)
SELECT
    year,
    total_topup_amount
FROM
    last_five_years
ORDER BY
    year;


-- name: GetMonthlyWithdrawAmount :many
WITH months AS (
    SELECT generate_series(
        date_trunc('year', $1::timestamp),
        date_trunc('year', $1::timestamp) + interval '1 year' - interval '1 day',
        interval '1 month'
    ) AS month
)
SELECT
    TO_CHAR(m.month, 'Mon') AS month,
    COALESCE(SUM(w.withdraw_amount), 0)::int AS total_withdraw_amount
FROM
    months m
LEFT JOIN
    withdraws w ON EXTRACT(MONTH FROM w.withdraw_time) = EXTRACT(MONTH FROM m.month)
    AND EXTRACT(YEAR FROM w.withdraw_time) = EXTRACT(YEAR FROM m.month)
    AND w.deleted_at IS NULL
LEFT JOIN
    cards c ON w.card_number = c.card_number
    AND c.deleted_at IS NULL
GROUP BY
    m.month
ORDER BY
    m.month;




-- name: GetYearlyWithdrawAmount :many
WITH last_five_years AS (
    SELECT
        EXTRACT(YEAR FROM w.withdraw_time) AS year,
        SUM(w.withdraw_amount) AS total_withdraw_amount
    FROM
        withdraws w
    JOIN
        cards c ON w.card_number = c.card_number
    WHERE
        w.deleted_at IS NULL
        AND c.deleted_at IS NULL
        AND EXTRACT(YEAR FROM w.withdraw_time) >= $1 - 4
        AND EXTRACT(YEAR FROM w.withdraw_time) <= $1
    GROUP BY
        EXTRACT(YEAR FROM w.withdraw_time)
)
SELECT
    year,
    total_withdraw_amount
FROM
    last_five_years
ORDER BY
    year;


-- name: GetMonthlyTransactionAmount :many
WITH months AS (
    SELECT generate_series(
        date_trunc('year', $1::timestamp),
        date_trunc('year', $1::timestamp) + interval '1 year' - interval '1 day',
        interval '1 month'
    ) AS month
)
SELECT
    TO_CHAR(m.month, 'Mon') AS month,
    COALESCE(SUM(t.amount), 0)::int AS total_transaction_amount
FROM
    months m
LEFT JOIN
    transactions t ON EXTRACT(MONTH FROM t.transaction_time) = EXTRACT(MONTH FROM m.month)
    AND EXTRACT(YEAR FROM t.transaction_time) = EXTRACT(YEAR FROM m.month)
    AND t.deleted_at IS NULL
LEFT JOIN
    cards c ON t.card_number = c.card_number
    AND c.deleted_at IS NULL
GROUP BY
    m.month
ORDER BY
    m.month;



-- name: GetYearlyTransactionAmount :many
WITH last_five_years AS (
    SELECT
        EXTRACT(YEAR FROM t.transaction_time) AS year,
        SUM(t.amount) AS total_transaction_amount
    FROM
        transactions t
    JOIN
        cards c ON t.card_number = c.card_number
    WHERE
        t.deleted_at IS NULL
        AND c.deleted_at IS NULL
        AND EXTRACT(YEAR FROM t.transaction_time) >= $1 - 4
        AND EXTRACT(YEAR FROM t.transaction_time) <= $1
    GROUP BY
        EXTRACT(YEAR FROM t.transaction_time)
)
SELECT
    year,
    total_transaction_amount
FROM
    last_five_years
ORDER BY
    year;


-- name: GetMonthlyTransferAmountSender :many
WITH months AS (
    SELECT generate_series(
        date_trunc('year', $1::timestamp),
        date_trunc('year', $1::timestamp) + interval '1 year' - interval '1 day',
        interval '1 month'
    ) AS month
)
SELECT
    TO_CHAR(m.month, 'Mon') AS month,
    COALESCE(SUM(t.transfer_amount), 0)::int AS total_sent_amount
FROM
    months m
LEFT JOIN
    transfers t ON EXTRACT(MONTH FROM t.transfer_time) = EXTRACT(MONTH FROM m.month)
    AND EXTRACT(YEAR FROM t.transfer_time) = EXTRACT(YEAR FROM m.month)
    AND t.deleted_at IS NULL
GROUP BY
    m.month
ORDER BY
    m.month;


-- name: GetMonthlyTransferAmountReceiver :many
WITH months AS (
    SELECT generate_series(
        date_trunc('year', $1::timestamp),
        date_trunc('year', $1::timestamp) + interval '1 year' - interval '1 day',
        interval '1 month'
    ) AS month
)
SELECT
    TO_CHAR(m.month, 'Mon') AS month,
    COALESCE(SUM(t.transfer_amount), 0)::int AS total_received_amount
FROM
    months m
LEFT JOIN
    transfers t ON EXTRACT(MONTH FROM t.transfer_time) = EXTRACT(MONTH FROM m.month)
    AND EXTRACT(YEAR FROM t.transfer_time) = EXTRACT(YEAR FROM m.month)
    AND t.deleted_at IS NULL
GROUP BY
    m.month
ORDER BY
    m.month;


-- name: GetYearlyTransferAmountSender :many
WITH last_five_years AS (
    SELECT
        EXTRACT(YEAR FROM t.transfer_time) AS year,
        SUM(t.transfer_amount) AS total_sent_amount
    FROM
        transfers t
    WHERE
        t.deleted_at IS NULL
        AND EXTRACT(YEAR FROM t.transfer_time) >= $1 - 4
        AND EXTRACT(YEAR FROM t.transfer_time) <= $1
    GROUP BY
        EXTRACT(YEAR FROM t.transfer_time)
)
SELECT
    year,
    total_sent_amount
FROM
    last_five_years
ORDER BY
    year;

-- name: GetYearlyTransferAmountReceiver :many
WITH last_five_years AS (
    SELECT
        EXTRACT(YEAR FROM t.transfer_time) AS year,
        SUM(t.transfer_amount) AS total_received_amount
    FROM
        transfers t
    WHERE
        t.deleted_at IS NULL
        AND EXTRACT(YEAR FROM t.transfer_time) >= $1 - 4
        AND EXTRACT(YEAR FROM t.transfer_time) <= $1
    GROUP BY
        EXTRACT(YEAR FROM t.transfer_time)
)
SELECT
    year,
    total_received_amount
FROM
    last_five_years
ORDER BY
    year;




-- name: GetMonthlyBalancesByCardNumber :many
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
LEFT JOIN
    cards c ON s.card_number = c.card_number
    AND c.deleted_at IS NULL
    AND c.card_number = $2
GROUP BY
    m.month
ORDER BY
    m.month;



-- name: GetYearlyBalancesByCardNumber :many
WITH last_five_years AS (
    SELECT
        EXTRACT(YEAR FROM s.created_at) AS year,
        SUM(s.total_balance) AS total_balance
    FROM
        saldos s
    JOIN
        cards c ON s.card_number = c.card_number
    WHERE
        s.deleted_at IS NULL AND c.deleted_at IS NULL
        AND EXTRACT(YEAR FROM s.created_at) >= $1 - 4
        AND EXTRACT(YEAR FROM s.created_at) <= $1
        AND c.card_number = $2
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



-- name: GetMonthlyTopupAmountByCardNumber :many
WITH months AS (
    SELECT generate_series(
        date_trunc('year', $2::timestamp),
        date_trunc('year', $2::timestamp) + interval '1 year' - interval '1 day',
        interval '1 month'
    ) AS month
)
SELECT
    TO_CHAR(m.month, 'Mon') AS month,
    COALESCE(SUM(t.topup_amount), 0)::int AS total_topup_amount
FROM
    months m
LEFT JOIN
    topups t ON EXTRACT(MONTH FROM t.topup_time) = EXTRACT(MONTH FROM m.month)
    AND EXTRACT(YEAR FROM t.topup_time) = EXTRACT(YEAR FROM m.month)
    AND t.deleted_at IS NULL
LEFT JOIN
    cards c ON t.card_number = c.card_number
    AND c.deleted_at IS NULL
    AND t.card_number = $1
GROUP BY
    m.month
ORDER BY
    m.month;


-- name: GetYearlyTopupAmountByCardNumber :many
WITH last_five_years AS (
    SELECT
        EXTRACT(YEAR FROM t.topup_time) AS year,
        SUM(t.topup_amount) AS total_topup_amount
    FROM
        topups t
    JOIN
        cards c ON t.card_number = c.card_number
    WHERE
        t.deleted_at IS NULL
        AND c.deleted_at IS NULL
        AND t.card_number = $1
        AND EXTRACT(YEAR FROM t.topup_time) >= $2 - 4
        AND EXTRACT(YEAR FROM t.topup_time) <= $2
    GROUP BY
        EXTRACT(YEAR FROM t.topup_time)
)
SELECT
    year,
    total_topup_amount
FROM
    last_five_years
ORDER BY
    year;


-- name: GetMonthlyWithdrawAmountByCardNumber :many
WITH months AS (
    SELECT generate_series(
        date_trunc('year', $2::timestamp),
        date_trunc('year', $2::timestamp) + interval '1 year' - interval '1 day',
        interval '1 month'
    ) AS month
)
SELECT
    TO_CHAR(m.month, 'Mon') AS month,
    COALESCE(SUM(w.withdraw_amount), 0)::int AS total_withdraw_amount
FROM
    months m
LEFT JOIN
    withdraws w ON EXTRACT(MONTH FROM w.withdraw_time) = EXTRACT(MONTH FROM m.month)
    AND EXTRACT(YEAR FROM w.withdraw_time) = EXTRACT(YEAR FROM m.month)
    AND w.deleted_at IS NULL
LEFT JOIN
    cards c ON w.card_number = c.card_number
    AND c.deleted_at IS NULL
    AND w.card_number = $1
GROUP BY
    m.month
ORDER BY
    m.month;



-- name: GetYearlyWithdrawAmountByCardNumber :many
WITH last_five_years AS (
    SELECT
        EXTRACT(YEAR FROM w.withdraw_time) AS year,
        SUM(w.withdraw_amount) AS total_withdraw_amount
    FROM
        withdraws w
    JOIN
        cards c ON w.card_number = c.card_number
    WHERE
        w.deleted_at IS NULL
        AND c.deleted_at IS NULL
        AND w.card_number = $1
        AND EXTRACT(YEAR FROM w.withdraw_time) >= $2 - 4
        AND EXTRACT(YEAR FROM w.withdraw_time) <= $2
    GROUP BY
        EXTRACT(YEAR FROM w.withdraw_time)
)
SELECT
    year,
    total_withdraw_amount
FROM
    last_five_years
ORDER BY
    year;



-- name: GetMonthlyTransactionAmountByCardNumber :many
WITH months AS (
    SELECT generate_series(
        date_trunc('year', $2::timestamp),
        date_trunc('year', $2::timestamp) + interval '1 year' - interval '1 day',
        interval '1 month'
    ) AS month
)
SELECT
    TO_CHAR(m.month, 'Mon') AS month,
    COALESCE(SUM(t.amount), 0)::int AS total_transaction_amount
FROM
    months m
LEFT JOIN
    transactions t ON EXTRACT(MONTH FROM t.transaction_time) = EXTRACT(MONTH FROM m.month)
    AND EXTRACT(YEAR FROM t.transaction_time) = EXTRACT(YEAR FROM m.month)
    AND t.deleted_at IS NULL
LEFT JOIN
    cards c ON t.card_number = c.card_number
    AND c.deleted_at IS NULL
    AND t.card_number = $1
GROUP BY
    m.month
ORDER BY
    m.month;



-- name: GetYearlyTransactionAmountByCardNumber :many
WITH last_five_years AS (
    SELECT
        EXTRACT(YEAR FROM t.transaction_time) AS year,
        SUM(t.amount) AS total_transaction_amount
    FROM
        transactions t
    JOIN
        cards c ON t.card_number = c.card_number
    WHERE
        t.deleted_at IS NULL
        AND c.deleted_at IS NULL
        AND t.card_number = $1
        AND EXTRACT(YEAR FROM t.transaction_time) >= $2 - 4
        AND EXTRACT(YEAR FROM t.transaction_time) <= $2
    GROUP BY
        EXTRACT(YEAR FROM t.transaction_time)
)
SELECT
    year,
    total_transaction_amount
FROM
    last_five_years
ORDER BY
    year;


-- name: GetMonthlyTransferAmountBySender :many
WITH months AS (
    SELECT generate_series(
        date_trunc('year', $2::timestamp),
        date_trunc('year', $2::timestamp) + interval '1 year' - interval '1 day',
        interval '1 month'
    ) AS month
)
SELECT
    TO_CHAR(m.month, 'Mon') AS month,
    COALESCE(SUM(t.transfer_amount), 0)::int AS total_sent_amount
FROM
    months m
LEFT JOIN
    transfers t ON EXTRACT(MONTH FROM t.transfer_time) = EXTRACT(MONTH FROM m.month)
    AND EXTRACT(YEAR FROM t.transfer_time) = EXTRACT(YEAR FROM m.month)
    AND t.deleted_at IS NULL
    AND t.transfer_from = $1
GROUP BY
    m.month
ORDER BY
    m.month;


-- name: GetMonthlyTransferAmountByReceiver :many
WITH months AS (
    SELECT generate_series(
        date_trunc('year', $2::timestamp),
        date_trunc('year', $2::timestamp) + interval '1 year' - interval '1 day',
        interval '1 month'
    ) AS month
)
SELECT
    TO_CHAR(m.month, 'Mon') AS month,
    COALESCE(SUM(t.transfer_amount), 0)::int AS total_received_amount
FROM
    months m
LEFT JOIN
    transfers t ON EXTRACT(MONTH FROM t.transfer_time) = EXTRACT(MONTH FROM m.month)
    AND EXTRACT(YEAR FROM t.transfer_time) = EXTRACT(YEAR FROM m.month)
    AND t.deleted_at IS NULL
    AND t.transfer_to = $1
GROUP BY
    m.month
ORDER BY
    m.month;




-- name: GetYearlyTransferAmountBySender :many
WITH last_five_years AS (
    SELECT
        EXTRACT(YEAR FROM t.transfer_time) AS year,
        SUM(t.transfer_amount) AS total_sent_amount
    FROM
        transfers t
    WHERE
        t.deleted_at IS NULL
        AND t.transfer_from = $1
        AND EXTRACT(YEAR FROM t.transfer_time) >= $2 - 4
        AND EXTRACT(YEAR FROM t.transfer_time) <= $2
    GROUP BY
        EXTRACT(YEAR FROM t.transfer_time)
)
SELECT
    year,
    total_sent_amount
FROM
    last_five_years
ORDER BY
    year;


-- name: GetYearlyTransferAmountByReceiver :many
WITH last_five_years AS (
    SELECT
        EXTRACT(YEAR FROM t.transfer_time) AS year,
        SUM(t.transfer_amount) AS total_received_amount
    FROM
        transfers t
    WHERE
        t.deleted_at IS NULL
        AND t.transfer_to = $1
        AND EXTRACT(YEAR FROM t.transfer_time) >= $2 - 4
        AND EXTRACT(YEAR FROM t.transfer_time) <= $2
    GROUP BY
        EXTRACT(YEAR FROM t.transfer_time)
)
SELECT
    year,
    total_received_amount
FROM
    last_five_years
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
