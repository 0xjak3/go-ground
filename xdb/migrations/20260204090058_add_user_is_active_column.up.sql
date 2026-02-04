-- create "users" table
CREATE TABLE "public"."users" (
  "id" bigserial NOT NULL,
  "name" text NULL,
  "age" bigint NULL,
  PRIMARY KEY ("id")
);
