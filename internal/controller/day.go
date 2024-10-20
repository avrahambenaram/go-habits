package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/avrahambenaram/go-habits/internal/model"
)

type DayController struct {
	server   *http.ServeMux
	tmpls    *template.Template
	dayModel *model.DayModel
}

type Item struct {
	model.TableItem
	Date    string
	Weekday string
}

func NewDayController(server *http.ServeMux, dayModel *model.DayModel, tmpls *template.Template) DayController {
	dayController := DayController{
		server,
		tmpls,
		dayModel,
	}
	server.HandleFunc("GET /table", dayController.Table)
	server.HandleFunc("GET /day/habits/{ID}", dayController.DayHabits)

	return dayController
}

func (c *DayController) DayHabits(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("ID")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Id deve ser um número inteiro", http.StatusForbidden)
		return
	}

	items, err1 := c.dayModel.FindHabitsByDayID(id)
	if err1 != nil {
		http.Error(w, "Id não encontrado", http.StatusNotFound)
		return
	}

	now := time.Now()
	editingDay, _ := c.dayModel.FindById(id)
	editingDate := editingDay.Date
	isEditable := false
	if editingDate.Year() == now.Year() && editingDate.Month() == now.Month() && editingDate.Day() == now.Day() {
		isEditable = true
	}

	c.tmpls.ExecuteTemplate(w, "day-habits", map[string]interface{}{
		"HabitItems": items,
		"Editable":   isEditable,
	})
}

func (c *DayController) Table(w http.ResponseWriter, r *http.Request) {
	table := c.dayModel.GetTable()
	items := []Item{}
	c.formatItems(&items, table)

	c.tmpls.ExecuteTemplate(w, "table", map[string]interface{}{
		"Days": items,
	})
}

func (c DayController) formatItems(dest *[]Item, table model.Table) {
	for _, tableItem := range table {
		date := fmt.Sprintf("%02d/%02d", tableItem.Date.Day(), tableItem.Date.Month())
		weekdays := []string{
			"Domingo",
			"Segunda-feira",
			"Terça-feira",
			"Quarta-feira",
			"Quinta-feira",
			"Sexta-feira",
			"Sábado",
		}
		weekday := weekdays[tableItem.Date.Weekday()]
		item := Item{
			tableItem,
			date,
			weekday,
		}
		*dest = append(*dest, item)
	}
}
