package entity

import "time"

type Day struct {
	ID   uint `gorm:"primarykey"`
	Date time.Time
}
