-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "transfers" (
    "transfer_id" SERIAL PRIMARY KEY,
    "transfer_from" INTEGER REFERENCES "users"(user_id) ON DELETE CASCADE ON UPDATE CASCADE NOT NULL,
    "transfer_to" INTEGER REFERENCES "users"(user_id) ON DELETE CASCADE ON UPDATE CASCADE NOT NULL,
    "transfer_amount" INTEGER NOT NULL DEFAULT 0,
    "transfer_time" TIMESTAMP NOT NULL,
    "created_at" timestamp DEFAULT current_timestamp,
    "updated_at" timestamp DEFAULT current_timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXITS "transfers";
-- +goose StatementEnd
