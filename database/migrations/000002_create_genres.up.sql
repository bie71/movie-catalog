CREATE TABLE IF NOT EXISTS "genres" (
  "id" SERIAL PRIMARY KEY,
  "created_by_id" int,
  "name" VARCHAR(100),
  "slug" VARCHAR(100),
  "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE "genres"
ADD CONSTRAINT fk_genres_created_by
FOREIGN KEY ("created_by_id") REFERENCES "users"("id")
ON DELETE SET NULL;

