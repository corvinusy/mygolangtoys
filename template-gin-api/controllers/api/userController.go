package api

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"vv-gin-api/models"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	_ "github.com/lib/pq"
)

func (ct *Controller) UserRegisterHandler(c *gin.Context) ([]byte, int) {

	type (
		okResponse struct {
			Res    string `json:"result"`
			UserID int    `json:"user_id"`
			Rid    int    `json:"rid"`
		}
		registerForm struct {
			Login    string `form:"login" binding:"required"`
			Name     string `form:"name" binding:"required"`
			Password string `form:"password" binding:"required"`
		}
	)

	var (
		user   models.User
		params registerForm
		resp   okResponse
		db     *sql.DB
		err    error
		hash   []byte
	)

	c.BindWith(&params, binding.Form)

	if len(params.Login) < 4 || len(params.Login) > 31 {
		return []byte("Login parameter error\n" + params.Login), http.StatusNotAcceptable
	}

	if len(params.Name) < 4 || len(params.Name) > 15 {
		return []byte("Name parameter error\n" + params.Name), http.StatusNotAcceptable
	}

	hash, err = ct.HashPassword(params.Password)
	if err != nil {
		return []byte("Password parameter error"), http.StatusInternalServerError
	}

	//create user with db
	db = ct.GetDatabase(c)
	err = ct.CreateUser(db, params.Login, params.Name, hash)
	if err != nil {
		return []byte(err.Error()), http.StatusNotAcceptable
	}

	resp = okResponse{"OK", user.UserId, 0}
	body, err := json.MarshalIndent(resp, "", "\t")
	if err != nil {
		return []byte("Json error"), http.StatusInternalServerError
	}

	return body, http.StatusOK
}

func (ct *Controller) UserLoginHandler(c *gin.Context) ([]byte, int) {
	type (
		okResponse struct {
			Res     string `json:"result"`
			Token   int    `json:"token"`
			Expired int    `json:"expired"`
			Rid     int    `json:"rid"`
		}
		loginForm struct {
			Login    string `form:"login" binding:"required"`
			Password string `form:"password" binding:"required"`
		}
	)

	var (
		user   models.User
		params loginForm
		resp   okResponse
		db     *sql.DB
		err    error
	)

	c.BindWith(&params, binding.Form)

	db = ct.GetDatabase(c)
	hash := user.GetUserAuth(db, paramsLogin)

	return nil, http.StatusOK
}

func (ct *Controller) UserDestroy(c *gin.Context) ([]byte, int) {
	// stub
	return nil, http.StatusOK
}
