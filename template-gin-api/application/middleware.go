package application

import "github.com/gin-gonic/gin"

// forwarding database access to controllers
func (app *Application) ApplyDatabase(c *gin.Context) {
	c.Set("Database", app.Database)
	c.Set("DbVersion", app.DbVersion)
}

// forwarding storage access to controllers
func (app *Application) ApplyStorage(c *gin.Context) {
	c.Set("Storage", app.Storage)
	c.Set("ApiVersion", app.ApiVersion)
}

// forwarding cache access to controllers
func (app *Application) ApplyCache(c *gin.Context) {
	c.Set("Cache", app.Cache)
}
