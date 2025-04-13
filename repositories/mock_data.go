package repositories

import (
	"go-todos-api/models"
	"time"
)

func GetUserListMockData() []models.User {
	return []models.User{
		{
			ID:        1,
			Name:      "Lam",
			BirthDay:  "2000-01-01",
			Gender:    "Male",
			PhotoURL:  "",
			LastLogin: time.Now().Add(-2),
			Active:    true,
			UpdatedAt: time.Now().Add(-10),
		},
		{
			ID:        2,
			Name:      "Lam",
			BirthDay:  "2001-01-15",
			Gender:    "Male",
			PhotoURL:  "",
			LastLogin: time.Now().Add(-2),
			Active:    true,
			UpdatedAt: time.Now().Add(-10),
		},
		{
			ID:        3,
			Name:      "Lam",
			BirthDay:  "2002-03-02",
			Gender:    "Male",
			PhotoURL:  "",
			LastLogin: time.Now().Add(-2),
			Active:    true,
			UpdatedAt: time.Now().Add(-10),
		},
	}
}

func GetTodoListMockData() []models.Todo {
	return []models.Todo{
		{
			ID:          1,
			Username:    "todo",
			Description: "Learn AWS",
			TargetDate:  time.Now().AddDate(0, 1, 5),
			Done:        false,
		},
		{
			ID:          2,
			Username:    "todo",
			Description: "Learn Azure",
			TargetDate:  time.Now().AddDate(0, 2, 10),
			Done:        false,
		},
		{
			ID:          3,
			Username:    "todo",
			Description: "Learn DevOp",
			TargetDate:  time.Now().AddDate(0, 3, 15),
			Done:        false,
		},
		{
			ID:          4,
			Username:    "todo",
			Description: "Learn Deno",
			TargetDate:  time.Now().AddDate(0, 4, 22),
			Done:        false,
		},
	}
}
