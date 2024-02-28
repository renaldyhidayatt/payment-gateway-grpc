CREATE TABLE IF NOT EXISTS "users"(
    "user_id" serial PRIMARY KEY,
    "firstname" VARCHAR(100) NOT NULL,
    "lastname" varchar(100) NOT NULL,
    "email" varchar(100) UNIQUE NOT NULL,
    "password" varchar(100) NOT NULL,
    "noc_transfer" varchar(100) UNIQUE NOT NULL DEFAULT 0,
    "created_at" timestamp DEFAULT current_timestamp,
    "updated_at" timestamp DEFAULT current_timestamp
);

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


CREATE TABLE IF NOT EXISTS "saldo" (
    "saldo_id" SERIAL PRIMARY KEY,
    "user_id" INTEGER REFERENCES "users"(user_id) ON DELETE CASCADE ON UPDATE CASCADE NOT NULL,
    "total_balance" INTEGER NOT NULL,
    "withdraw_amount" INTEGER DEFAULT 0,
    "withdraw_time" TIMESTAMP DEFAULT NULL,
    "created_at" timestamp DEFAULT current_timestamp,
    "updated_at" timestamp DEFAULT current_timestamp
);



CREATE TABLE IF NOT EXISTS "transfers" (
    "transfer_id" SERIAL PRIMARY KEY,
    "transfer_from" INTEGER REFERENCES "users"(user_id) ON DELETE CASCADE ON UPDATE CASCADE NOT NULL,
    "transfer_to" INTEGER REFERENCES "users"(user_id) ON DELETE CASCADE ON UPDATE CASCADE NOT NULL,
    "transfer_amount" INTEGER NOT NULL DEFAULT 0,
    "transfer_time" TIMESTAMP NOT NULL,
    "created_at" timestamp DEFAULT current_timestamp,
    "updated_at" timestamp DEFAULT current_timestamp
);



CREATE TABLE IF NOT EXISTS "withdraws" (
    "withdraw_id" SERIAL PRIMARY KEY,
    "user_id" INTEGER REFERENCES users(user_id) ON DELETE CASCADE ON UPDATE CASCADE NOT NULL,
    "withdraw_amount" INTEGER NOT NULL,
    "withdraw_time" TIMESTAMP NOT NULL,
    "created_at" timestamp DEFAULT current_timestamp,
    "updated_at" timestamp DEFAULT current_timestamp
);