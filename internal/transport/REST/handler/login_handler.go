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

	token, err := checkUser.AuthenticateUser(checkUser)

	fmt.Println(token)

	if err != nil {

		return ctx.JSON(fiber.Map{"error": fiber.StatusUnauthorized, "token": nil})
	}

	return ctx.JSON(fiber.Map{"message": "Password correct", "token": token})
}
