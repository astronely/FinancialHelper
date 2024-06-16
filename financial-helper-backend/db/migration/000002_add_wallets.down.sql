ALTER TABLE "wallets" ADD FOREIGN KEY ("owner") REFERENCES "users" ("id");

ALTER TABLE "expenses" ADD FOREIGN KEY ("owner") REFERENCES "users" ("id");

ALTER TABLE "expenses" ADD FOREIGN KEY ("wallet") REFERENCES "wallets" ("id");

DROP TABLE IF EXISTS "expenses";

DROP TABLE IF EXISTS "wallets";

