-- migrate:up
CREATE TABLE IF NOT EXISTS "categories" (
    "id" uuid PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
    "name" varchar(40) NOT NULL,
    "created_at" timestamp with time zone DEFAULT now() NOT NULL,
    "updated_at" timestamp with time zone
);
CREATE UNIQUE INDEX IF NOT EXISTS "name_idx" ON "categories" ("name");

-- migrate:down
DROP TABLE categories;