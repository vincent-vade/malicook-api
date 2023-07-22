-- migrate:up
CREATE TABLE IF NOT EXISTS "users" (
    "id" uuid PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
    "first_name" varchar(80) NOT NULL,
    "last_name" varchar(60) NOT NULL,
    "email" varchar(60) NOT NULL,
    "password" varchar NOT NULL,
    "is_active" boolean DEFAULT false,
    "created_at" timestamp with time zone DEFAULT now() NOT NULL,
    "updated_at" timestamp with time zone
);
CREATE INDEX IF NOT EXISTS "name_idx" ON "users" ("first_name");
CREATE UNIQUE INDEX IF NOT EXISTS "email_idx" ON "users" ("email");

-- migrate:down
DROP TABLE users;