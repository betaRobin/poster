BEGIN;

ALTER TABLE posts
RENAME description TO content;

COMMIT;