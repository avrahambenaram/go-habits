package controller

import (
	"net/http"
	"strconv"

	"github.com/avrahambenaram/go-habits/internal/entity"
	"github.com/avrahambenaram/go-habits/internal/model"
)

type HabitController struct {
	server     *http.ServeMux
	habitModel *model.HabitModel
}

func NewHabitController(server *http.ServeMux, habitModel *model.HabitModel) HabitController {
	habitController := HabitController{
		server,
		habitModel,
	}
	server.HandleFunc("POST /habit", habitController.Create)

	return habitController
}

func (c HabitController) Create(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	title := r.PostFormValue("title")
	weekDaysForm := r.Form["week_days"]
	weekDays := 0
	for _, day := range weekDaysForm {
		dayInt, err := strconv.Atoi(day)
		if err == nil {
			weekDays += dayInt
		}
	}
	habit := entity.Habit{
		Title: title,
		Days:  uint(weekDays),
	}

	errCreate := c.habitModel.Create(habit)
	if errCreate != nil {
		w.Write([]byte("Erro"))
		return
	}

	w.Write([]byte("Teste"))
}
