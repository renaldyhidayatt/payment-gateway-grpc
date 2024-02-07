-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "saldo" (
    "saldo_id" SERIAL PRIMARY KEY,
    "user_id" INTEGER REFERENCES "users"(user_id) ON DELETE CASCADE ON UPDATE CASCADE NOT NULL,
    "total_balance" INTEGER NOT NULL,
    "withdraw_amount" INTEGER DEFAULT 0,
    "withdraw_time" TIMESTAMP DEFAULT NULL,
    "created_at" timestamp DEFAULT current_timestamp,
    "updated_at" timestamp DEFAULT current_timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXITS "saldo";
-- +goose StatementEnd
