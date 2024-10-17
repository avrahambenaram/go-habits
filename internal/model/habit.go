package model

import (
	"errors"

	"github.com/avrahambenaram/go-habits/internal/entity"
)

type HabitModel struct {
}

func (c HabitModel) Create(habit entity.Habit) error {
	result := DB.Create(&habit)
	if result.RowsAffected != 1 {
		return errors.New("Failed to create new habit")
	}
	return nil
}
