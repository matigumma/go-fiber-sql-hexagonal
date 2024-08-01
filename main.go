package main

import (
	"log"

	"github.com/matigumma/go-fiber-sql-hexagonal/internal/adapters/left/rest"
	"github.com/matigumma/go-fiber-sql-hexagonal/internal/adapters/right/db"
	"github.com/matigumma/go-fiber-sql-hexagonal/internal/app/api"
)

func main() {
	db, err := db.New()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	api := api.NewApplication(db)

	fiberAdapter := rest.NewAdapter(*api)

	fiberAdapter.Start()

}
