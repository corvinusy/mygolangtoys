package core

import (
	"net/http"

	"github.com/zenazn/goji/web"
)

// forwarding database access to controllers
func (application *Application) ApplyDatabase(c *web.C, h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		c.Env["Database"] = application.Database
		c.Env["DbVersion"] = application.DbVersion
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

// forwarding storage access to controllers
func (application *Application) ApplyStorage(c *web.C, h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		c.Env["Storage"] = application.Storage
		c.Env["ApiVersion"] = application.ApiVersion
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

// forwarding cache access to controllers
func (application *Application) ApplyCache(c *web.C, h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		c.Env["Cache"] = application.Cache
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
