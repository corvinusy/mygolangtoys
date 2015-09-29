package router

import (
	"net/http"

	"github.com/corvinusy/vv-goji-api/controllers/api"

	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

type Router struct {
}

// json-responder-helper
func JsonResponse(w http.ResponseWriter, body []byte, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(body)
}

// main router
// returns func(c web.C, w http.ResponseWriter, r *http.Request)
func (r *Router) Route(controller interface{}) interface{} {
	// neat closure
	fn := func(c web.C, w http.ResponseWriter, r *http.Request) {

		handler := controller.(func(c web.C, r *http.Request) ([]byte, int))
		body, code := handler(c, r)
		JsonResponse(w, body, code)
	}
	return fn
}

func (r *Router) ApplyRoutes() {

	var ct = &api.Controller{}

	// root
	goji.Get("/", r.Route(ct.RootVersion))

	// user register
	goji.Get("/user/register/:login/:name/:pass", r.Route(ct.UserRegister))

	/*	// user login
		goji.Get("/user/login/:login,:pass", application.Route("User.Login"))

		// user destroy
		goji.Get("/user/destroy/:token", application.Route("User.Destroy"))
	*/
}
