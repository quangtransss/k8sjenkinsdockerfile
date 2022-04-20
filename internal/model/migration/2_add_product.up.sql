
CREATE TABLE "orders" (
  "id" BIGSERIAL PRIMARY KEY NOT NULL,
  "customerid" bigint NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "order_detail" (
  "order_id" bigint PRIMARY KEY NOT NULL,
  "product_id" bigint NOT NULL,
  "active" status NOT NULL,
  "total" bigint NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "product" (
  "id" bigint PRIMARY KEY NOT NULL,
  "name_product" varchar(255) NOT NULL,
  "price" bigint NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "product_category" (
  "category_id" bigint PRIMARY KEY NOT NULL,
  "product_id" bigint NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "category" (
  "id" BIGSERIAL PRIMARY KEY NOT NULL,
  "name_category" varchar(255) NOT NULL
);




ALTER TABLE "orders" ADD FOREIGN KEY ("id") REFERENCES "order_detail" ("order_id");

ALTER TABLE "product" ADD FOREIGN KEY ("id") REFERENCES "order_detail" ("product_id");

ALTER TABLE "product" ADD FOREIGN KEY ("id") REFERENCES "product_category" ("product_id");

ALTER TABLE "product_category" ADD FOREIGN KEY ("category_id") REFERENCES "category" ("id");



CREATE INDEX ON "product_category" ("product_id");
CREATE INDEX ON "order_detail" ("order_id");

CREATE INDEX ON "product" ("id");
CREATE UNIQUE INDEX ON "orders" ("id", "customerid");

CREATE UNIQUE INDEX ON "order_detail" ("order_id", "product_id");
CREATE UNIQUE INDEX ON "product_category" ("category_id", "product_id");