package entity

import (
	"gorm.io/gorm"
)

type Habit struct {
	gorm.Model
	Title string
	Days  uint
}
