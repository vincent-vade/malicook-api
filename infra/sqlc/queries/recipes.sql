-- name: CreateRecipe :one
INSERT INTO recipes (
                title,
                description,
                difficulty,
                cooking_time,
                preparation_duration,
                pricing,
                number_person,
                user_id
) VALUES (
             $1, $2, $3, $4, $5, $6, $7, $8
         )
    RETURNING *;

-- name: ListRecipes :many
SELECT * FROM recipes;

-- name: GetRecipe :one
SELECT * FROM recipes WHERE id = $1 LIMIT 1;

-- name: EditRecipe :exec
UPDATE recipes SET
                 title = $2,
                 cooking_time = $3,
                 pricing = $4,
                 preparation_duration = $5,
                 number_person = $6,
                 description = $7,
                 user_id = $8,
                 difficulty = $9,
                 updated_at = now()
WHERE id = $1;

-- name: DeleteRecipe :exec
DELETE FROM recipes WHERE id = $1;
