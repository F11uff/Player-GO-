//package app
//
//import (
//	"database/sql"
//	"fmt"
//	"github.com/gofiber/fiber/v2"
//	"net/http"
//	"player/internal/config"
//	"player/internal/handler"
//	"player/internal/storage/postgresql/user"
//	"sync"
//)
//
//func OpenApp() {
//	app := fiber.New()
//	wg := sync.WaitGroup{}
//	cfg := config.DefaultConfig()
//
//	fmt.Println(cfg)
//	fmt.Println(cfg.HTTPServerConfig.Address)
//
//	app.Static("/", "../../web/build")
//
//	wg.Add(1)
//	go func() {
//		defer wg.Done()
//
//		app.Get("/register", handler.RegisterHandler)
//
//		if err := app.Listen(cfg.HTTPServerConfig.Address); err != nil {
//			fmt.Printf("Ошибка запуска сервера: %v\n", err)
//		}
//	}()
//	wg.Wait()
//}

package app

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"player/internal/config"
	"player/internal/handler"
)

func OpenApp() {
	app := fiber.New()
	cfg := config.DefaultConfig()

	app.Static("/", "../../web/build")

	//// СДЕЛАТЬ СЛЕДУЮЩИЕ СТРОЧКИ В ГОРУТИНЕ !!!!!!!!!!!!!!!!!!

	app.Get("/", handler.RegisterHandler)

	app.Post("/", handler.PostLoginHandler)
	app.Post("/registration", handler.PostRegisterHandler)

	if err := app.Listen(cfg.HTTPServerConfig.Address); err != nil {
		fmt.Printf("Ошибка запуска сервера: %v\n", err)
	}

}
