-- name: CreateUser :one
INSERT INTO users (
    first_name, last_name, email, password
) VALUES (
             $1, $2, $3, $4
         )
RETURNING *;

-- name: ListUsers :many
SELECT * FROM users;

-- name: GetUser :one
SELECT * FROM users WHERE id = $1 LIMIT 1;

-- name: GetUserEmail :one
SELECT * FROM users WHERE email = $1 LIMIT 1;

-- name: EditUser :exec
UPDATE users SET
                 email = $2,
                 first_name = $3,
                 last_name = $4,
                 password = 5,
                 updated_at = now()
WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;
