package models

import "go-todos-api/models"

func GetUserListMockData() []models.User {
	return []models.User{
		{
			ID:        1,
			Name:      "Lam",
			BirthDay:  "2000-01-01",
			Gender:    "Male",
			PhotoURL:  "",
			Time:      20250303,
			Active:    true,
			UpdatedAt: 20250401,
		},
		{
			ID:        2,
			Name:      "Lam",
			BirthDay:  "2001-01-15",
			Gender:    "Male",
			PhotoURL:  "",
			Time:      20250303,
			Active:    true,
			UpdatedAt: 20250401,
		},
		{
			ID:        3,
			Name:      "Lam",
			BirthDay:  "2002-03-02",
			Gender:    "Male",
			PhotoURL:  "",
			Time:      20250303,
			Active:    true,
			UpdatedAt: 20250401,
		},
	}
}
