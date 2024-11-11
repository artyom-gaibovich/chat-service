-- +goose Up
CREATE TABLE IF NOT EXISTS "user".users_to_chats
(
    id      SERIAL PRIMARY KEY,
    user_id INTEGER,
    chat_id INTEGER,
    CONSTRAINT fk_users FOREIGN KEY (user_id) REFERENCES "user".users (id),
    CONSTRAINT fk_chats FOREIGN KEY (chat_id) REFERENCES "user".chats (id)
);
-- +goose Down
DROP TABLE "user".users_to_chats ;
