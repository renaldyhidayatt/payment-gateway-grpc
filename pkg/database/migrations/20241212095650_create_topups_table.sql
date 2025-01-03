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
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "topups";
-- +goose StatementEnd