package repositories

import (
	"context"
	"errors"
	"go-todos-api/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Mock data
/*
var userList = GetUserListMockData()

func (h UserRepository) GetAllUsers() []models.User {
	return userList
}

func (h UserRepository) GetByID(id uint) *models.User {

	for _, user := range userList {
		if user.ID == id {
			return &user
		}
	}

	return nil
}

func (h UserRepository) InsertNewUser(newUser models.User) models.User {

	userList = append(userList, newUser)

	return newUser
}
*/

// Implement repository methods

func (r *UserRepository) Create(ctx context.Context, user *models.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *UserRepository) FindByID(ctx context.Context, id uint) (*models.User, error) {
	var user models.User
	err := r.db.WithContext(ctx).First(&user, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &user, err
}

func (r *UserRepository) Update(ctx context.Context, user *models.User) error {
	return r.db.WithContext(ctx).Save(user).Error
}

func (r *UserRepository) Delete(ctx context.Context, id uint) error {
	// Hard Delete
	// return r.WithContext(ctx).Unscoped().Delete(&models.User{}, id).Error
	// Soft Delete
	return r.db.WithContext(ctx).Delete(&models.User{}, id).Error
}

func (r *UserRepository) List(ctx context.Context, page, limit int) ([]*models.User, error) {
	var users []*models.User
	offset := (page - 1) * limit
	err := r.db.WithContext(ctx).Offset(offset).Limit(limit).Find(&users).Error
	return users, err
}

func (r *UserRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&models.User{}).Count(&count).Error
	return count, err
}

// func (r *UserRepository) WithTransaction(ctx context.Context, fn func(context.Context) error) error {
// 	return r.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
// 		txCtx := context.WithValue(ctx, "tx", tx)
// 		return fn(txCtx)
// 	})
// }

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &user, err
}

func (r *UserRepository) FindByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User
	err := r.db.WithContext(ctx).Where("username = ?", username).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &user, err
}
