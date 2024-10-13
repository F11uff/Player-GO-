package app

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"player/internal/config"
	"player/internal/transport/REST/handler"
	"sync"
)

func OpenApp() {
	app := fiber.New()
	cfg := config.DefaultConfig()
	wq := sync.WaitGroup{}

	app.Static("/", "../../web/build")

	wq.Add(2)
	go func() {
		defer wq.Done()

		app.Get("/", handler.RegisterHandler)

		app.Post("/", handler.PostLoginHandler)
		app.Post("/registration", handler.PostRegisterHandler)
	}()

	go func() {
		defer wq.Done()

		if err := app.Listen(cfg.HTTPServerConfig.Address); err != nil {

			fmt.Printf("Ошибка запуска сервера: %v\n", err)
		}
	}()

	wq.Wait()

}
