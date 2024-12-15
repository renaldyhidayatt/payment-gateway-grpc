-- +goose Up
-- +goose StatementBegin
CREATE TABLE transfers (
    "transfer_id" SERIAL PRIMARY KEY,
    "transfer_from" VARCHAR(16) NOT NULL REFERENCES cards ("card_number"),
    "transfer_to" VARCHAR(16) NOT NULL REFERENCES cards ("card_number"),
    "transfer_amount" INT NOT NULL,
    "transfer_time" TIMESTAMP NOT NULL,
    "created_at" timestamp DEFAULT current_timestamp,
    "updated_at" timestamp DEFAULT current_timestamp,
    "deleted_at" TIMESTAMP DEFAULT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "transfers";
-- +goose StatementEnd