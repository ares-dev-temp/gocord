-- +goose Up
CREATE TABLE chatlogs(
    id UUID PRIMARY KEY,
    sent_at TIMESTAMP NOT NULL,
    message TEXT NOT NULL
);

-- +goose Down
DROP TABLE chatlogs;
