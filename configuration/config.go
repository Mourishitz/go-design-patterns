package configuration

import (
	"database/sql"
	"sync"

	"dpatterns/models"
)

type Application struct {
	Models *models.Models
}

var (
	instance *Application
	once     sync.Once
	db       *sql.DB
)

func New(pool *sql.DB) *Application {
	db = pool
	return GetInstance()
}

func GetInstance() *Application {
	once.Do(func() {
		instance = &Application{
			Models: models.New(db),
		}
	})
	return instance
}
