package api

import (
	"database/sql"
	"log"

	"code.google.com/p/go.crypto/bcrypt"

	"github.com/gin-gonic/gin"
	"gopkg.in/redis.v2"
)

type Controller struct {
}

func (controller *Controller) GetDatabase(c *gin.Context) *sql.DB {
	db, _ := c.Get("Database")
	return db.(*sql.DB)
}

func (controller *Controller) GetStorage(c *gin.Context) *redis.Client {
	storage, _ := c.Get("Storage")
	return storage.(*redis.Client)
}

func (controller *Controller) GetCache(c *gin.Context) *redis.Client {
	cache, _ := c.Get("Cache")
	return cache.(*redis.Client)
}

func (controller *Controller) HashPassword(password string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Couldn't hash password: %v", err)
		return nil, err
	}
	return hash, nil
}

func (controller *Controller) IsAuthentic(hash []byte, password string) bool {
	err := bcrypt.CompareHashAndPassword(hash, []byte(password))
	if err != nil {
		return false
	}
	return true
}
