ALTER TABLE "contents" DROP CONSTRAINT IF EXISTS fk_contents_genre;
ALTER TABLE "contents" DROP CONSTRAINT IF EXISTS fk_contents_director;

DROP TABLE IF EXISTS directors;