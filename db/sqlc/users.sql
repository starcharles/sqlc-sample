-- name: GetUser :one
SELECT * FROM users
WHERE users.id = ?;

-- name: GetUsers :many
SELECT * FROM users
ORDER BY users.id
LIMIT ?;
-- name: CreateUser :execresult
INSERT INTO users (name, email, password)
VALUES (?, ?, ?);