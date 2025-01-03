-- name: GetRoles :many
SELECT 
    role_id, 
    role_name, 
    created_at, 
    updated_at, 
    deleted_at 
FROM 
    roles
WHERE 
    role_name ILIKE '%' || $1 || '%'
ORDER BY 
    created_at ASC
LIMIT $2 OFFSET $3;


-- name: GetRole :one
SELECT 
    role_id, 
    role_name, 
    created_at, 
    updated_at, 
    deleted_at 
FROM 
    roles
WHERE 
    role_id = $1;


-- name: GetRoleByName :one
SELECT 
    role_id, 
    role_name, 
    created_at, 
    updated_at, 
    deleted_at 
FROM 
    roles
WHERE 
    role_name = $1;


-- name: GetUserRoles :many
SELECT 
    r.role_id, 
    r.role_name, 
    r.created_at, 
    r.updated_at, 
    r.deleted_at 
FROM 
    roles r
JOIN 
    user_roles ur ON ur.role_id = r.role_id
WHERE 
    ur.user_id = $1
ORDER BY 
    r.created_at ASC;


-- name: GetActiveRoles :many
SELECT 
    role_id, 
    role_name, 
    created_at, 
    updated_at, 
    deleted_at 
FROM 
    roles
WHERE 
    deleted_at IS NULL
    AND role_name ILIKE '%' || $1 || '%'
ORDER BY 
    created_at ASC
LIMIT $2 OFFSET $3;


-- name: GetTrashedRoles :many
SELECT 
    role_id, 
    role_name, 
    created_at, 
    updated_at, 
    deleted_at 
FROM 
    roles
WHERE 
    deleted_at IS NOT NULL
    AND role_name ILIKE '%' || $1 || '%'
ORDER BY 
    deleted_at DESC
LIMIT $2 OFFSET $3;


-- name: CreateRole :one
INSERT INTO roles (
    role_name, 
    created_at, 
    updated_at
) VALUES (
    $1, 
    current_timestamp, 
    current_timestamp
) RETURNING 
    role_id, 
    role_name, 
    created_at, 
    updated_at, 
    deleted_at;

-- name: UpdateRole :one
UPDATE roles
SET 
    role_name = $2,
    updated_at = current_timestamp
WHERE 
    role_id = $1
RETURNING 
    role_id, 
    role_name, 
    created_at, 
    updated_at, 
    deleted_at;


-- name: TrashRole :exec
UPDATE roles
SET 
    deleted_at = current_timestamp
WHERE 
    role_id = $1;


-- name: RestoreRole :exec
UPDATE roles
SET 
    deleted_at = NULL
WHERE 
    role_id = $1;


-- name: DeletePermanentRole :exec
DELETE FROM roles
WHERE 
    role_id = $1;
