-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "user_roles" (
    "user_role_id" SERIAL PRIMARY KEY,
    "user_id" INT NOT NULL REFERENCES "users" ("user_id") ON DELETE CASCADE,
    "role_id" INT NOT NULL REFERENCES "roles" ("role_id") ON DELETE CASCADE,
    "created_at" timestamp DEFAULT current_timestamp,
    "updated_at" timestamp DEFAULT current_timestamp,
    "deleted_at" TIMESTAMP DEFAULT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "user_roles";
-- +goose StatementEnd
