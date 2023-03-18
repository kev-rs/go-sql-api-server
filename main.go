package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	usr "github.com/kev-tsx/go-sql-server/controllers"
	"github.com/kev-tsx/go-sql-server/db"
)

var DB *sql.DB

func main() {
	// Capture connection properties.
	db.Conect()

	app := fiber.New()
	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	app.Use(cors.New())
	app.Static("/", "./client/dist")

	app.Get("/users", usr.GetAll)

	app.Listen(":4000")
	fmt.Println("Server listen on port 4000")
}
