package server

import (
	"crypto/md5"
	"encoding/hex"
	"net/http"
	"strconv"
	"time"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/go-xorm/xorm"
)

// User is a resource for /users requests
type User struct {
	ID      int64     `xorm:"'id'" json:"id"`
	Message string    `xorm:"varchar(1024)" json:"message"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
	Hash    string    `xorm:"varchar(32)" json:"hash"`
}

// TableName uses by xorm to realize name of table
func (u *User) TableName() string {
	return "users"
}

// UserHandler is handler implementation for users
type UserHandler struct {
	Engine *xorm.Engine
}

// GetAllUsers is a GET /users ending
func (i *UserHandler) GetAllUsers(w rest.ResponseWriter, r *rest.Request) {
	users := []User{}
	i.Engine.Find(&users)
	w.WriteJson(&users)
}

// GetUser is a GET /users/{id} ending
func (i *UserHandler) GetUser(w rest.ResponseWriter, r *rest.Request) {
	var (
		id    int64
		err   error
		user  User
		found bool
	)

	id, err = strconv.ParseInt(r.PathParam("id"), 10, 0)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user = User{ID: id}
	found, err = i.Engine.Get(&user)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !found {
		rest.NotFound(w, r)
		return
	}
	w.WriteJson(&user)
}

// PostUser is a POST /users ending
func (i *UserHandler) PostUser(w rest.ResponseWriter, r *rest.Request) {
	var err error
	user := User{}
	if err = r.DecodeJsonPayload(&user); err != nil {
		rest.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user.Hash = i.getHash(&user)

	affected, err := i.Engine.Insert(&user)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if affected == 0 {
		rest.Error(w, "Always exists", http.StatusConflict)
		return
	}

	w.WriteJson(&user)
}

// PutUser is a PUT /users/{id} ending
func (i *UserHandler) PutUser(w rest.ResponseWriter, r *rest.Request) {

	var (
		id, affected  int64
		err           error
		found         bool
		user, updated User
	)

	id, err = strconv.ParseInt(r.PathParam("id"), 10, 0)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user = User{ID: id}
	found, err = i.Engine.Id(id).Get(&user)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !found {
		rest.NotFound(w, r)
		return
	}

	updated = User{}
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
	user.Message = updated.Message
	user.Updated = updated.Updated
	user.Hash = updated.Hash
	w.WriteJson(&user)
}

// DeleteUser is a DELETE /users/{id} ending
func (i *UserHandler) DeleteUser(w rest.ResponseWriter, r *rest.Request) {

	var (
		id, affected int64
		err          error
		found        bool
		user         User
	)

	id, err = strconv.ParseInt(r.PathParam("id"), 10, 0)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user = User{ID: id}
	found, err = i.Engine.Get(&user)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !found {
		rest.NotFound(w, r)
		return
	}

	affected, err = i.Engine.Delete(&user)
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
func (i *UserHandler) getHash(user *User) string {
	sumByteArray := md5.Sum([]byte(user.Message))
	sumString := hex.EncodeToString(sumByteArray[:])
	return sumString
}
