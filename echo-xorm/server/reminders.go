package server

import (
	"net/http"
	"strconv"
	"time"

	"github.com/go-xorm/xorm"
	"github.com/labstack/echo"
)

// Reminder is a resource for /reminders requests
type Reminder struct {
	ID      uint64    `xorm:"'id' pk autoincr unique notnull" json:"id"`
	Message string    `xorm:"text" json:"message"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
	Hash    string    `xorm:"text" json:"hash"`
}

// TableName used by xorm to set table name for entity
func (r *Reminder) TableName() string {
	return "reminders"
}

// reminderHandler is a container for handlers and app data
type reminderHandler struct {
	Orm *xorm.Engine
}

// FindAllReminders is a GET /reminders handler
func (h *reminderHandler) FindAllReminders(c echo.Context) error {
	reminders := []Reminder{}
	err := h.Orm.Find(&reminders)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, reminders)
}

// FindReminder is a GET /reminders/{id} handler
func (h *reminderHandler) FindReminder(c echo.Context) error {
	var reminder Reminder

	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	found, err := h.Orm.Id(id).Get(&reminder)
	if err != nil {
		return c.String(http.StatusServiceUnavailable, err.Error())
	}
	if !found {
		return c.NoContent(http.StatusNoContent)
	}
	return c.JSON(http.StatusOK, reminder)
}

// CreateReminder is a POST /reminders handler
func (h *reminderHandler) CreateReminder(c echo.Context) error {
	var (
		err      error
		reminder Reminder
	)

	if err = c.Bind(&reminder); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	reminder.Hash = getSHA3Hash(reminder.Message)
	if err != nil {
		return err
	}

	affected, err := h.Orm.InsertOne(&reminder)
	if err != nil {
		return c.String(http.StatusServiceUnavailable, err.Error())
	}
	if affected == 0 {
		return c.String(http.StatusConflict, err.Error())
	}

	return c.JSON(http.StatusCreated, reminder)
}

// UpdateReminder is a PUT /reminders/{id} handler
func (h *reminderHandler) UpdateReminder(c echo.Context) error {
	var input, reminder Reminder
	// parse id
	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	// get reminder from db
	found, err := h.Orm.Id(id).Get(&reminder)
	if err != nil || !found {
		return c.NoContent(http.StatusNoContent)
	}
	// parse request body
	if err = c.Bind(&input); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	// update in db
	input.Hash = getSHA3Hash(input.Message)
	if err != nil {
		return err
	}
	affected, err := h.Orm.Id(id).Update(&input)
	if err != nil {
		return c.String(http.StatusServiceUnavailable, err.Error())
	}
	if affected == 0 {
		return c.String(http.StatusConflict, err.Error())
	}

	// assemble response struct
	reminder = Reminder{}
	found, err = h.Orm.Id(id).Get(&reminder)
	if err != nil || !found {
		return c.String(http.StatusServiceUnavailable, err.Error())
	}

	return c.JSON(http.StatusOK, reminder)
}

// DeleteReminder is a DELETE /reminders/{id} ending
func (h *reminderHandler) DeleteReminder(c echo.Context) error {

	var (
		id       uint64
		affected int64
		err      error
		reminder Reminder
	)

	id, err = strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	// get reminder from db
	found, err := h.Orm.Id(id).Get(&reminder)
	if err != nil || !found {
		return c.NoContent(http.StatusNoContent)
	}

	affected, err = h.Orm.Id(id).Delete(&Reminder{})
	if err != nil {
		return c.String(http.StatusServiceUnavailable, err.Error())
	}
	if affected == 0 {
		return c.String(http.StatusConflict, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}
