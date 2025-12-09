CREATE TABLE "user" (
  "id" uuid PRIMARY KEY,
  "username" varchar NOT NULL,
  "balance" float DEFAULT 0,
  "affiliate_id" uuid,
  "created_at" timestamp,
  "updated_at" timestamp
);

CREATE TABLE "product" (
  "id" uuid PRIMARY KEY,
  "name" varchar NOT NULL,
  "quantity" integer NOT NULL,
  "price" float NOT NULL,
  "created_at" timestamp,
  "updated_at" timestamp
);

CREATE TABLE "order" (
  "id" uuid PRIMARY KEY,
  "user_id" uuid NOT NULL,
  "product_detail" jsonb NOT NULL,
  "total_amount" float NOT NULL DEFAULT 0,
  "total_commission" float NOT NULL DEFAULT 0,
  "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "commission" (
  "id" uuid PRIMARY KEY,
  "order_id" uuid NOT NULL,
  "affiliate_id" uuid NULL,
  "amount" float NOT NULL DEFAULT 0,
  "created_at" timestamp,
  "updated_at" timestamp
);

CREATE TABLE "affiliate" (
  "id" uuid PRIMARY KEY,
  "name" varchar(255) NOT NULL,
  "master_affiliate" uuid,
  "balance" float NOT NULL DEFAULT 0,
  "created_at" timestamp,
  "updated_at" timestamp
);

CREATE INDEX ON "user" ("affiliate_id");

CREATE INDEX ON "order" ("user_id");

CREATE INDEX ON "commission" ("order_id");

CREATE INDEX ON "commission" ("affiliate_id");

ALTER TABLE "user" ADD FOREIGN KEY ("affiliate_id") REFERENCES "affiliate" ("id");

ALTER TABLE "order" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "commission" ADD FOREIGN KEY ("order_id") REFERENCES "order" ("id");

ALTER TABLE "commission" ADD FOREIGN KEY ("affiliate_id") REFERENCES "affiliate" ("id");
