BEGIN;

INSERT INTO posts(title, description, user_id)
SELECT 'Title 1', 'Desc 1', u.id
FROM users u
WHERE u.username = 'UserOne' AND u.password = 'PwOne';

INSERT INTO posts(title, description, user_id)
SELECT 'Title 2', 'Desc 2', u.id
FROM users u
WHERE u.username = 'UserTwo' AND u.password = 'PwTwo';

COMMIT;