package repository

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"recipes-api/domain/recipes"
	"recipes-api/infra/sqlc"
	"time"
)

type recipeRepository struct {
	sqlc *sqlc.Queries
}

func NewRecipeRepository(sqlc *sqlc.Queries) recipes.RecipeRepository {
	return &recipeRepository{
		sqlc: sqlc,
	}
}

func (r recipeRepository) Create(c context.Context, recipe sqlc.Recipe) (sqlc.Recipe, error) {
	fmt.Println(recipe.PreparationDuration)
	p := time.Date(2023, 12, 24, 00, 30, 00, 0, time.UTC)
	recipe, err := r.sqlc.CreateRecipe(c, sqlc.CreateRecipeParams{
		Title:               recipe.Title,
		Description:         recipe.Description,
		Difficulty:          recipe.Difficulty,
		CookingTime:         time.Now(),
		PreparationDuration: p,
		Pricing:             recipe.Pricing,
		NumberPerson:        recipe.NumberPerson,
		UserID:              recipe.UserID,
	})

	return recipe, err
}

func (r recipeRepository) List(c context.Context) ([]sqlc.Recipe, error) {
	rcp, err := r.sqlc.ListRecipes(c)
	return rcp, err
}

func (r recipeRepository) GetByID(c context.Context, id uuid.UUID) (sqlc.Recipe, error) {
	rcp, err := r.sqlc.GetRecipe(c, id)
	return rcp, err
}

func (r recipeRepository) Edit(c context.Context, recipe sqlc.Recipe) error {
	err := r.sqlc.EditRecipe(c, sqlc.EditRecipeParams{
		ID:                  recipe.ID,
		Title:               recipe.Title,
		CookingTime:         recipe.CookingTime,
		Pricing:             recipe.Pricing,
		PreparationDuration: recipe.PreparationDuration,
		NumberPerson:        recipe.NumberPerson,
		Description:         recipe.Description,
		UserID:              recipe.UserID,
		Difficulty:          recipe.Difficulty,
	})

	return err
}

func (r recipeRepository) Remove(c context.Context, id uuid.UUID) error {
	err := r.sqlc.DeleteRecipe(c, id)
	return err
}
