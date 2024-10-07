package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"player/internal/storage/postgresql/user"
)

func RegisterHandler(ctx *fiber.Ctx) error {

	return ctx.SendFile("../../web/build/index.html")
}

func PostRegisterHandler(ctx *fiber.Ctx) error {
	var newUser user.User

	if err := ctx.BodyParser(&newUser); err != nil {

		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid JSON"})
	}

	fmt.Println("logs : ", newUser)

	if err := newUser.AddUser(newUser); err != nil {

		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to add user"})
	}

	return ctx.JSON(fiber.Map{"message": "User registered successfully!"})
}
