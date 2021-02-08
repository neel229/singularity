CREATE TABLE "currency" (
  "id" bigserial PRIMARY KEY,
  "code" varchar(8),
  "name" varchar(64),
  "is_base" bool
);

CREATE TABLE "currency_rate" (
  "id" bigserial PRIMARY KEY,
  "currency_id" int,
  "base_currency_id" int,
  "rate" decimal(16,6),
  "ts" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "stock" (
  "id" bigserial PRIMARY KEY,
  "ticker" varchar(8),
  "creator_id" bigserial,
  "details" text
);

CREATE TABLE "price" (
  "id" bigserial PRIMARY KEY,
  "stock_id" bigserial,
  "currency_id" int,
  "buy" decimal(16,6),
  "sell" decimal(16,6),
  "ts" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "report" (
  "id" bigserial PRIMARY KEY,
  "trading_date" date,
  "stock_id" bigserial,
  "currency_id" int,
  "first_price" decimal(16,6),
  "last_price" decimal(16,6),
  "min_price" decimal(16,6),
  "max_price" decimal(16,6),
  "avg_price" decimal(16,6),
  "total_amount" decimal(16,6),
  "volume" decimal(16,6)
);

CREATE TABLE "fan" (
  "id" bigserial PRIMARY KEY,
  "first_name" varchar(64),
  "last_name" varchar(64),
  "user_name" varchar(64),
  "password" varchar(64),
  "email" varchar(128),
  "time_registered" timestamptz NOT NULL DEFAULT (now()),
  "time_confirmed" timestamptz NOT NULL DEFAULT (now()),
  "preferred_currency_id" int
);

CREATE TABLE "creator" (
  "id" bigserial PRIMARY KEY,
  "first_name" varchar(64),
  "last_name" varchar(64),
  "user_name" varchar(64),
  "password" varchar(64),
  "email" varchar(128),
  "time_registered" timestamptz NOT NULL DEFAULT (now()),
  "time_confirmed" timestamptz NOT NULL DEFAULT (now()),
  "preferred_currency_id" int,
  "creator_stock" bigserial,
  "virgin_tokens_left" int
);

CREATE TABLE "portfolio" (
  "id" bigserial PRIMARY KEY,
  "fan_id" bigserial,
  "stock_id" bigserial,
  "quantity" decimal(16,6)
);

CREATE TABLE "virgin_offer" (
  "id" bigserial PRIMARY KEY,
  "creator_id" bigserial,
  "stock_id" bigserial,
  "quantity" decimal(16,6),
  "price" decimal(16,6),
  "ts" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "offer" (
  "id" bigserial PRIMARY KEY,
  "trader_id" bigserial,
  "stock_id" bigserial,
  "quantity" decimal(16,6),
  "buy" bool,
  "sell" bool,
  "price" decimal(16,6),
  "ts" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "virgin_trade" (
  "id" bigserial PRIMARY KEY,
  "stock_id" bigserial,
  "creator_id" bigserial,
  "buyer_id" bigserial,
  "quantity" decimal(16,6),
  "unit_price" decimal(16,6),
  "details" text,
  "virgin_offer_id" bigserial
);

CREATE TABLE "trade" (
  "id" bigserial PRIMARY KEY,
  "stock_id" bigserial,
  "buyer_id" bigserial,
  "seller_id" bigserial,
  "quantity" decimal(16,6),
  "unit_price" decimal(16,6),
  "details" text,
  "offer_id" bigserial
);

ALTER TABLE "currency_rate" ADD FOREIGN KEY ("currency_id") REFERENCES "currency" ("id");

ALTER TABLE "currency_rate" ADD FOREIGN KEY ("base_currency_id") REFERENCES "currency" ("id");

ALTER TABLE "price" ADD FOREIGN KEY ("currency_id") REFERENCES "currency" ("id");

ALTER TABLE "report" ADD FOREIGN KEY ("currency_id") REFERENCES "currency" ("id");

ALTER TABLE "fan" ADD FOREIGN KEY ("preferred_currency_id") REFERENCES "currency" ("id");

ALTER TABLE "creator" ADD FOREIGN KEY ("preferred_currency_id") REFERENCES "currency" ("id");

ALTER TABLE "creator" ADD FOREIGN KEY ("id") REFERENCES "stock" ("id");

ALTER TABLE "price" ADD FOREIGN KEY ("stock_id") REFERENCES "stock" ("id");

ALTER TABLE "report" ADD FOREIGN KEY ("stock_id") REFERENCES "stock" ("id");

ALTER TABLE "trade" ADD FOREIGN KEY ("stock_id") REFERENCES "stock" ("id");

ALTER TABLE "offer" ADD FOREIGN KEY ("trader_id") REFERENCES "creator" ("id");

ALTER TABLE "portfolio" ADD FOREIGN KEY ("fan_id") REFERENCES "creator" ("id");

ALTER TABLE "virgin_trade" ADD FOREIGN KEY ("creator_id") REFERENCES "creator" ("id");

ALTER TABLE "virgin_trade" ADD FOREIGN KEY ("buyer_id") REFERENCES "creator" ("id");

ALTER TABLE "trade" ADD FOREIGN KEY ("buyer_id") REFERENCES "creator" ("id");

ALTER TABLE "trade" ADD FOREIGN KEY ("seller_id") REFERENCES "creator" ("id");

ALTER TABLE "offer" ADD FOREIGN KEY ("trader_id") REFERENCES "fan" ("id");

ALTER TABLE "portfolio" ADD FOREIGN KEY ("fan_id") REFERENCES "fan" ("id");

ALTER TABLE "virgin_trade" ADD FOREIGN KEY ("buyer_id") REFERENCES "fan" ("id");

ALTER TABLE "trade" ADD FOREIGN KEY ("buyer_id") REFERENCES "fan" ("id");

ALTER TABLE "trade" ADD FOREIGN KEY ("seller_id") REFERENCES "fan" ("id");

ALTER TABLE "portfolio" ADD FOREIGN KEY ("stock_id") REFERENCES "stock" ("id");

ALTER TABLE "offer" ADD FOREIGN KEY ("stock_id") REFERENCES "stock" ("id");

ALTER TABLE "trade" ADD FOREIGN KEY ("offer_id") REFERENCES "offer" ("id");

ALTER TABLE "virgin_trade" ADD FOREIGN KEY ("virgin_offer_id") REFERENCES "offer" ("id");

COMMENT ON COLUMN "fan"."password" IS 'Should be 64 based encoded value';

COMMENT ON COLUMN "creator"."password" IS 'Should be 64 based encoded value';

COMMENT ON COLUMN "creator"."virgin_tokens_left" IS 'Number of stocks the creator owns of his/hers';
