-- name: CreatePost :one
INSERT INTO posts (id, created_at, updated_at, title, url, description, published_at, feed_id)
VALUES(
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8
)
RETURNING *;

-- name: GetPostsForUser :many
SELECT p.*, feeds.name as feed_name
FROM posts p
INNER JOIN feed_follows f
ON f.feed_id = p.feed_id
INNER JOIN feeds
ON feeds.id = p.feed_id
WHERE f.user_id = $1
ORDER BY p.published_at DESC
LIMIT $2;