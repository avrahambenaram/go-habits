package controller

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/avrahambenaram/go-habits/internal/entity"
	"github.com/avrahambenaram/go-habits/internal/model"
)

type HabitController struct {
	server     *http.ServeMux
	tmpls      *template.Template
	habitModel *model.HabitModel
}

func NewHabitController(server *http.ServeMux, habitModel *model.HabitModel, tmpls *template.Template) HabitController {
	habitController := HabitController{
		server,
		tmpls,
		habitModel,
	}
	server.HandleFunc("POST /habit", habitController.Create)

	return habitController
}

func (c *HabitController) Create(w http.ResponseWriter, r *http.Request) {
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
		if err != nil {
			break
		}
		if dayInt < 0 || dayInt > 6 {
			break
		}
		dayMask := 1 << dayInt
		weekDays |= dayMask
	}
	habit := entity.Habit{
		Title: title,
		Days:  uint(weekDays),
	}

	errCreate := c.habitModel.Create(habit)
	if errCreate != nil {
		w.WriteHeader(400)
		c.tmpls.ExecuteTemplate(w, "result-fail", "Erro ao cadastrar o hábito")
		return
	}

	c.tmpls.ExecuteTemplate(w, "result-success", "Cadastrado com sucesso")
}
