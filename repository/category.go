package repository

import (
	"a21hc3NpZ25tZW50/entity"
	"context"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetCategoriesByUserId(ctx context.Context, id int) ([]entity.Category, error)
	StoreCategory(ctx context.Context, category *entity.Category) (categoryId int, err error)
	StoreManyCategory(ctx context.Context, categories []entity.Category) error
	GetCategoryByID(ctx context.Context, id int) (entity.Category, error)
	UpdateCategory(ctx context.Context, category *entity.Category) error
	DeleteCategory(ctx context.Context, id int) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db}
}

func (r *categoryRepository) GetCategoriesByUserId(ctx context.Context, id int) ([]entity.Category, error) {
	categories := make([]entity.Category, 0)
	result := r.db.Where("user_id = ?", id).Find(&categories)
	if result.Error != nil {
		return []entity.Category{}, result.Error
	}
	return categories, nil // TODO: replace this
}

func (r *categoryRepository) StoreCategory(ctx context.Context, category *entity.Category) (categoryId int, err error) {
	result := r.db.Create(&category)
	if result.Error != nil {
		return 0, result.Error
	}
	return category.ID, nil // TODO: replace this
}

func (r *categoryRepository) StoreManyCategory(ctx context.Context, categories []entity.Category) error {
	result := r.db.Create(&categories)
	if result.Error != nil {
		return result.Error
	}
	return nil // TODO: replace this
}

func (r *categoryRepository) GetCategoryByID(ctx context.Context, id int) (entity.Category, error) {
	category := entity.Category{}
	result := r.db.First(&category, id)
	if result.Error != nil {
		return entity.Category{}, result.Error
	}
	return category, nil // TODO: replace this
}

func (r *categoryRepository) UpdateCategory(ctx context.Context, category *entity.Category) error {
	result := r.db.Model(&category).Updates(category)
	if result.Error != nil {
		return result.Error
	}
	return nil // TODO: replace this
}

func (r *categoryRepository) DeleteCategory(ctx context.Context, id int) error {
	result := r.db.Delete(&entity.Category{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil // TODO: replace this
}
