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

-- name: EditUser :exec
UPDATE users SET email = $2,
                 first_name = $3,
                 last_name = $4,
                 password = 5,
                 updated_at = now()
WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;

-- name: CreateCategory :one
INSERT INTO category (
    name
) VALUES (
 $1
)
RETURNING *;

-- name: ListCategories :many
SELECT * FROM category;

-- name: GetCategory :one
SELECT * FROM category WHERE id = $1 LIMIT 1;

-- name: CountCategory :one
SELECT count(*) FROM category;

-- name: EditCategory :exec
UPDATE category SET name = $2, updatedAt = now()
WHERE id = $1;

-- name: DeleteCategory :exec
DELETE FROM category WHERE id = $1;