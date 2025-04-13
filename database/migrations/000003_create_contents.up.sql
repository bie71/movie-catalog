CREATE TABLE IF NOT EXISTS "contents" (
  "id" SERIAL PRIMARY KEY,
  "genre_id" int,
  "director_id" int,
  "title" VARCHAR(200),
  "excerpt" VARCHAR(100),
  "description" TEXT,
  "image" VARCHAR(200),
  "publish_year" int,
  "status" BOOLEAN,
  "tags" VARCHAR(200),
  "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);



