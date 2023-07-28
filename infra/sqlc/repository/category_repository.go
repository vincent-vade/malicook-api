package repository

import (
	"context"
	"github.com/google/uuid"
	"log"
	"recipes-api/domain/categories"
	sqlc2 "recipes-api/infra/sqlc"
)

type categoryRepository struct {
	sqlc *sqlc2.Queries
}

func NewCategoryRepository(sqlc *sqlc2.Queries) categories.CategoryRepository {
	return &categoryRepository{
		sqlc: sqlc,
	}
}

func (ct *categoryRepository) List(c context.Context) ([]sqlc2.Category, error) {
	lc, err := ct.sqlc.ListCategories(c)
	return lc, err
}

func (ct *categoryRepository) Create(c context.Context, name string) error {
	_, err := ct.sqlc.CreateCategory(c, name)
	return err
}

func (ct *categoryRepository) Edit(c context.Context, category sqlc2.Category) error {
	err := ct.sqlc.EditCategory(c, sqlc2.EditCategoryParams{
		Name: category.Name,
	})
	return err
}

func (ct *categoryRepository) Remove(c context.Context, id uuid.UUID) error {
	err := ct.sqlc.DeleteCategory(c, id)
	return err
}

func (ct *categoryRepository) GetByID(c context.Context, id uuid.UUID) (sqlc2.Category, error) {
	cat, err := ct.sqlc.GetCategory(c, id)
	if err != nil {
		log.Fatal(err)
	}
	return cat, err
}
