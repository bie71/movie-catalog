ALTER TABLE "content_ratings" DROP CONSTRAINT IF EXISTS fk_content_ratings_content;
ALTER TABLE "content_ratings" DROP CONSTRAINT IF EXISTS fk_content_ratings_user;

DROP TABLE IF EXISTS content_ratings;