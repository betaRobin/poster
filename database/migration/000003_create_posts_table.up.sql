BEGIN;

CREATE TABLE IF NOT EXISTS posts (
   id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
   title TEXT NOT NULL,
   description TEXT NOT NULL,
   user_id UUID,
   CONSTRAINT fk_user
        FOREIGN KEY(user_id)
            REFERENCES users(id)
            ON DELETE CASCADE
);

COMMIT;