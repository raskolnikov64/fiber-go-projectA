package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go-fiber-project-1/database"
	"go-fiber-project-1/pkg/routes"
)

const Pnum string = ":8080"

func main() {
	database.DbInit()

	app := fiber.New()

	routes.RoutesInit(app)

	fmt.Println(fmt.Sprintf("App is running on port %s", Pnum))

	err := app.Listen(Pnum)
	if err != nil {
		return
	}
}
