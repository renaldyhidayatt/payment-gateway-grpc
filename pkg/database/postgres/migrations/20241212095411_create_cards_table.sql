-- +goose Up
-- +goose StatementBegin
CREATE TABLE "cards" (
    "card_id" SERIAL PRIMARY KEY,
    "user_id" INT NOT NULL REFERENCES "users" ("user_id"),
    "card_number" VARCHAR(16) UNIQUE NOT NULL,
    "card_type" VARCHAR(50) NOT NULL,
    "expire_date" DATE NOT NULL,
    "cvv" VARCHAR(3) NOT NULL,
    "card_provider" VARCHAR(50) NOT NULL,
    "created_at" timestamp DEFAULT current_timestamp,
    "updated_at" timestamp DEFAULT current_timestamp,
    "deleted_at" TIMESTAMP DEFAULT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "cards";
-- +goose StatementEnd