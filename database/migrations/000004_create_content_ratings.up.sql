CREATE TABLE IF NOT EXISTS "content_ratings" (
  "id" SERIAL PRIMARY KEY,
  "content_id" int,
  "user_id" int,
  "rating" int,
  "review" TEXT,
  "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE "content_ratings"
ADD CONSTRAINT fk_content_ratings_content
FOREIGN KEY ("content_id") REFERENCES "contents"("id") ON DELETE CASCADE;

ALTER TABLE "content_ratings"
ADD CONSTRAINT fk_content_ratings_user
FOREIGN KEY ("user_id") REFERENCES "users"("id") ON DELETE CASCADE;
