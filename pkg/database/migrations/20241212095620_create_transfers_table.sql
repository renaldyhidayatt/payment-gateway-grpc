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

CREATE INDEX idx_transfers_transfer_from ON transfers (transfer_from);

CREATE INDEX idx_transfers_transfer_to ON transfers (transfer_to);

CREATE INDEX idx_transfers_transfer_time ON transfers (transfer_time);

CREATE INDEX idx_transfers_transfer_amount ON transfers (transfer_amount);

CREATE INDEX idx_transfers_transfer_from_transfer_time ON transfers (transfer_from, transfer_time);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_transfers_transfer_from;

DROP INDEX IF EXISTS idx_transfers_transfer_to;

DROP INDEX IF EXISTS idx_transfers_transfer_time;

DROP INDEX IF EXISTS idx_transfers_transfer_amount;

DROP INDEX IF EXISTS idx_transfers_transfer_from_transfer_time;

DROP TABLE IF EXISTS "transfers";

-- +goose StatementEnd
