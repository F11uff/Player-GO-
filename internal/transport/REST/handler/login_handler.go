package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"player/internal/storage/postgresql/user"
)

func PostLoginHandler(ctx *fiber.Ctx) error {
	var checkUser user.UserLogin

	if err := ctx.BodyParser(&checkUser); err != nil {

		return ctx.JSON(fiber.Map{"error": fiber.StatusBadRequest})
	}

	token, error := checkUser.AuthenticateUser(checkUser)

	fmt.Println(token)

	if error != nil {
		fmt.Println(error)

		return ctx.JSON(fiber.Map{"error": fiber.StatusUnauthorized})
	}

	return ctx.JSON(fiber.Map{"message": "Password correct", "token": token})
}
