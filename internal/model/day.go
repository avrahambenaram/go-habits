package model

import (
	"time"

	"github.com/avrahambenaram/go-habits/internal/entity"
)

type DayModel struct {
	HabitModel *HabitModel
}

type TableItem struct {
	entity.Day
	Points   float32
	Editable bool
}
type Table []TableItem

func (c *Table) Add(item TableItem) {
	*c = append(*c, item)
}

func (c DayModel) GetTable() Table {
	now := time.Now()
	today := now.Truncate(24 * time.Hour)
	after10Days := today.Add(10 * 24 * time.Hour)
	before := time.Date(
		today.Year(),
		today.Month(),
		today.Day()-27-int(now.Weekday()),
		0, 0, 0, 0,
		today.Location(),
	)
	table := Table{}

	days := []entity.Day{}
	for current := before; current.Unix() <= after10Days.Unix(); current = current.Add(24 * time.Hour) {
		day := entity.Day{}
		DB.FirstOrCreate(&day, entity.Day{
			Date: current,
		})
		days = append(days, day)
	}
	for _, day := range days {
		dayHabits := []entity.DayHabit{}
		weekday := uint(day.Date.Weekday())
		weekdayHabits := c.HabitModel.getHabitsByWeekday(weekday)
		DB.Where("day_id = ?", day.ID).Find(&dayHabits)
		var points float32 = 0
		if len(weekdayHabits) != 0 {
			points = float32(len(dayHabits)) / float32(len(weekdayHabits))
		}
		editable := false
		hours3 := 3600 * 3
		// FIX: UTC-0300 conflict
		if day.Date.Unix() == today.Unix()+int64(hours3) {
			editable = true
		}
		table.Add(TableItem{
			day,
			points,
			editable,
		})
	}
	return table
}
