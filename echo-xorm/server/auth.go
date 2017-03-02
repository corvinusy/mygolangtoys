package server

import (
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/go-xorm/xorm"
	"github.com/labstack/echo"
)

// authHandler represents handler for '/auth'
type authHandler struct {
	Orm *xorm.Engine
}

type versionResponse struct {
	ServerTime uint64 `json:"server_time"`
	Version    string `json:"version"`
}

// handler for /version
func (h *authHandler) getVersion(c echo.Context) error {
	vr := versionResponse{
		ServerTime: uint64(time.Now().Unix()),
		Version:    time.Now().String(),
	}
	return c.JSON(http.StatusOK, vr)
}

type authResponse struct {
	Result string `json:"result"`
	Token  string `json:"token"`
}

// handler for /auth
func (h *authHandler) getToken(c echo.Context) error {
	var (
		input userInput
		user  User
		err   error
		found bool
	)

	if err = c.Bind(&input); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	// find user
	user = User{Login: input.Login}
	found, err = h.Orm.Get(&user)
	if err != nil || !found {
		return c.String(http.StatusForbidden, "invalid credentials")
	}

	//validate user credentials
	if input.Password != user.Password {
		return c.String(http.StatusForbidden, "invalid credentials")
	}

	//create a HMAC SHA256 signer
	token := jwt.New(jwt.SigningMethodHS256)

	//set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["iss"] = "Jon Snow"
	claims["iat"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	//signer.Claims["jti"] = "1" // should be what?
	claims["CustomUserInfo"] = input.Login

	t, err := token.SignedString(signingKey)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error while signing the token:"+err.Error())
	}

	resp := authResponse{
		Result: "OK",
		Token:  t,
	}
	return c.JSON(http.StatusOK, resp)
}
