-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "topups" (
    "topup_id" SERIAL PRIMARY KEY,
    "user_id" INTEGER REFERENCES users(user_id) ON DELETE CASCADE ON UPDATE CASCADE NOT NULL,
    "topup_no" TEXT NOT NULL,
    "topup_amount" INTEGER NOT NULL,
    "topup_method" TEXT NOT NULL,
    "topup_time" TIMESTAMP NOT NULL,
    "created_at" timestamp DEFAULT current_timestamp,
    "updated_at" timestamp DEFAULT current_timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "topups";
-- +goose StatementEnd
