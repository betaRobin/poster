BEGIN;

ALTER TABLE posts
RENAME content TO description;

ALTER TABLE posts
DROP COLUMN type;

COMMIT;