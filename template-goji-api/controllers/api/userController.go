package api

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/corvinusy/vv-goji-api/models"

	_ "github.com/lib/pq"
	"github.com/zenazn/goji/web"
)

type UserRegisterOK struct {
	Res    string `json:"result"`
	UserID int    `json:"user_id"`
	Rid    int    `json:"rid"`
}

func (ct *Controller) UserRegister(c web.C, r *http.Request) ([]byte, int) {

	var (
		user   models.User
		userOK UserRegisterOK
		db     *sql.DB
		err    error
	)

	user = models.User{UserLogin: c.URLParams["login"], UserName: c.URLParams["name"]}

	loginLen := len(user.UserLogin)
	if loginLen < 4 || loginLen > 31 {
		return []byte("Login parameter error\n" + user.UserLogin), http.StatusNotAcceptable
	}

	nameLen := len(user.UserName)
	if nameLen < 4 || nameLen > 15 {
		return []byte("Name parameter error"), http.StatusNotAcceptable
	}

	hash, err := user.HashPassword(c.URLParams["pass"])
	if err != nil {
		return []byte("Password parameter error"), http.StatusInternalServerError
	}

	db = ct.GetDatabase(c)
	err = user.Create(db, hash)

	if err != nil || user.UserId == 0 {
		return nil, http.StatusNotAcceptable
	}

	userOK = UserRegisterOK{"OK", user.UserId, 0}
	body, err := json.MarshalIndent(userOK, "", "\t")
	if err != nil {
		return []byte("Json error"), http.StatusInternalServerError
	}

	return body, http.StatusOK
}

func (ct *Controller) UserLogin(c web.C, r *http.Request) ([]byte, int) {
	// stub
	type ver struct {
		result  string
		version string
		rid     int
	}
	result := &ver{result: "OK", version: "3.0", rid: 0}
	body, err := json.Marshal(result)
	if err != nil {
		return []byte("Json error"), http.StatusInternalServerError
	}
	return body, http.StatusOK
}

func (ct *Controller) UserDestroy(c web.C, r *http.Request) ([]byte, int) {
	// stub
	type ver struct {
		result  string
		version string
		rid     int
	}
	result := &ver{result: "OK", version: "3.0", rid: 0}
	body, err := json.Marshal(result)
	if err != nil {
		return []byte("Json error"), http.StatusInternalServerError
	}
	return body, http.StatusOK
}
