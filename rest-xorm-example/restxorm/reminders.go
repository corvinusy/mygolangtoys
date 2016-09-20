package restxorm

import (
	"crypto/md5"
	"encoding/hex"
	"net/http"
	"strconv"
	"time"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/go-xorm/xorm"
)

// Reminder is a resource for /reminders requests
type Reminder struct {
	ID      int64     `xorm:"'id'" json:"id"`
	Message string    `xorm:"varchar(1024)" json:"message"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
	Hash    string    `xorm:"varchar(32)" json:"hash"`
}

// TableName uses by xorm to realize name of table
func (r *Reminder) TableName() string {
	return "reminders"
}

// ReminderImpl is ORM database mapping implementation
type ReminderImpl struct {
	Engine *xorm.Engine
}

// GetAllReminders is a GET /reminders ending
func (i *ReminderImpl) GetAllReminders(w rest.ResponseWriter, r *rest.Request) {
	reminders := []Reminder{}
	i.Engine.Find(&reminders)
	w.WriteJson(&reminders)
}

// GetReminder is a GET /reminders/{id} ending
func (i *ReminderImpl) GetReminder(w rest.ResponseWriter, r *rest.Request) {
	var (
		id       int64
		err      error
		reminder Reminder
		found    bool
	)

	id, err = strconv.ParseInt(r.PathParam("id"), 10, 0)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	reminder = Reminder{ID: id}
	found, err = i.Engine.Get(&reminder)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !found {
		rest.NotFound(w, r)
		return
	}
	w.WriteJson(&reminder)
}

// PostReminder is a POST /reminders ending
func (i *ReminderImpl) PostReminder(w rest.ResponseWriter, r *rest.Request) {
	var err error
	reminder := Reminder{}
	if err = r.DecodeJsonPayload(&reminder); err != nil {
		rest.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	reminder.Hash = i.getHash(&reminder)

	affected, err := i.Engine.Insert(&reminder)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if affected == 0 {
		rest.Error(w, "Always exists", http.StatusConflict)
		return
	}

	w.WriteJson(&reminder)
}

// PutReminder is a PUT /reminders/{id} ending
func (i *ReminderImpl) PutReminder(w rest.ResponseWriter, r *rest.Request) {

	var (
		id, affected      int64
		err               error
		found             bool
		reminder, updated Reminder
	)

	id, err = strconv.ParseInt(r.PathParam("id"), 10, 0)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	reminder = Reminder{ID: id}
	found, err = i.Engine.Id(id).Get(&reminder)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !found {
		rest.NotFound(w, r)
		return
	}

	updated = Reminder{}
	if err = r.DecodeJsonPayload(&updated); err != nil {
		rest.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updated.Hash = i.getHash(&updated)

	affected, err = i.Engine.Id(id).Update(&updated)

	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if affected == 0 {
		rest.Error(w, "Always exists", http.StatusConflict)
		return
	}

	// assemble response struct
	reminder.Message = updated.Message
	reminder.Updated = updated.Updated
	reminder.Hash = updated.Hash
	w.WriteJson(&reminder)
}

// DeleteReminder is a DELETE /reminders/{id} ending
func (i *ReminderImpl) DeleteReminder(w rest.ResponseWriter, r *rest.Request) {

	var (
		id, affected int64
		err          error
		found        bool
		reminder     Reminder
	)

	id, err = strconv.ParseInt(r.PathParam("id"), 10, 0)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	reminder = Reminder{ID: id}
	found, err = i.Engine.Get(&reminder)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !found {
		rest.NotFound(w, r)
		return
	}

	affected, err = i.Engine.Delete(&reminder)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if affected == 0 {
		rest.Error(w, "Cannot be deleted", http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// getHash returns Hash of record
func (i *ReminderImpl) getHash(reminder *Reminder) string {
	sumByteArray := md5.Sum([]byte(reminder.Message))
	sumString := hex.EncodeToString(sumByteArray[:])
	return sumString
}
