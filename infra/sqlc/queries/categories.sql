-- name: CreateCategory :one
INSERT INTO categories (
    name
) VALUES (
             $1
         )
RETURNING *;

-- name: ListCategories :many
SELECT * FROM categories;

-- name: GetCategory :one
SELECT * FROM categories WHERE id = $1 LIMIT 1;

-- name: CountCategory :one
SELECT count(*) FROM categories;

-- name: EditCategory :exec
UPDATE categories SET name = $2, updated_at = now()
WHERE id = $1;

-- name: DeleteCategory :exec
DELETE FROM categories WHERE id = $1;