-- +goose Up
CREATE TABLE users(
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    username TEXT NOT NULL
);

ALTER TABLE chatlogs
ADD user_id UUID REFERENCES users (id) ON DELETE CASCADE;

-- +goose Down
DROP TABLE users;
