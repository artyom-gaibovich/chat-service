-- +goose Up
CREATE TABLE IF NOT EXISTS "user".chats
(
    id BIGSERIAL PRIMARY KEY
);
-- +goose Down
DROP TABLE IF EXISTS "user".chats;
