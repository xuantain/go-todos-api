package models

import (
	"go-todos-api/models"

	"gorm.io/gorm"
)

type UserRepo struct {
	*gorm.DB
}

var userList = GetUserListMockData()

func (h UserRepo) GetAllUsers() []models.User {
	return userList
}

func (h UserRepo) GetByID(id uint) *models.User {

	for _, user := range userList {
		if user.ID == id {
			return &user
		}
	}

	return nil
}

func (h UserRepo) InsertNewUser(newUser models.User) models.User {

	userList = append(userList, newUser)

	return newUser
}
