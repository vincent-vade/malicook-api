package recipes

import (
	"github.com/google/uuid"
	"recipes-api/domain/core"
	"recipes-api/infra/sqlc"
)

type recipeService struct {
	recipeRepository RecipeRepository
}

type IRecipeService interface {
	NewRecipe(ch *core.ContextHandler, rcp sqlc.Recipe) error
	FindRecipe(ch *core.ContextHandler, id uuid.UUID) sqlc.Recipe
	ListRecipes(ch *core.ContextHandler) ([]sqlc.Recipe, error)
}

func NewRecipeService(r RecipeRepository) IRecipeService {
	return &recipeService{
		recipeRepository: r,
	}
}

func (rs recipeService) NewRecipe(ch *core.ContextHandler, rcp sqlc.Recipe) error {
	_, err := rs.recipeRepository.Create(ch.Ctx, rcp)
	return err
}

func (rs recipeService) FindRecipe(ch *core.ContextHandler, id uuid.UUID) sqlc.Recipe {
	rcp, _ := rs.recipeRepository.GetByID(ch.Ctx, id)
	return rcp
}

func (rs recipeService) ListRecipes(ch *core.ContextHandler) ([]sqlc.Recipe, error) {
	rcps, err := rs.recipeRepository.List(ch.Ctx)
	return rcps, err
}
