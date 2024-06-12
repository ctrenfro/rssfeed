-- +goose Up
CREATE TABLE feeds(
    id UUID primary key,
    created_at timestamp not null,
    updated_at timestamp not null,
    name text not null, 
    url TEXT NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE feeds;