-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
    INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
    VALUES(
        $1,
        $2,
        $3,
        $4,
        $5
    )
    RETURNING *
)
SELECT i.*, f.name As feed_name, u.name AS user_name
FROM inserted_feed_follow i
INNER JOIN feeds f
ON i.feed_id = f.id
INNER JOIN users u
ON u.id = i.user_id;

-- name: GetFeedFollowsForUser :many
SELECT f.name As feed_name, u.name AS user_name
FROM feed_follows ff
INNER JOIN feeds f
ON ff.feed_id = f.id
INNER JOIN users u
ON u.id = ff.user_id
WHERE ff.user_id = $1;

-- name: DeleteFollow :exec
DELETE FROM feed_follows where user_id = $1 AND feed_id = $2;