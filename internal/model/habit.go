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

func (c HabitModel) getHabitsByWeekday(weekday uint) []entity.Habit {
	weekdaysBits := []uint{1, 2, 4, 8, 16, 32, 64}
	weekdayBit := weekdaysBits[weekday]
	habits := []entity.Habit{}
	DB.Where("(Days & ?) = ?", weekdayBit, weekdayBit).Find(&habits)
	return habits
}
