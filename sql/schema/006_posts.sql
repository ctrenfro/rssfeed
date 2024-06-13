-- +goose Up
CREATE TABLE posts(
    id UUID primary key,
    created_at timestamp not null,
    updated_at timestamp not null,
    title TEXT NOT NULL,
    description TEXT,
    published_at TIMESTAMP NOT NULL,
    url TEXT NOT NULL UNIQUE,
    feed_id UUID NOT NULL REFERENCES feeds(id) ON DELETE CASCADE
);


-- +goose Down
 DROP TABLE posts;
