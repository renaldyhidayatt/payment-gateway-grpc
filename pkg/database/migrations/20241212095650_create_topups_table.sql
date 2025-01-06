-- +goose Up
-- +goose StatementBegin
CREATE TABLE "topups" (
    "topup_id" SERIAL PRIMARY KEY,
    "card_number" VARCHAR(16) NOT NULL REFERENCES "cards" ("card_number"),
    "topup_no" VARCHAR(50) UNIQUE NOT NULL,
    "topup_amount" INT NOT NULL,
    "topup_method" VARCHAR(50) NOT NULL,
    "topup_time" TIMESTAMP NOT NULL,
    "created_at" timestamp DEFAULT current_timestamp,
    "updated_at" timestamp DEFAULT current_timestamp,
    "deleted_at" TIMESTAMP DEFAULT NULL
);

CREATE INDEX idx_topups_card_number ON topups (card_number);

CREATE INDEX idx_topups_topup_no ON topups (topup_no);

CREATE INDEX idx_topups_topup_time ON topups (topup_time);

CREATE INDEX idx_topups_topup_method ON topups (topup_method);

CREATE INDEX idx_topups_card_number_topup_time ON topups (card_number, topup_time);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_topups_card_number;

DROP INDEX IF EXISTS idx_topups_topup_no;

DROP INDEX IF EXISTS idx_topups_topup_time;

DROP INDEX IF EXISTS idx_topups_topup_method;

DROP INDEX IF EXISTS idx_topups_card_number_topup_time;

DROP TABLE IF EXISTS "topups";

-- +goose StatementEnd
