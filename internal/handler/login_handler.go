package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"player/internal/storage/postgresql/user"
)

func PostLoginHandler(ctx *fiber.Ctx) error {
	var checkUser user.UserLogin

	if err := ctx.BodyParser(&checkUser); err != nil {

		return ctx.JSON(fiber.Map{"status": fiber.StatusBadRequest})
	}

	fmt.Println(checkUser)

	fmt.Println("wfvwrfvcerwfe")
	return ctx.JSON(fiber.Map{})
}
