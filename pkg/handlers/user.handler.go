package handlers

import "github.com/gofiber/fiber/v2"

func UserHandlerRead(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"Username": "John",
	})

}

func HandlerHomepage(ctx *fiber.Ctx) error {
	return ctx.SendString("This is the home page")
}
