-- modify "users" table
UPDATE "public"."users" SET "role" = 'user' WHERE "role" IS NULL;
ALTER TABLE "public"."users" ALTER COLUMN "role" SET NOT NULL, ALTER COLUMN "role" SET DEFAULT 'user';

