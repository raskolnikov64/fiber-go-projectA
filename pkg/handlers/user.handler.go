package handlers

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-project-1/database"
	"go-fiber-project-1/pkg/model/entity"
	"log"
)

func UserHandlerRead(ctx *fiber.Ctx) error {
	var users []entity.UserEntity

	result := database.DB.Find(&users)

	if result.Error != nil {
		log.Println(result.Error)
	}
	return ctx.JSON(users)
}

func HandlerHomepage(ctx *fiber.Ctx) error {
	return ctx.SendString("This is the home page")
}
