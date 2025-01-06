// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: merchant.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createMerchant = `-- name: CreateMerchant :one
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
    ) RETURNING merchant_id, name, api_key, user_id, status, created_at, updated_at, deleted_at
`

type CreateMerchantParams struct {
	Name   string `json:"name"`
	ApiKey string `json:"api_key"`
	UserID int32  `json:"user_id"`
	Status string `json:"status"`
}

// Create Merchant
func (q *Queries) CreateMerchant(ctx context.Context, arg CreateMerchantParams) (*Merchant, error) {
	row := q.db.QueryRowContext(ctx, createMerchant,
		arg.Name,
		arg.ApiKey,
		arg.UserID,
		arg.Status,
	)
	var i Merchant
	err := row.Scan(
		&i.MerchantID,
		&i.Name,
		&i.ApiKey,
		&i.UserID,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return &i, err
}

const deleteMerchantPermanently = `-- name: DeleteMerchantPermanently :exec
DELETE FROM merchants WHERE merchant_id = $1
`

// Delete Merchant Permanently
func (q *Queries) DeleteMerchantPermanently(ctx context.Context, merchantID int32) error {
	_, err := q.db.ExecContext(ctx, deleteMerchantPermanently, merchantID)
	return err
}

const findAllTransactionsByMerchantID = `-- name: FindAllTransactionsByMerchantID :many
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
    t.transaction_time DESC
`

type FindAllTransactionsByMerchantIDRow struct {
	TransactionID   int32        `json:"transaction_id"`
	CardNumber      string       `json:"card_number"`
	Amount          int32        `json:"amount"`
	PaymentMethod   string       `json:"payment_method"`
	MerchantID      int32        `json:"merchant_id"`
	MerchantName    string       `json:"merchant_name"`
	TransactionTime time.Time    `json:"transaction_time"`
	CreatedAt       sql.NullTime `json:"created_at"`
	UpdatedAt       sql.NullTime `json:"updated_at"`
	DeletedAt       sql.NullTime `json:"deleted_at"`
}

func (q *Queries) FindAllTransactionsByMerchantID(ctx context.Context, merchantID int32) ([]*FindAllTransactionsByMerchantIDRow, error) {
	rows, err := q.db.QueryContext(ctx, findAllTransactionsByMerchantID, merchantID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*FindAllTransactionsByMerchantIDRow
	for rows.Next() {
		var i FindAllTransactionsByMerchantIDRow
		if err := rows.Scan(
			&i.TransactionID,
			&i.CardNumber,
			&i.Amount,
			&i.PaymentMethod,
			&i.MerchantID,
			&i.MerchantName,
			&i.TransactionTime,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getActiveMerchants = `-- name: GetActiveMerchants :many
SELECT
    merchant_id, name, api_key, user_id, status, created_at, updated_at, deleted_at,
    COUNT(*) OVER() AS total_count
FROM merchants
WHERE deleted_at IS NULL
    AND ($1::TEXT IS NULL OR name ILIKE '%' || $1 || '%' OR api_key ILIKE '%' || $1 || '%' OR status ILIKE '%' || $1 || '%')
ORDER BY merchant_id
LIMIT $2 OFFSET $3
`

type GetActiveMerchantsParams struct {
	Column1 string `json:"column_1"`
	Limit   int32  `json:"limit"`
	Offset  int32  `json:"offset"`
}

type GetActiveMerchantsRow struct {
	MerchantID int32        `json:"merchant_id"`
	Name       string       `json:"name"`
	ApiKey     string       `json:"api_key"`
	UserID     int32        `json:"user_id"`
	Status     string       `json:"status"`
	CreatedAt  sql.NullTime `json:"created_at"`
	UpdatedAt  sql.NullTime `json:"updated_at"`
	DeletedAt  sql.NullTime `json:"deleted_at"`
	TotalCount int64        `json:"total_count"`
}

func (q *Queries) GetActiveMerchants(ctx context.Context, arg GetActiveMerchantsParams) ([]*GetActiveMerchantsRow, error) {
	rows, err := q.db.QueryContext(ctx, getActiveMerchants, arg.Column1, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*GetActiveMerchantsRow
	for rows.Next() {
		var i GetActiveMerchantsRow
		if err := rows.Scan(
			&i.MerchantID,
			&i.Name,
			&i.ApiKey,
			&i.UserID,
			&i.Status,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
			&i.TotalCount,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getMerchantByApiKey = `-- name: GetMerchantByApiKey :one
SELECT merchant_id, name, api_key, user_id, status, created_at, updated_at, deleted_at FROM merchants WHERE api_key = $1 AND deleted_at IS NULL
`

// Get Merchant by API Key
func (q *Queries) GetMerchantByApiKey(ctx context.Context, apiKey string) (*Merchant, error) {
	row := q.db.QueryRowContext(ctx, getMerchantByApiKey, apiKey)
	var i Merchant
	err := row.Scan(
		&i.MerchantID,
		&i.Name,
		&i.ApiKey,
		&i.UserID,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return &i, err
}

const getMerchantByID = `-- name: GetMerchantByID :one
SELECT merchant_id, name, api_key, user_id, status, created_at, updated_at, deleted_at
FROM merchants
WHERE
    merchant_id = $1
    AND deleted_at IS NULL
`

// Get Merchant by ID
func (q *Queries) GetMerchantByID(ctx context.Context, merchantID int32) (*Merchant, error) {
	row := q.db.QueryRowContext(ctx, getMerchantByID, merchantID)
	var i Merchant
	err := row.Scan(
		&i.MerchantID,
		&i.Name,
		&i.ApiKey,
		&i.UserID,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return &i, err
}

const getMerchantByName = `-- name: GetMerchantByName :one
SELECT merchant_id, name, api_key, user_id, status, created_at, updated_at, deleted_at FROM merchants WHERE name = $1 AND deleted_at IS NULL
`

// Get Merchant by Name
func (q *Queries) GetMerchantByName(ctx context.Context, name string) (*Merchant, error) {
	row := q.db.QueryRowContext(ctx, getMerchantByName, name)
	var i Merchant
	err := row.Scan(
		&i.MerchantID,
		&i.Name,
		&i.ApiKey,
		&i.UserID,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return &i, err
}

const getMerchants = `-- name: GetMerchants :many
SELECT
    merchant_id, name, api_key, user_id, status, created_at, updated_at, deleted_at,
    COUNT(*) OVER() AS total_count
FROM merchants
WHERE deleted_at IS NULL
    AND ($1::TEXT IS NULL OR name ILIKE '%' || $1 || '%' OR api_key ILIKE '%' || $1 || '%' OR status ILIKE '%' || $1 || '%')
ORDER BY merchant_id
LIMIT $2 OFFSET $3
`

type GetMerchantsParams struct {
	Column1 string `json:"column_1"`
	Limit   int32  `json:"limit"`
	Offset  int32  `json:"offset"`
}

type GetMerchantsRow struct {
	MerchantID int32        `json:"merchant_id"`
	Name       string       `json:"name"`
	ApiKey     string       `json:"api_key"`
	UserID     int32        `json:"user_id"`
	Status     string       `json:"status"`
	CreatedAt  sql.NullTime `json:"created_at"`
	UpdatedAt  sql.NullTime `json:"updated_at"`
	DeletedAt  sql.NullTime `json:"deleted_at"`
	TotalCount int64        `json:"total_count"`
}

func (q *Queries) GetMerchants(ctx context.Context, arg GetMerchantsParams) ([]*GetMerchantsRow, error) {
	rows, err := q.db.QueryContext(ctx, getMerchants, arg.Column1, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*GetMerchantsRow
	for rows.Next() {
		var i GetMerchantsRow
		if err := rows.Scan(
			&i.MerchantID,
			&i.Name,
			&i.ApiKey,
			&i.UserID,
			&i.Status,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
			&i.TotalCount,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getMerchantsByUserID = `-- name: GetMerchantsByUserID :many
SELECT merchant_id, name, api_key, user_id, status, created_at, updated_at, deleted_at FROM merchants WHERE user_id = $1 AND deleted_at IS NULL
`

// Get Merchants by User ID
func (q *Queries) GetMerchantsByUserID(ctx context.Context, userID int32) ([]*Merchant, error) {
	rows, err := q.db.QueryContext(ctx, getMerchantsByUserID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*Merchant
	for rows.Next() {
		var i Merchant
		if err := rows.Scan(
			&i.MerchantID,
			&i.Name,
			&i.ApiKey,
			&i.UserID,
			&i.Status,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getMonthlyAmountMerchant = `-- name: GetMonthlyAmountMerchant :many
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
    EXTRACT(MONTH FROM t.transaction_time)
`

type GetMonthlyAmountMerchantRow struct {
	Month       string `json:"month"`
	TotalAmount int64  `json:"total_amount"`
}

func (q *Queries) GetMonthlyAmountMerchant(ctx context.Context, transactionTime time.Time) ([]*GetMonthlyAmountMerchantRow, error) {
	rows, err := q.db.QueryContext(ctx, getMonthlyAmountMerchant, transactionTime)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*GetMonthlyAmountMerchantRow
	for rows.Next() {
		var i GetMonthlyAmountMerchantRow
		if err := rows.Scan(&i.Month, &i.TotalAmount); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getMonthlyAmountsByMerchant = `-- name: GetMonthlyAmountsByMerchant :many
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
    EXTRACT(MONTH FROM t.transaction_time)
`

type GetMonthlyAmountsByMerchantParams struct {
	MerchantID      int32     `json:"merchant_id"`
	TransactionTime time.Time `json:"transaction_time"`
}

type GetMonthlyAmountsByMerchantRow struct {
	Month       string `json:"month"`
	TotalAmount int64  `json:"total_amount"`
}

func (q *Queries) GetMonthlyAmountsByMerchant(ctx context.Context, arg GetMonthlyAmountsByMerchantParams) ([]*GetMonthlyAmountsByMerchantRow, error) {
	rows, err := q.db.QueryContext(ctx, getMonthlyAmountsByMerchant, arg.MerchantID, arg.TransactionTime)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*GetMonthlyAmountsByMerchantRow
	for rows.Next() {
		var i GetMonthlyAmountsByMerchantRow
		if err := rows.Scan(&i.Month, &i.TotalAmount); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getMonthlyPaymentMethodsMerchant = `-- name: GetMonthlyPaymentMethodsMerchant :many
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
    EXTRACT(MONTH FROM t.transaction_time)
`

type GetMonthlyPaymentMethodsMerchantRow struct {
	Month         string `json:"month"`
	PaymentMethod string `json:"payment_method"`
	TotalAmount   int64  `json:"total_amount"`
}

func (q *Queries) GetMonthlyPaymentMethodsMerchant(ctx context.Context, transactionTime time.Time) ([]*GetMonthlyPaymentMethodsMerchantRow, error) {
	rows, err := q.db.QueryContext(ctx, getMonthlyPaymentMethodsMerchant, transactionTime)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*GetMonthlyPaymentMethodsMerchantRow
	for rows.Next() {
		var i GetMonthlyPaymentMethodsMerchantRow
		if err := rows.Scan(&i.Month, &i.PaymentMethod, &i.TotalAmount); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTrashedMerchantByID = `-- name: GetTrashedMerchantByID :one
SELECT merchant_id, name, api_key, user_id, status, created_at, updated_at, deleted_at
FROM merchants
WHERE
    merchant_id = $1
    AND deleted_at IS NOT NULL
`

// Get Trashed By Merchant ID
func (q *Queries) GetTrashedMerchantByID(ctx context.Context, merchantID int32) (*Merchant, error) {
	row := q.db.QueryRowContext(ctx, getTrashedMerchantByID, merchantID)
	var i Merchant
	err := row.Scan(
		&i.MerchantID,
		&i.Name,
		&i.ApiKey,
		&i.UserID,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return &i, err
}

const getTrashedMerchants = `-- name: GetTrashedMerchants :many
SELECT
    merchant_id, name, api_key, user_id, status, created_at, updated_at, deleted_at,
    COUNT(*) OVER() AS total_count
FROM merchants
WHERE deleted_at IS NOT NULL
    AND ($1::TEXT IS NULL OR name ILIKE '%' || $1 || '%' OR api_key ILIKE '%' || $1 || '%' OR status ILIKE '%' || $1 || '%')
ORDER BY merchant_id
LIMIT $2 OFFSET $3
`

type GetTrashedMerchantsParams struct {
	Column1 string `json:"column_1"`
	Limit   int32  `json:"limit"`
	Offset  int32  `json:"offset"`
}

type GetTrashedMerchantsRow struct {
	MerchantID int32        `json:"merchant_id"`
	Name       string       `json:"name"`
	ApiKey     string       `json:"api_key"`
	UserID     int32        `json:"user_id"`
	Status     string       `json:"status"`
	CreatedAt  sql.NullTime `json:"created_at"`
	UpdatedAt  sql.NullTime `json:"updated_at"`
	DeletedAt  sql.NullTime `json:"deleted_at"`
	TotalCount int64        `json:"total_count"`
}

func (q *Queries) GetTrashedMerchants(ctx context.Context, arg GetTrashedMerchantsParams) ([]*GetTrashedMerchantsRow, error) {
	rows, err := q.db.QueryContext(ctx, getTrashedMerchants, arg.Column1, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*GetTrashedMerchantsRow
	for rows.Next() {
		var i GetTrashedMerchantsRow
		if err := rows.Scan(
			&i.MerchantID,
			&i.Name,
			&i.ApiKey,
			&i.UserID,
			&i.Status,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
			&i.TotalCount,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getYearlyAmountMerchant = `-- name: GetYearlyAmountMerchant :many
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
    year
`

type GetYearlyAmountMerchantRow struct {
	Year        string `json:"year"`
	TotalAmount int64  `json:"total_amount"`
}

func (q *Queries) GetYearlyAmountMerchant(ctx context.Context) ([]*GetYearlyAmountMerchantRow, error) {
	rows, err := q.db.QueryContext(ctx, getYearlyAmountMerchant)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*GetYearlyAmountMerchantRow
	for rows.Next() {
		var i GetYearlyAmountMerchantRow
		if err := rows.Scan(&i.Year, &i.TotalAmount); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getYearlyAmountsByMerchant = `-- name: GetYearlyAmountsByMerchant :many
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
    year
`

type GetYearlyAmountsByMerchantRow struct {
	Year        string `json:"year"`
	TotalAmount int64  `json:"total_amount"`
}

func (q *Queries) GetYearlyAmountsByMerchant(ctx context.Context, merchantID int32) ([]*GetYearlyAmountsByMerchantRow, error) {
	rows, err := q.db.QueryContext(ctx, getYearlyAmountsByMerchant, merchantID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*GetYearlyAmountsByMerchantRow
	for rows.Next() {
		var i GetYearlyAmountsByMerchantRow
		if err := rows.Scan(&i.Year, &i.TotalAmount); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getYearlyPaymentMethodMerchant = `-- name: GetYearlyPaymentMethodMerchant :many
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
    year
`

type GetYearlyPaymentMethodMerchantRow struct {
	Year          string `json:"year"`
	PaymentMethod string `json:"payment_method"`
	TotalAmount   int64  `json:"total_amount"`
}

func (q *Queries) GetYearlyPaymentMethodMerchant(ctx context.Context) ([]*GetYearlyPaymentMethodMerchantRow, error) {
	rows, err := q.db.QueryContext(ctx, getYearlyPaymentMethodMerchant)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*GetYearlyPaymentMethodMerchantRow
	for rows.Next() {
		var i GetYearlyPaymentMethodMerchantRow
		if err := rows.Scan(&i.Year, &i.PaymentMethod, &i.TotalAmount); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const restoreMerchant = `-- name: RestoreMerchant :exec
UPDATE merchants
SET
    deleted_at = NULL
WHERE
    merchant_id = $1
    AND deleted_at IS NOT NULL
`

// Restore Trashed Merchant
func (q *Queries) RestoreMerchant(ctx context.Context, merchantID int32) error {
	_, err := q.db.ExecContext(ctx, restoreMerchant, merchantID)
	return err
}

const trashMerchant = `-- name: TrashMerchant :exec
UPDATE merchants
SET
    deleted_at = current_timestamp
WHERE
    merchant_id = $1
    AND deleted_at IS NULL
`

// Trash Merchant
func (q *Queries) TrashMerchant(ctx context.Context, merchantID int32) error {
	_, err := q.db.ExecContext(ctx, trashMerchant, merchantID)
	return err
}

const updateMerchant = `-- name: UpdateMerchant :exec
UPDATE merchants
SET
    name = $2,
    user_id = $3,
    status = $4,
    updated_at = current_timestamp
WHERE
    merchant_id = $1
    AND deleted_at IS NULL
`

type UpdateMerchantParams struct {
	MerchantID int32  `json:"merchant_id"`
	Name       string `json:"name"`
	UserID     int32  `json:"user_id"`
	Status     string `json:"status"`
}

// Update Merchant
func (q *Queries) UpdateMerchant(ctx context.Context, arg UpdateMerchantParams) error {
	_, err := q.db.ExecContext(ctx, updateMerchant,
		arg.MerchantID,
		arg.Name,
		arg.UserID,
		arg.Status,
	)
	return err
}
