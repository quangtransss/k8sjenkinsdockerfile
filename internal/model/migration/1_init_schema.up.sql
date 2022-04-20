CREATE TYPE "status" AS ENUM (
  'active',
  'disable'
);

CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "full_name" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "username" varchar NOT NULL,
  "hashed_password" varchar NOT NULL,
  "email" varchar NOT NULL,
  "mobile" bigint NOT NULL,
  "roleid" bigint NOT NULL,
  "password_change_at" timestamptz NOT NULL DEFAULT '0001-01-01'
);

CREATE TABLE "role" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "title" varchar NOT NULL,
  "slug" varchar NOT NULL,
  "active" status NOT NULL,
  "description" text,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01'
);

CREATE TABLE "role_permisstion" (
  "roleid" bigserial PRIMARY KEY NOT NULL,
  "permissionid" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "permission" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "title" text,
  "content" text,
  "active" status NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "orders" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "customerid" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "order_detail" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "order_id" bigint NOT NULL,
  "product_id" bigint NOT NULL,
  "active" status NOT NULL,
  "total" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "product" (
  "id" bigint PRIMARY KEY NOT NULL,
  "name_product" varchar(255) NOT NULL,
  "price" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "product_category" (
  "category_id" bigint PRIMARY KEY NOT NULL,
  "product_id" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "category" (
  "id" BIGSERIAL PRIMARY KEY NOT NULL,
  "name_category" varchar(255) NOT NULL
);

ALTER TABLE "order_detail" ADD FOREIGN KEY ("order_id") REFERENCES "orders" ("id");

ALTER TABLE "order_detail" ADD FOREIGN KEY ("product_id") REFERENCES "product" ("id");

ALTER TABLE "product_category" ADD FOREIGN KEY ("product_id") REFERENCES "product" ("id");

ALTER TABLE "product_category" ADD FOREIGN KEY ("category_id") REFERENCES "category" ("id");

ALTER TABLE "role_permisstion" ADD FOREIGN KEY ("roleid") REFERENCES "role" ("id");

ALTER TABLE "users" ADD FOREIGN KEY ("roleid") REFERENCES "role" ("id");

ALTER TABLE "role_permisstion" ADD FOREIGN KEY ("permissionid") REFERENCES "permission" ("id");

ALTER TABLE "orders" ADD FOREIGN KEY ("customerid") REFERENCES "users" ("id");

CREATE INDEX ON "users" ("id");

CREATE UNIQUE INDEX ON "users" ("username", "email");

CREATE INDEX ON "role" ("id");

CREATE UNIQUE INDEX ON "role" ("slug");

CREATE INDEX ON "permission" ("id");

CREATE INDEX ON "orders" ("customerid");

CREATE INDEX ON "order_detail" ("order_id");

CREATE UNIQUE INDEX ON "product" ("id", "name_product");

CREATE INDEX ON "product_category" ("product_id");

CREATE UNIQUE INDEX ON "product_category" ("category_id", "product_id");