package service

import (
	"context"
	"database/sql"
	"go-restfulapi/helper"
	"go-restfulapi/model/domain"
	"go-restfulapi/model/web/category_web"
	"go-restfulapi/repository"

	"github.com/go-playground/validator/v10"
)

type CategoryService interface {
	Create(ctx context.Context, request category_web.CategoryCreateRequest) category_web.CategoryResponse
	Update(ctx context.Context, request category_web.CategoryUpdateRequest) *category_web.CategoryUpdateResponse
	Delete(ctx context.Context, categoryId int)  *category_web.CategoryDeleteResponse
	FindById(ctx context.Context, categoryId int) category_web.CategoryResponse
	FindAll(ctx context.Context) []category_web.CategoryResponse
}

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB *sql.DB
	Validate *validator.Validate
}

func NewCategoryService(DB *sql.DB, validate *validator.Validate) CategoryService {
	repo := repository.NewCategoryRepository()
	return &CategoryServiceImpl{
		CategoryRepository: repo,
		DB: DB,
		Validate: validate,
	}
}

func (service *CategoryServiceImpl) Create(ctx context.Context, request category_web.CategoryCreateRequest) category_web.CategoryResponse {

	err := service.Validate.Struct(request)
	helper.PanicError(err)

	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	category := domain.Category{
		Name: request.Name,
	}

	category = service.CategoryRepository.Save(ctx, tx, category)

	return category_web.CategoryResponse{
		Id: category.Id,
		Name: category.Name,
	}
}

func (service *CategoryServiceImpl) Update(ctx context.Context, request category_web.CategoryUpdateRequest) *category_web.CategoryUpdateResponse {

	err := service.Validate.Struct(request)
	helper.PanicError(err)

	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	category := domain.Category{
		Id: request.Id,
		Name: request.Name,
	}

	response := service.CategoryRepository.Update(ctx, tx, category)

	return response
}

func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryId int) *category_web.CategoryDeleteResponse {
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	category := domain.Category{
		Id: categoryId,
	}

	response := service.CategoryRepository.Delete(ctx, tx, category)
	return response
}

func (service *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) category_web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	helper.PanicError(err)

	return category_web.CategoryResponse{
		Id: category.Id,
		Name: category.Name,
	}
}

func (service *CategoryServiceImpl) FindAll(ctx context.Context) []category_web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	categories := service.CategoryRepository.FindAll(ctx,tx)
	var responses []category_web.CategoryResponse

	for _, category := range categories {
		responses = append(responses, category_web.CategoryResponse(category))
	}

	return responses

}

