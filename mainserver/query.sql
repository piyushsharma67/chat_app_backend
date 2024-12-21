-- name: GetUser :one
SELECT * from users
WHERE users.Email = $1;

-- name: CreateUser :one
INSERT INTO users (
  name, email,password
) VALUES ($1, $2, $3)
RETURNING *;