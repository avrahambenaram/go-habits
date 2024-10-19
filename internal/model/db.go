package model

import (
	"github.com/avrahambenaram/go-habits/internal/entity"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	db, err := gorm.Open(sqlite.Open("./habits.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	db.AutoMigrate(&entity.Habit{})
	db.AutoMigrate(&entity.Day{})
	db.AutoMigrate(&entity.DayHabit{})
	DB = db
}
