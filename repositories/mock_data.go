package repositories

import (
	"go-todos-api/models"
	"go-todos-api/pkg/types"
	"time"
)

func GetUserListMockData() []models.User {
	return []models.User{
		{
			ID:        1,
			Name:      "Test User 1",
			Username:  "todo",
			Email:     "test1@example.com",
			BirthDay:  "2000-01-01",
			Gender:    "Male",
			PhotoURL:  "",
			LastLogin: time.Now().Add(-2),
			Active:    true,
			UpdatedAt: time.Now().Add(-10),
			Password:  "f7e726f3d70567c06772b3b32a4c5bfa4dca451e014aa2eecd3e575a8b12091f",
		},
		{
			ID:        2,
			Name:      "Test User 2",
			Username:  "aha",
			Email:     "test2@example.com",
			BirthDay:  "2001-01-15",
			Gender:    "Male",
			PhotoURL:  "",
			LastLogin: time.Now().Add(-2),
			Active:    true,
			UpdatedAt: time.Now().Add(-10),
			Password:  "f7e726f3d70567c06772b3b32a4c5bfa4dca451e014aa2eecd3e575a8b12091f",
		},
		{
			ID:        3,
			Name:      "Test User 13",
			Username:  "oho",
			Email:     "test3@example.com",
			BirthDay:  "2002-03-02",
			Gender:    "Male",
			PhotoURL:  "",
			LastLogin: time.Now().Add(-2),
			Active:    true,
			UpdatedAt: time.Now().Add(-10),
			Password:  "f7e726f3d70567c06772b3b32a4c5bfa4dca451e014aa2eecd3e575a8b12091f",
		},
	}
}

func GetTodoListMockData() []models.Todo {
	return []models.Todo{
		{
			ID:          1,
			UserID:      1,
			Description: "Learn AWS",
			TargetDate:  types.DateType{Time: time.Now().AddDate(0, 1, 5)},
			Done:        false,
		},
		{
			ID:          2,
			UserID:      1,
			Description: "Learn Azure",
			TargetDate:  types.DateType{Time: time.Now().AddDate(0, 2, 10)},
			Done:        false,
		},
		{
			ID:          3,
			UserID:      1,
			Description: "Learn DevOp",
			TargetDate:  types.DateType{Time: time.Now().AddDate(0, 3, 15)},
			Done:        false,
		},
		{
			ID:          4,
			UserID:      1,
			Description: "Learn Deno",
			TargetDate:  types.DateType{Time: time.Now().AddDate(0, 4, 22)},
			Done:        false,
		},
		{
			ID:          5,
			UserID:      2,
			Description: "Learn AWS",
			TargetDate:  types.DateType{Time: time.Now().AddDate(0, 1, 5)},
			Done:        false,
		},
		{
			ID:          6,
			UserID:      2,
			Description: "Learn Azure",
			TargetDate:  types.DateType{Time: time.Now().AddDate(0, 2, 10)},
			Done:        false,
		},
		{
			ID:          7,
			UserID:      2,
			Description: "Learn DevOp",
			TargetDate:  types.DateType{Time: time.Now().AddDate(0, 3, 15)},
			Done:        false,
		},
		{
			ID:          8,
			UserID:      2,
			Description: "Learn Deno",
			TargetDate:  types.DateType{Time: time.Now().AddDate(0, 4, 22)},
			Done:        false,
		},
		{
			ID:          9,
			UserID:      3,
			Description: "Learn AWS",
			TargetDate:  types.DateType{Time: time.Now().AddDate(0, 1, 5)},
			Done:        false,
		},
		{
			ID:          10,
			UserID:      3,
			Description: "Learn Azure",
			TargetDate:  types.DateType{Time: time.Now().AddDate(0, 2, 10)},
			Done:        false,
		},
		{
			ID:          11,
			UserID:      3,
			Description: "Learn DevOp",
			TargetDate:  types.DateType{Time: time.Now().AddDate(0, 3, 15)},
			Done:        false,
		},
		{
			ID:          12,
			UserID:      3,
			Description: "Learn Deno",
			TargetDate:  types.DateType{Time: time.Now().AddDate(0, 4, 22)},
			Done:        false,
		},
	}
}
