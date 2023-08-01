-- name: GetUser :one
SELECT id, name, email FROM users
WHERE users.id = ?;

-- name: GetUsers :many
SELECT  id, name, email FROM users
ORDER BY users.id
LIMIT ?;
-- name: CreateUser :execresult
INSERT INTO users (name, email, password)
VALUES (?, ?, ?);