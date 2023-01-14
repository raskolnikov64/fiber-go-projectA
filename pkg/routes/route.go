package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-project-1/pkg/handlers"
)

func RoutesInit(o *fiber.App) {
	o.Get("/home", handlers.HandlerHomepage)

	o.Get("/", handlers.HandlerHomepage)

	o.Get("/user", handlers.UserHandlerRead)

}
