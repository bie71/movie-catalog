CREATE TABLE IF NOT EXISTS "actor_content" (
  "actor_id" INT NOT NULL,
  "content_id" INT NOT NULL,
  "role" VARCHAR(100),
  PRIMARY KEY ("actor_id", "content_id"),
  CONSTRAINT fk_actor_content_actor
    FOREIGN KEY ("actor_id") REFERENCES "actors"("id") ON DELETE CASCADE,
  CONSTRAINT fk_actor_content_content
    FOREIGN KEY ("content_id") REFERENCES "contents"("id") ON DELETE CASCADE
);

