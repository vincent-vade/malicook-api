-- migrate:up
CREATE TABLE IF NOT EXISTS "category" (
    "id" uuid PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
    "name" varchar(40) NOT NULL,
    "createdAt" timestamp with time zone DEFAULT now() NOT NULL,
    "updatedAt" timestamp with time zone
);
CREATE UNIQUE INDEX IF NOT EXISTS "name_idx" ON "category" ("name");

-- migrate:down
DROP TABLE category;