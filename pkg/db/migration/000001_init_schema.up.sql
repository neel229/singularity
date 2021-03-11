CREATE TABLE "currency" (
  "id" bigserial PRIMARY KEY,
  "code" varchar(8) NOT NULL,
  "name" varchar(64) NOT NULL,
  "is_base" bool NOT NULL
);
CREATE TABLE "currency_rate" (
  "id" bigserial PRIMARY KEY,
  "currency_id" bigserial NOT NULL,
  "base_currency_id" bigserial NOT NULL,
  "rate" decimal(16, 6) NOT NULL,
  "ts" timestamptz NOT NULL DEFAULT (now())
);
CREATE TABLE "stock" (
  "id" bigserial PRIMARY KEY,
  "ticker" varchar(8) NOT NULL,
  "details" text NOT NULL,
  "mint_price" int NOT NULL,
  "current_price" int NOT NULL
);
CREATE TABLE "report" (
  "id" bigserial PRIMARY KEY,
  "trading_date" date NOT NULL,
  "stock_id" bigserial NOT NULL,
  "currency_id" bigserial NOT NULL,
  "first_price" decimal(16, 6) NOT NULL,
  "last_price" decimal(16, 6) NOT NULL,
  "min_price" decimal(16, 6) NOT NULL,
  "max_price" decimal(16, 6) NOT NULL,
  "avg_price" decimal(16, 6) NOT NULL,
  "total_amount" decimal(16, 6) NOT NULL,
  "volume" decimal(16, 6) NOT NULL
);
CREATE TABLE "fan" (
  "id" bigserial PRIMARY KEY,
  "first_name" varchar(64) NOT NULL,
  "last_name" varchar(64) NOT NULL,
  "user_name" varchar(64) NOT NULL,
  "password" varchar(64) NOT NULL,
  "email" varchar(128) NOT NULL,
  "time_registered" timestamptz NOT NULL DEFAULT (now()),
  "time_confirmed" timestamptz NOT NULL DEFAULT (now()),
  "preferred_currency_id" bigserial NOT NULL
);
CREATE TABLE "creator" (
  "id" bigserial PRIMARY KEY,
  "first_name" varchar(64) NOT NULL,
  "last_name" varchar(64) NOT NULL,
  "user_name" varchar(64) NOT NULL,
  "password" varchar(64) NOT NULL,
  "email" varchar(128) NOT NULL,
  "time_registered" timestamptz NOT NULL DEFAULT (now()),
  "time_confirmed" timestamptz NOT NULL DEFAULT (now()),
  "preferred_currency_id" bigserial NOT NULL,
  "virgin_tokens_left" int NOT NULL
);
CREATE TABLE "creator_stock" (
  "id" bigserial PRIMARY KEY,
  "creator_id" bigserial NOT NULL,
  "stock_id" bigserial NOT NULL
);
CREATE TABLE "fan_portfolio" (
  "id" bigserial PRIMARY KEY,
  "fan_id" bigserial NOT NULL,
  "stock_id" bigserial NOT NULL,
  "quantity" decimal(16, 6) NOT NULL
);
CREATE TABLE "creator_portfolio" (
  "id" bigserial PRIMARY KEY,
  "creator_id" bigserial NOT NULL,
  "stock_id" bigserial NOT NULL,
  "quantity" decimal(16, 6) NOT NULL
);
CREATE TABLE "virgin_offer" (
  "id" bigserial PRIMARY KEY,
  "creator_id" bigserial NOT NULL,
  "stock_id" bigserial NOT NULL,
  "quantity" decimal(16, 6) NOT NULL,
  "price" decimal(16, 6) NOT NULL,
  "ts" timestamptz NOT NULL DEFAULT (now())
);
CREATE TABLE "offer" (
  "id" bigserial PRIMARY KEY,
  "trader_id" bigserial NOT NULL,
  "stock_id" bigserial NOT NULL,
  "quantity" decimal(16, 6) NOT NULL,
  "buy" bool NOT NULL,
  "sell" bool NOT NULL,
  "price" decimal(16, 6) NOT NULL,
  "ts" timestamptz NOT NULL DEFAULT (now())
);
CREATE TABLE "virgin_trade" (
  "id" bigserial PRIMARY KEY,
  "stock_id" bigserial NOT NULL,
  "creator_id" bigserial NOT NULL,
  "fan_id" bigserial NOT NULL,
  "quantity" decimal(16, 6) NOT NULL,
  "unit_price" decimal(16, 6) NOT NULL,
  "details" text NOT NULL,
  "virgin_offer_id" bigserial NOT NULL
);
CREATE TABLE "trade" (
  "id" bigserial PRIMARY KEY,
  "stock_id" bigserial NOT NULL,
  "buyer_id" bigserial NOT NULL,
  "seller_id" bigserial NOT NULL,
  "quantity" decimal(16, 6) NOT NULL,
  "unit_price" decimal(16, 6) NOT NULL,
  "details" text NOT NULL,
  "offer_id" bigserial NOT NULL
);
ALTER TABLE "currency_rate"
ADD FOREIGN KEY ("currency_id") REFERENCES "currency" ("id");
ALTER TABLE "currency_rate"
ADD FOREIGN KEY ("base_currency_id") REFERENCES "currency" ("id");
ALTER TABLE "report"
ADD FOREIGN KEY ("currency_id") REFERENCES "currency" ("id");
ALTER TABLE "fan"
ADD FOREIGN KEY ("preferred_currency_id") REFERENCES "currency" ("id");
ALTER TABLE "creator"
ADD FOREIGN KEY ("preferred_currency_id") REFERENCES "currency" ("id");
ALTER TABLE "report"
ADD FOREIGN KEY ("stock_id") REFERENCES "stock" ("id");
ALTER TABLE "trade"
ADD FOREIGN KEY ("stock_id") REFERENCES "stock" ("id");
ALTER TABLE "offer"
ADD FOREIGN KEY ("trader_id") REFERENCES "creator" ("id");
ALTER TABLE "creator_portfolio"
ADD FOREIGN KEY ("creator_id") REFERENCES "creator" ("id");
ALTER TABLE "virgin_trade"
ADD FOREIGN KEY ("creator_id") REFERENCES "creator" ("id");
ALTER TABLE "trade"
ADD FOREIGN KEY ("buyer_id") REFERENCES "creator" ("id");
ALTER TABLE "trade"
ADD FOREIGN KEY ("seller_id") REFERENCES "creator" ("id");
ALTER TABLE "creator_stock"
ADD FOREIGN KEY ("creator_id") REFERENCES "creator" ("id");
ALTER TABLE "creator_stock"
ADD FOREIGN KEY ("stock_id") REFERENCES "stock" ("id");
ALTER TABLE "offer"
ADD FOREIGN KEY ("trader_id") REFERENCES "fan" ("id");
ALTER TABLE "fan_portfolio"
ADD FOREIGN KEY ("fan_id") REFERENCES "fan" ("id");
ALTER TABLE "virgin_trade"
ADD FOREIGN KEY ("fan_id") REFERENCES "fan" ("id");
ALTER TABLE "trade"
ADD FOREIGN KEY ("buyer_id") REFERENCES "fan" ("id");
ALTER TABLE "trade"
ADD FOREIGN KEY ("seller_id") REFERENCES "fan" ("id");
ALTER TABLE "fan_portfolio"
ADD FOREIGN KEY ("stock_id") REFERENCES "stock" ("id");
ALTER TABLE "creator_portfolio"
ADD FOREIGN KEY ("stock_id") REFERENCES "stock" ("id");
ALTER TABLE "offer"
ADD FOREIGN KEY ("stock_id") REFERENCES "stock" ("id");
ALTER TABLE "trade"
ADD FOREIGN KEY ("offer_id") REFERENCES "offer" ("id");
ALTER TABLE "virgin_trade"
ADD FOREIGN KEY ("virgin_offer_id") REFERENCES "virgin_offer" ("id");
COMMENT ON COLUMN "fan"."password" IS 'Should be 64 based encoded value';
COMMENT ON COLUMN "creator"."password" IS 'Should be 64 based encoded value';
COMMENT ON COLUMN "creator"."virgin_tokens_left" IS 'Number of stocks the creator owns of his/hers';