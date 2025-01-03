-- +goose Up
-- +goose StatementBegin
CREATE TABLE "withdraws" (
    "withdraw_id" SERIAL PRIMARY KEY,
    "card_number" VARCHAR(16) NOT NULL REFERENCES cards ("card_number"),
    "withdraw_amount" INT NOT NULL,
    "withdraw_time" TIMESTAMP NOT NULL,
    "created_at" timestamp DEFAULT current_timestamp,
    "updated_at" timestamp DEFAULT current_timestamp,
    "deleted_at" TIMESTAMP DEFAULT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "withdraws";
-- +goose StatementEnd