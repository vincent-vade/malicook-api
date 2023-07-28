-- migrate:up
CREATE TYPE difficulty AS ENUM ('Easy', 'Medium', 'Hard');
CREATE TYPE pricing AS ENUM ('Cheap', 'Not Too Expensive', 'Expensive');
CREATE TABLE IF NOT EXISTS "recipes" (
    "id" uuid PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
    "title" varchar(120) NOT NULL,
    "description" varchar(120),
    "difficulty" "difficulty" NOT NULL,
    "cooking_time" time NOT NULL,
    "preparation_duration" time NOT NULL ,
    "pricing"  "pricing" NOT NULL,
    "number_person" SMALLINT NOT NULL,
    "user_id" uuid NOT NULL,
    "created_at" timestamp with time zone DEFAULT now() NOT NULL,
    "updated_at" timestamp with time zone
);

ALTER TABLE recipes ADD CONSTRAINT "recipes_user_id_users_id_fk" FOREIGN KEY ("user_id") REFERENCES  "users"("id") ON DELETE no action ON UPDATE no action;

-- migrate:down
DROP TABLE recipes;
