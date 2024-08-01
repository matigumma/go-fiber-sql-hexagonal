package rest

import (
	"log"

	"github.com/matigumma/go-fiber-sql-hexagonal/internal/app/api"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/static"
)

type Adapter struct {
	api api.Application
}

func NewAdapter(api api.Application) *Adapter {
	return &Adapter{api: api}
}

func (s *Adapter) Start() {
	app := fiber.New()

	app.Use(cors.New())

	index := app.Group("/")
	index.Use("/", static.New("./public"))

	// Start the server on port 3000
	log.Fatal(app.Listen(":3000"))
}
