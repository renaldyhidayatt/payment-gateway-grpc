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

CREATE INDEX idx_merchants_api_key ON merchants (api_key);

CREATE INDEX idx_merchants_user_id ON merchants (user_id);

CREATE INDEX idx_merchants_status ON merchants (status);

CREATE INDEX idx_merchants_name ON merchants (name);

CREATE INDEX idx_merchants_user_id_status ON merchants (user_id, status);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_merchants_api_key;

DROP INDEX IF EXISTS idx_merchants_user_id;

DROP INDEX IF EXISTS idx_merchants_status;

DROP INDEX IF EXISTS idx_merchants_name;

DROP INDEX IF EXISTS idx_merchants_user_id_status;

DROP TABLE IF EXISTS "merchants";

-- +goose StatementEnd
