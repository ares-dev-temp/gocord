-- +goose Up
CREATE TABLE chatlogs(
    id UUID PRIMARY KEY,
    sent_at TIMESTAMP NOT NULL,
    message TEXT NOT NULL,
    user_id UUID REFERENCES users (id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE chatlogs;
