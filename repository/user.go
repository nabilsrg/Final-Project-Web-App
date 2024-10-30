package repository

import (
	"a21hc3NpZ25tZW50/entity"
	"context"
	"errors"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByID(ctx context.Context, id int) (entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (entity.User, error)
	CreateUser(ctx context.Context, user entity.User) (entity.User, error)
	UpdateUser(ctx context.Context, user entity.User) (entity.User, error)
	DeleteUser(ctx context.Context, id int) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) GetUserByID(ctx context.Context, id int) (entity.User, error) {
	user := entity.User{}
	result := r.db.First(&user, id)
	if result.Error != nil {
		return entity.User{}, result.Error
	}
	return user, nil // TODO: replace this <done>
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {
	user := entity.User{}
	result := r.db.Where("email = ?", email).First(&user)

	// Jika ada error dan bukan record not found, kembalikan error
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return entity.User{}, nil // User tidak ditemukan, kembalikan User kosong dan nil
		}
		return entity.User{}, result.Error // Kembalikan error lainnya
	}
	return user, nil // Kembalikan user yang ditemukan dan nil // TODO: replace this <done>
}

func (r *userRepository) CreateUser(ctx context.Context, user entity.User) (entity.User, error) {
	result := r.db.Create(&user)
	if result.Error != nil {
		return entity.User{}, result.Error
	}
	return user, nil // TODO: replace this
}

func (r *userRepository) UpdateUser(ctx context.Context, user entity.User) (entity.User, error) {
	result := r.db.Model(&user).Updates(user)
	if result.Error != nil {
		return entity.User{}, result.Error
	}
	return user, nil // TODO: replace this
}

func (r *userRepository) DeleteUser(ctx context.Context, id int) error {
	result := r.db.Delete(&entity.User{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil // TODO: replace this
}
