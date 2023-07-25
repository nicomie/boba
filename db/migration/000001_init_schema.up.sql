CREATE TABLE "accounts" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "balance" bigint NOT NULL,
  "currency" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "items" (
  "item_id" bigserial PRIMARY KEY,
  "item_name" varchar NOT NULL,
  "amount" bigint NOT NULL
);

CREATE TABLE "order_details" (
  "order_detail_id" bigserial PRIMARY KEY,
  "order_id" bigserial,
  "item_id" bigserial,
  "quantity" int NOT NULL
);

CREATE TABLE "orders" (
  "id" bigserial PRIMARY KEY
);

CREATE TABLE "entries" (
  "id" bigserial PRIMARY KEY,
  "order_id" bigserial,
  "account_id" bigint,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "transfers" (
  "id" bigserial PRIMARY KEY,
  "from_account_id" bigint,
  "to_account_id" bigint,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "accounts" ("name");

CREATE INDEX ON "entries" ("id");

CREATE INDEX ON "transfers" ("to_account_id");

CREATE INDEX ON "transfers" ("from_account_id", "to_account_id");

COMMENT ON COLUMN "transfers"."amount" IS 'it must be positive';

ALTER TABLE "order_details" ADD FOREIGN KEY ("order_id") REFERENCES "orders" ("id");

ALTER TABLE "order_details" ADD FOREIGN KEY ("item_id") REFERENCES "items" ("item_id");

ALTER TABLE "entries" ADD FOREIGN KEY ("order_id") REFERENCES "orders" ("id");

ALTER TABLE "entries" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("from_account_id") REFERENCES "accounts" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("to_account_id") REFERENCES "accounts" ("id");