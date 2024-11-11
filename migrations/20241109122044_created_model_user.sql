-- +goose Up
CREATE TYPE user_role AS ENUM ('ADMIN', 'USER');

CREATE TABLE IF NOT EXISTS "user".users (
                                            id BIGSERIAL PRIMARY KEY,
                                            email VARCHAR UNIQUE,
                                            password VARCHAR,
                                            name VARCHAR,
                                            role user_role DEFAULT 'USER'
);

-- +goose Down
DROP TABLE IF EXISTS "user".users;
DROP TYPE IF EXISTS user_role;