package entity

type DayHabit struct {
	ID      uint `gorm:"primarykey"`
	DayID   uint
	HabitID uint

	Day
	Habit
}
