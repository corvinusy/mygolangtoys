package api

import (
	"database/sql"

	"github.com/zenazn/goji/web"
	"gopkg.in/redis.v2"
)

type Controller struct {
}

func (controller *Controller) GetDatabase(c web.C) *sql.DB {
	db := c.Env["Database"].(*sql.DB)
	return db
}

func (controller *Controller) GetStorage(c web.C) *redis.Client {
	storage := c.Env["Storage"].(*redis.Client)
	return storage
}

func (controller *Controller) GetCache(c web.C) *redis.Client {
	cache := c.Env["Cache"].(*redis.Client)
	return cache
}
