package server

import (
	"net/http"
	"strconv"
	"time"

	"github.com/go-xorm/xorm"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

// User is a resource for /users requests
type User struct {
	ID       uint64    `xorm:"'id' pk autoincr unique notnull" json:"id"`
	Login    string    `xorm:"text" json:"login"`
	Password string    `xorm:"text" json:"-"`
	Created  time.Time `xorm:"created"`
	Updated  time.Time `xorm:"updated"`
	Hash     string    `xorm:"'hash' text" json:"hash"`
}

// userInput represents payload data format
type userInput struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

// TableName used by xorm to set table name for entity
func (r *User) TableName() string {
	return "users"
}

// userHandler is a container for handlers and app data
type userHandler struct {
	Orm *xorm.Engine
}

// FindAllUsers is a GET /users handler
func (h *userHandler) FindAllUsers(c echo.Context) error {
	users := []User{}
	err := h.Orm.Find(&users)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, users)
}

// FindUser is a GET /users/{id} handler
func (h *userHandler) FindUser(c echo.Context) error {
	var user User

	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	found, err := h.Orm.Id(id).Get(&user)
	if err != nil {
		return c.String(http.StatusServiceUnavailable, err.Error())
	}
	if !found {
		return c.NoContent(http.StatusNoContent)
	}
	return c.JSON(http.StatusOK, user)
}

// CreateUser is a POST /users handler
func (h *userHandler) CreateUser(c echo.Context) error {
	var (
		err   error
		user  User
		input userInput
	)

	if err = c.Bind(&input); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	// encrypt password
	passHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	hash := getSHA3Hash(input.Login)
	if err != nil {
		return err
	}
	// construct user
	user = User{
		Login:    input.Login,
		Password: string(passHash),
		Hash:     hash,
	}

	affected, err := h.Orm.InsertOne(&user)
	if err != nil {
		return c.String(http.StatusServiceUnavailable, err.Error())
	}
	if affected == 0 {
		return c.String(http.StatusConflict, err.Error())
	}

	return c.JSON(http.StatusCreated, user)
}

// UpdateUser is a PUT /users/{id} handler
func (h *userHandler) UpdateUser(c echo.Context) error {
	var (
		input userInput
		user  User
		id    uint64
		err   error
	)
	// parse id
	id, err = strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	// get user from db
	found, err := h.Orm.Id(id).Get(&user)
	if err != nil {
		return c.String(http.StatusServiceUnavailable, err.Error())
	}
	if !found {
		return c.NoContent(http.StatusNoContent)
	}
	// parse request body
	if err = c.Bind(&input); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	// encrypt password
	passHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	hash := getSHA3Hash(input.Login)
	if err != nil {
		return err
	}
	// construct user
	user = User{
		Login:    input.Login,
		Password: string(passHash),
		Hash:     hash,
	}

	// update
	affected, err := h.Orm.Id(id).Update(&user)
	if err != nil {
		return c.String(http.StatusServiceUnavailable, err.Error())
	}
	if affected == 0 {
		return c.String(http.StatusConflict, err.Error())
	}

	// assemble response struct
	user = User{}
	found, err = h.Orm.Id(id).Get(&user)
	if err != nil {
		return c.String(http.StatusServiceUnavailable, err.Error())
	}
	if !found {
		return c.NoContent(http.StatusNoContent)
	}

	return c.JSON(http.StatusOK, user)
}

// DeleteUser is a DELETE /users/{id} ending
func (h *userHandler) DeleteUser(c echo.Context) error {

	var (
		id       uint64
		affected int64
		err      error
		user     User
	)

	id, err = strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	// get user from db
	found, err := h.Orm.Id(id).Get(&user)
	if err != nil || !found {
		return c.NoContent(http.StatusNoContent)
	}

	affected, err = h.Orm.Id(id).Delete(&User{})
	if err != nil {
		return c.String(http.StatusServiceUnavailable, err.Error())
	}
	if affected == 0 {
		return c.String(http.StatusConflict, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}
