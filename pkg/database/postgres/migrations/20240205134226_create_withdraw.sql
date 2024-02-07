-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "withdraws" (
    "withdraw_id" SERIAL PRIMARY KEY,
    "user_id" INTEGER REFERENCES users(user_id) ON DELETE CASCADE ON UPDATE CASCADE NOT NULL,
    "withdraw_amount" INTEGER NOT NULL,
    "withdraw_time" TIMESTAMP NOT NULL,
    "created_at" timestamp DEFAULT current_timestamp,
    "updated_at" timestamp DEFAULT current_timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXITS "withdraws";
-- +goose StatementEnd
