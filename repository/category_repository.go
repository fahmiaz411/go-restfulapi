package repository

import (
	"context"
	"database/sql"
	"errors"
	"go-restfulapi/helper"
	"go-restfulapi/model/domain"
	"go-restfulapi/model/web/category_web"
)

type CategoryRepository interface {
	Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Update(ctx context.Context, tx *sql.Tx, category domain.Category) *category_web.CategoryUpdateResponse
	Delete(ctx context.Context, tx *sql.Tx, category domain.Category) *category_web.CategoryDeleteResponse
	FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Category
}

type CategoryRepositoryImpl struct {
	
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "INSERT INTO category(name) VALUES (?)"
	result, err := tx.ExecContext(ctx, SQL, category.Name)
	helper.PanicError(err)

	id, err := result.LastInsertId()
	helper.PanicError(err)

	category.Id = int(id)
	return category
}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) *category_web.CategoryUpdateResponse {
	SQL := "UPDATE category SET name = ? WHERE id = ?"
	result, err := tx.ExecContext(ctx, SQL, category.Name, category.Id)
	helper.PanicError(err)

	updatedCount, err := result.RowsAffected()
	helper.PanicError(err)

	return &category_web.CategoryUpdateResponse{
		UpdatedCount: int(updatedCount),
	}
}

func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) *category_web.CategoryDeleteResponse {
	SQL := "DELETE FROM category WHERE id = ?"
	result, err := tx.ExecContext(ctx, SQL, category.Id)
	helper.PanicError(err)

	deletedCount, err := result.RowsAffected()
	return &category_web.CategoryDeleteResponse{
		DeletedCount: int(deletedCount),
	}

}

func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
	SQL := "SELECT id,name FROM category WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, categoryId)
	helper.PanicError(err)
	defer rows.Close()

	category := domain.Category{}
	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicError(err)

		return category, nil
	} else {
		 return category, errors.New("category is not found")
	}
}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	SQL := "SELECT id, name FROM category"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicError(err)
	defer rows.Close()

	var categories []domain.Category
	for rows.Next(){
		category := domain.Category{}

		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicError(err)

		categories = append(categories, category)
	}

	return categories
}