-- +goose Up
CREATE SCHEMA IF NOT EXISTS "user";

-- +goose Down
DROP SCHEMA IF EXISTS "user";
