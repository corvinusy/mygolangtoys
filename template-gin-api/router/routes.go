package router

import (
	"vv-gin-api/controllers/api"

	"github.com/gin-gonic/gin"
)

type (
	Router struct{}

	// shortcut types
	controllerHandler func(*gin.Context) ([]byte, int)
)

// main router
func (r *Router) Route(handler controllerHandler) gin.HandlerFunc {
	// neat closure
	fn := func(c *gin.Context) {
		body, code := handler(c)
		c.String(code, "%s", body)
	}
	return fn
}

func (r *Router) ApplyRoutes(e *gin.Engine) {

	var ct = &api.Controller{}

	// root
	e.GET("/", r.Route(ct.RootHandler))

	// user
	e.GET("/user/register", r.Route(ct.UserRegisterHandler))
	e.GET("/user/login", r.Route(ct.UserLoginHandler))
	e.GET("/user/logout", r.Route(ct.UserLogoutHandler))
	e.GET("/user/destroy", r.Route(ct.UserDestroyHandler))
}
