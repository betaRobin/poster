BEGIN;

ALTER TABLE posts
RENAME content TO description;

COMMIT;