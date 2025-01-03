-- +goose Up
-- +goose StatementBegin
CREATE TABLE "merchants" (
    "merchant_id" SERIAL PRIMARY KEY,
    "name" VARCHAR(255) NOT NULL,
    "api_key" VARCHAR(255) UNIQUE NOT NULL,
    "user_id" INT NOT NULL REFERENCES "users" (user_id),
    "status" VARCHAR(50) NOT NULL,
    "created_at" timestamp DEFAULT current_timestamp,
    "updated_at" timestamp DEFAULT current_timestamp,
    "deleted_at" TIMESTAMP DEFAULT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "merchants";
-- +goose StatementEnd