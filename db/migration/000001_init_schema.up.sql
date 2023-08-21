CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL,
  "hashed_password" varchar NOT NULL,
  "name" varchar NOT NULL,
  "balance" bigint DEFAULT 0,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "products" (
  "product_id" bigserial PRIMARY KEY,
  "name" varchar UNIQUE NOT NULL,
  "price" int4 NOT NULL
);

CREATE TABLE "orders" (
  "order_id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "status" varchar DEFAULT 'pending',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "order_items" (
  "order_item_id" bigserial PRIMARY KEY,
  "order_id" bigint NOT NULL,
  "product_id" bigint NOT NULL,
  "quantity" int4 NOT NULL
);

CREATE INDEX ON "users" ("name");

ALTER TABLE "orders" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "order_items" ADD FOREIGN KEY ("order_id") REFERENCES "orders" ("order_id");

ALTER TABLE "order_items" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("product_id");