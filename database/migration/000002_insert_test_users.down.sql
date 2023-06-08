BEGIN;

DELETE FROM users
WHERE username = 'UserOne' AND password = 'PwOne';

DELETE FROM users
WHERE username = 'UserTwo' AND password = 'PwTwo';

COMMIT;