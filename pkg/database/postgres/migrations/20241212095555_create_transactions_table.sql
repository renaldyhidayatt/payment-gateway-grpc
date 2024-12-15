-- +goose Up
-- +goose StatementBegin
CREATE TABLE "transactions" (
    "transaction_id" SERIAL PRIMARY KEY,
    "card_number" VARCHAR(16) NOT NULL REFERENCES "cards" ("card_number"),
    "amount" INT NOT NULL,
    "payment_method" VARCHAR(50) NOT NULL,
    "merchant_id" INT NOT NULL REFERENCES "merchants" ("merchant_id"),
    "transaction_time" TIMESTAMP NOT NULL,
    "created_at" timestamp DEFAULT current_timestamp,
    "updated_at" timestamp DEFAULT current_timestamp,
    "deleted_at" TIMESTAMP DEFAULT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "transactions";
-- +goose StatementEnd