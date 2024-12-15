// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: withdraw.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const countActiveWithdrawsByDate = `-- name: CountActiveWithdrawsByDate :one
SELECT COUNT(*)
FROM withdraws
WHERE deleted_at IS NULL AND withdraw_time::DATE = $1
`

// Count Active Withdraws by Date
func (q *Queries) CountActiveWithdrawsByDate(ctx context.Context, withdrawTime time.Time) (int64, error) {
	row := q.db.QueryRowContext(ctx, countActiveWithdrawsByDate, withdrawTime)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createWithdraw = `-- name: CreateWithdraw :one
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
    ) RETURNING withdraw_id, card_number, withdraw_amount, withdraw_time, created_at, updated_at, deleted_at
`

type CreateWithdrawParams struct {
	CardNumber     string    `json:"card_number"`
	WithdrawAmount int32     `json:"withdraw_amount"`
	WithdrawTime   time.Time `json:"withdraw_time"`
}

// Create Withdraw
func (q *Queries) CreateWithdraw(ctx context.Context, arg CreateWithdrawParams) (*Withdraw, error) {
	row := q.db.QueryRowContext(ctx, createWithdraw, arg.CardNumber, arg.WithdrawAmount, arg.WithdrawTime)
	var i Withdraw
	err := row.Scan(
		&i.WithdrawID,
		&i.CardNumber,
		&i.WithdrawAmount,
		&i.WithdrawTime,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return &i, err
}

const deleteWithdrawPermanently = `-- name: DeleteWithdrawPermanently :exec
DELETE FROM withdraws WHERE withdraw_id = $1
`

// Delete Withdraw Permanently
func (q *Queries) DeleteWithdrawPermanently(ctx context.Context, withdrawID int32) error {
	_, err := q.db.ExecContext(ctx, deleteWithdrawPermanently, withdrawID)
	return err
}

const getActiveWithdraws = `-- name: GetActiveWithdraws :many
SELECT withdraw_id, card_number, withdraw_amount, withdraw_time, created_at, updated_at, deleted_at
FROM withdraws
WHERE
    deleted_at IS NULL
ORDER BY withdraw_time DESC
`

// Get All Active Withdraws
func (q *Queries) GetActiveWithdraws(ctx context.Context) ([]*Withdraw, error) {
	rows, err := q.db.QueryContext(ctx, getActiveWithdraws)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*Withdraw
	for rows.Next() {
		var i Withdraw
		if err := rows.Scan(
			&i.WithdrawID,
			&i.CardNumber,
			&i.WithdrawAmount,
			&i.WithdrawTime,
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

const getTrashedWithdrawByID = `-- name: GetTrashedWithdrawByID :one
SELECT withdraw_id, card_number, withdraw_amount, withdraw_time, created_at, updated_at, deleted_at
FROM withdraws
WHERE
    withdraw_id = $1
    AND deleted_at IS NOT NULL
`

// Get Trashed By Withdraw ID
func (q *Queries) GetTrashedWithdrawByID(ctx context.Context, withdrawID int32) (*Withdraw, error) {
	row := q.db.QueryRowContext(ctx, getTrashedWithdrawByID, withdrawID)
	var i Withdraw
	err := row.Scan(
		&i.WithdrawID,
		&i.CardNumber,
		&i.WithdrawAmount,
		&i.WithdrawTime,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return &i, err
}

const getTrashedWithdraws = `-- name: GetTrashedWithdraws :many
SELECT withdraw_id, card_number, withdraw_amount, withdraw_time, created_at, updated_at, deleted_at
FROM withdraws
WHERE
    deleted_at IS NOT NULL
ORDER BY withdraw_time DESC
`

// Get Trashed Withdraws
func (q *Queries) GetTrashedWithdraws(ctx context.Context) ([]*Withdraw, error) {
	rows, err := q.db.QueryContext(ctx, getTrashedWithdraws)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*Withdraw
	for rows.Next() {
		var i Withdraw
		if err := rows.Scan(
			&i.WithdrawID,
			&i.CardNumber,
			&i.WithdrawAmount,
			&i.WithdrawTime,
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

const getWithdrawByID = `-- name: GetWithdrawByID :one
SELECT withdraw_id, card_number, withdraw_amount, withdraw_time, created_at, updated_at, deleted_at
FROM withdraws
WHERE
    withdraw_id = $1
    AND deleted_at IS NULL
`

// Get Withdraw by ID
func (q *Queries) GetWithdrawByID(ctx context.Context, withdrawID int32) (*Withdraw, error) {
	row := q.db.QueryRowContext(ctx, getWithdrawByID, withdrawID)
	var i Withdraw
	err := row.Scan(
		&i.WithdrawID,
		&i.CardNumber,
		&i.WithdrawAmount,
		&i.WithdrawTime,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return &i, err
}

const restoreWithdraw = `-- name: RestoreWithdraw :exec
UPDATE withdraws
SET
    deleted_at = NULL
WHERE
    withdraw_id = $1
    AND deleted_at IS NOT NULL
`

// Restore Withdraw (Undelete)
func (q *Queries) RestoreWithdraw(ctx context.Context, withdrawID int32) error {
	_, err := q.db.ExecContext(ctx, restoreWithdraw, withdrawID)
	return err
}

const searchWithdrawByCardNumber = `-- name: SearchWithdrawByCardNumber :many
SELECT withdraw_id, card_number, withdraw_amount, withdraw_time, created_at, updated_at, deleted_at
FROM withdraws
WHERE
    deleted_at IS NULL
    AND card_number ILIKE '%' || $1 || '%'
ORDER BY withdraw_time DESC
`

// Search Withdraw by Card Number
func (q *Queries) SearchWithdrawByCardNumber(ctx context.Context, dollar_1 sql.NullString) ([]*Withdraw, error) {
	rows, err := q.db.QueryContext(ctx, searchWithdrawByCardNumber, dollar_1)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*Withdraw
	for rows.Next() {
		var i Withdraw
		if err := rows.Scan(
			&i.WithdrawID,
			&i.CardNumber,
			&i.WithdrawAmount,
			&i.WithdrawTime,
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

const searchWithdraws = `-- name: SearchWithdraws :many
SELECT withdraw_id, card_number, withdraw_amount, withdraw_time, created_at, updated_at, deleted_at
FROM withdraws
WHERE deleted_at IS NULL
  AND ($1::TEXT IS NULL OR card_number ILIKE '%' || $1 || '%')
ORDER BY withdraw_time DESC
LIMIT $2 OFFSET $3
`

type SearchWithdrawsParams struct {
	Column1 string `json:"column_1"`
	Limit   int32  `json:"limit"`
	Offset  int32  `json:"offset"`
}

// Search Withdraws with Pagination
func (q *Queries) SearchWithdraws(ctx context.Context, arg SearchWithdrawsParams) ([]*Withdraw, error) {
	rows, err := q.db.QueryContext(ctx, searchWithdraws, arg.Column1, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*Withdraw
	for rows.Next() {
		var i Withdraw
		if err := rows.Scan(
			&i.WithdrawID,
			&i.CardNumber,
			&i.WithdrawAmount,
			&i.WithdrawTime,
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

const trashWithdraw = `-- name: TrashWithdraw :exec
UPDATE withdraws
SET
    deleted_at = current_timestamp
WHERE
    withdraw_id = $1
    AND deleted_at IS NULL
`

// Trash Withdraw (Soft Delete)
func (q *Queries) TrashWithdraw(ctx context.Context, withdrawID int32) error {
	_, err := q.db.ExecContext(ctx, trashWithdraw, withdrawID)
	return err
}

const updateWithdraw = `-- name: UpdateWithdraw :exec
UPDATE withdraws
SET
    card_number = $2,
    withdraw_amount = $3,
    withdraw_time = $4,
    updated_at = current_timestamp
WHERE
    withdraw_id = $1
    AND deleted_at IS NULL
`

type UpdateWithdrawParams struct {
	WithdrawID     int32     `json:"withdraw_id"`
	CardNumber     string    `json:"card_number"`
	WithdrawAmount int32     `json:"withdraw_amount"`
	WithdrawTime   time.Time `json:"withdraw_time"`
}

// Update Withdraw
func (q *Queries) UpdateWithdraw(ctx context.Context, arg UpdateWithdrawParams) error {
	_, err := q.db.ExecContext(ctx, updateWithdraw,
		arg.WithdrawID,
		arg.CardNumber,
		arg.WithdrawAmount,
		arg.WithdrawTime,
	)
	return err
}
