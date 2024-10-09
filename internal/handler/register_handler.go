package handler

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"player/internal/storage/postgresql/user"
	"strings"
)

func PostRegisterHandler(ctx *fiber.Ctx) error {
	var newUser user.UserRegistration

	if err := ctx.BodyParser(&newUser); err != nil {

		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid JSON"})
	}

	if err := strings.Contains(newUser.Email, "@"); !err {

		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid email"})
	}

	if err := newUser.AddUser(newUser); err != nil {

		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to add user"})
	}

	return ctx.JSON(fiber.Map{"message": "UserRegistration registered successfully!"})
}
