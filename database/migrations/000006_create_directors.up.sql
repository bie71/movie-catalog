CREATE TABLE IF NOT EXISTS "directors" (
  "id" SERIAL PRIMARY KEY,
  "name" VARCHAR(100),
  "birthdate" TIMESTAMP,
  "bio" VARCHAR(200),
  "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE "contents"
ADD CONSTRAINT fk_contents_genre
FOREIGN KEY ("genre_id") REFERENCES "genres"("id");

ALTER TABLE "contents"
ADD CONSTRAINT fk_contents_director
FOREIGN KEY ("director_id") REFERENCES "directors"("id");
