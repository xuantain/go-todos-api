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
			Time:      "2025-03-03",
			Active:    true,
			UpdatedAt: "2025-04-01",
		},
		{
			ID:        2,
			Name:      "Lam",
			BirthDay:  "2001-01-15",
			Gender:    "Male",
			PhotoURL:  "",
			Time:      "2025-03-03",
			Active:    true,
			UpdatedAt: "2025-04-01",
		},
		{
			ID:        3,
			Name:      "Lam",
			BirthDay:  "2002-03-02",
			Gender:    "Male",
			PhotoURL:  "",
			Time:      "2025-03-03",
			Active:    true,
			UpdatedAt: "2025-04-01",
		},
	}
}

func GetTodoListMockData() []models.Todo {
	return []models.Todo{
		{
			ID:          1,
			UserName:    "todo",
			Description: "Learn AWS",
			TargetDate:  "2025-06-15",
			Done:        false,
		},
		{
			ID:          2,
			UserName:    "todo",
			Description: "Learn Azure",
			TargetDate:  "2025-07-15",
			Done:        false,
		},
		{
			ID:          3,
			UserName:    "todo",
			Description: "Learn DevOp",
			TargetDate:  "2025-08-15",
			Done:        false,
		},
		{
			ID:          4,
			UserName:    "todo",
			Description: "Learn Deno",
			TargetDate:  "2025-09-15",
			Done:        false,
		},
	}
}
