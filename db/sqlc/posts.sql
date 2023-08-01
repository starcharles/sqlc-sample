-- name: GetPost :one
SELECT * FROM posts
JOIN users ON users.id = posts.user_id
WHERE posts.id = ?;

-- name: GetPosts :many
SELECT * FROM posts
JOIN users ON users.id = posts.user_id
ORDER BY posts.id
LIMIT ?;
-- name: CreatePost :execresult
INSERT INTO posts (title, body, user_id)
VALUES (?, ?, ?);