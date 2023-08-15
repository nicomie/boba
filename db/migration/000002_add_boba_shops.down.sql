ALTER TABLE "orders" DROP COLUMN "boba_shop";

ALTER TABLE "personnel" DROP CONSTRAINT "personnel_shop_fkey";

ALTER TABLE "orders" DROP CONSTRAINT "orders_boba_shop_fkey";

DROP TABLE IF EXISTS "personnel";

DROP TABLE IF EXISTS "boba_shop";


