-- +goose Up
-- +goose StatementBegin
CREATE TABLE "saldos" (
    "saldo_id" SERIAL PRIMARY KEY,
    "card_number" VARCHAR(16) NOT NULL REFERENCES "cards" ("card_number"),
    "total_balance" INT NOT NULL,
    "withdraw_amount" INT DEFAULT 0,
    "withdraw_time" TIMESTAMP DEFAULT current_timestamp,
    "created_at" timestamp DEFAULT current_timestamp,
    "updated_at" timestamp DEFAULT current_timestamp,
    "deleted_at" TIMESTAMP DEFAULT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "saldos";
-- +goose StatementEnd