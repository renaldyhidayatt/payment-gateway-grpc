-- Create User
-- name: CreateUser :one
INSERT INTO
    users (
        firstname,
        lastname,
        email,
        password,
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

-- Get User by ID
-- name: GetUserByID :one
SELECT * FROM users WHERE user_id = $1 AND deleted_at IS NULL;

-- Get User by Email
-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1 AND deleted_at IS NULL;

-- Get All Active Users
-- name: GetActiveUsers :many
SELECT *
FROM users
WHERE
    deleted_at IS NULL
ORDER BY created_at DESC;

-- Get Trashed Users
-- name: GetTrashedUsers :many
SELECT *
FROM users
WHERE
    deleted_at IS NOT NULL
ORDER BY created_at DESC;

-- Search Users with Pagination
-- name: SearchUsers :many
SELECT *
FROM users
WHERE deleted_at IS NULL
  AND ($1::TEXT IS NULL OR firstname ILIKE '%' || $1 || '%' OR lastname ILIKE '%' || $1 || '%' OR email ILIKE '%' || $1 || '%')
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- Count All Active Users
-- name: CountActiveUsers :one
SELECT COUNT(*) FROM users WHERE deleted_at IS NULL;

-- Trash User
-- name: TrashUser :exec
UPDATE users
SET
    deleted_at = current_timestamp
WHERE
    user_id = $1
    AND deleted_at IS NULL;

-- Restore Trashed User
-- name: RestoreUser :exec
UPDATE users
SET
    deleted_at = NULL
WHERE
    user_id = $1
    AND deleted_at IS NOT NULL;

-- Update User
-- name: UpdateUser :exec
UPDATE users
SET
    firstname = $2,
    lastname = $3,
    email = $4,
    password = $5,
    updated_at = current_timestamp
WHERE
    user_id = $1
    AND deleted_at IS NULL;

-- Delete User Permanently
-- name: DeleteUserPermanently :exec
DELETE FROM users WHERE user_id = $1;

-- Search Users by Email
-- name: SearchUsersByEmail :many
SELECT *
FROM users
WHERE
    deleted_at IS NULL
    AND email ILIKE '%' || $1 || '%'
ORDER BY created_at DESC;

-- Get Trashed By User ID
-- name: GetTrashedUserByID :one
SELECT *
FROM users
WHERE
    user_id = $1
    AND deleted_at IS NOT NULL;