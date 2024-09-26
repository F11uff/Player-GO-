package handler

import "github.com/gofiber/fiber/v2"

func RegisterHandler(ctx *fiber.Ctx) error {
	return ctx.SendFile("../../web/build/index.html")
}
