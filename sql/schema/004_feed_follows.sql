-- +goose Up
CREATE TABLE feed_follows(
    id UUID primary key,
    created_at timestamp not null,
    updated_at timestamp not null,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    feed_id UUID NOT NULL REFERENCES feeds(id) ON DELETE CASCADE,
    UNIQUE(user_id, feed_id)
);

-- +goose Down
DROP TABLE feed_follows;
