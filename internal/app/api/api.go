package api

import "github.com/matigumma/go-fiber-sql-hexagonal/internal/ports"

type Application struct {
	db ports.DbPort
}

func NewApplication(db ports.DbPort) *Application {
	return &Application{db: db}
}

