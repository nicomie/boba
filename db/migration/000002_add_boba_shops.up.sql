CREATE TABLE "boba_shop" (
  "shop_id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL
);

CREATE TABLE "personnel" (
  "badge_number" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "shop" bigint NOT NULL,
  "hashed_pin" int4 NOT NULL,
  "email" varchar DEFAULT 'editsbymystic@gmail.com',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "orders" ADD COLUMN "boba_shop" bigint;

ALTER TABLE "orders" ADD FOREIGN KEY ("boba_shop") REFERENCES "boba_shop" ("shop_id");

ALTER TABLE "personnel" ADD FOREIGN KEY ("shop") REFERENCES "boba_shop" ("shop_id");