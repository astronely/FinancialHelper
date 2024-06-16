CREATE TYPE "expense_category" AS ENUM (
  'Супермаркеты',
  'Развлечение',
  'Спорт',
  'Красота',
  'Медицина',
  'Фастфуд',
  'Рестораны',
  'Другое'
);

-- CREATE TABLE "users" (
--                          "id" bigserial PRIMARY KEY,
--                          "username" varchar NOT NULL,
--                          "hashed_password" varchar NOT NULL,
--                          "full_name" varchar NOT NULL,
--                          "email" varchar UNIQUE NOT NULL,
--                          "password_change_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
--                          "created_at" timestamptz NOT NULL DEFAULT (now())
-- );

CREATE TABLE "wallets" (
                           "id" bigserial PRIMARY KEY,
                           "owner" bigserial NOT NULL,
                           "name" varchar NOT NULL,
                           "balance" float DEFAULT 0 CHECK ( balance >= 0 ),
                           "currency" varchar NOT NULL,
                           "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "expenses" (
                            "id" bigserial PRIMARY KEY,
                            "owner" bigserial NOT NULL,
                            "wallet" bigserial NOT NULL,
                            "wallet_name" varchar NOT NULL,
                            "currency" varchar NOT NULL,
                            "value" float NOT NULL CHECK ( value >= 0 ),
                            "name" varchar NOT NULL,
                            "category" expense_category NOT NULL,
                            "date" date NOT NULL,
                            "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE INDEX ON "wallets" ("owner");

CREATE UNIQUE INDEX ON "wallets" ("owner", "name");

CREATE INDEX ON "expenses" ("owner");

CREATE INDEX ON "expenses" ("wallet");

CREATE INDEX ON "expenses" ("currency");

CREATE INDEX ON "expenses" ("name");

CREATE INDEX ON "expenses" ("category");

CREATE INDEX ON "expenses" ("date");

CREATE INDEX ON "expenses" ("owner", "wallet");

CREATE INDEX ON "expenses" ("owner", "name");

CREATE INDEX ON "expenses" ("owner", "category");

CREATE INDEX ON "expenses" ("owner", "date");
--
ALTER TABLE "wallets" ADD FOREIGN KEY ("owner") REFERENCES "users" ("id") ON DELETE CASCADE;

ALTER TABLE "expenses" ADD FOREIGN KEY ("owner") REFERENCES "users" ("id");

ALTER TABLE "expenses" ADD FOREIGN KEY ("wallet") REFERENCES "wallets" ("id");

