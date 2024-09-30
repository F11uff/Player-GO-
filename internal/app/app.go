package app

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"player/internal/config"
	"player/internal/handler"
	"sync"
)

func OpenApp() {
	app := fiber.New()
	wg := sync.WaitGroup{}
	cfg := config.DefaultConfig()

	fmt.Println(cfg)
	fmt.Println(cfg.HTTPServerConfig.Address)

	app.Static("/", "../../web/build")

	wg.Add(1)
	go func() {
		defer wg.Done()

		app.Get("/register", handler.RegisterHandler)

		if err := app.Listen(cfg.HTTPServerConfig.Address); err != nil {
			fmt.Printf("Ошибка запуска сервера: %v\n", err)
		}
	}()
	wg.Wait()
}
