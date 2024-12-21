-- name: GetUser :one
SELECT * from users
WHERE users.Email = $1;