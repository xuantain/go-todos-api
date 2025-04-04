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

func GetTodoListMockData() []models.Todo {
	return []models.Todo{
		{
			ID:          1,
			UserId:      1,
			Description: "Learn AWS",
			TargetDate:  20250615,
			Done:        false,
		},
		{
			ID:          2,
			UserId:      1,
			Description: "Learn Azure",
			TargetDate:  20250715,
			Done:        false,
		},
		{
			ID:          3,
			UserId:      1,
			Description: "Learn DevOp",
			TargetDate:  20250815,
			Done:        false,
		},
		{
			ID:          4,
			UserId:      1,
			Description: "Learn Deno",
			TargetDate:  20250915,
			Done:        false,
		},
	}
}
