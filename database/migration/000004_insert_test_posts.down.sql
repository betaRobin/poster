BEGIN;

DELETE FROM posts
WHERE title = 'Title 1' AND description = 'Desc 1' AND user_id = (SELECT id FROM users WHERE username = 'UserOne');

DELETE FROM posts
WHERE title = 'Title 2' AND description = 'Desc 2' AND user_id = (SELECT id FROM users WHERE username = 'UserTwo');

COMMIT;