package api

import (
	log "log"
	"os"

	fiber "github.com/gofiber/fiber/v2"
	basicauth "github.com/gofiber/fiber/v2/middleware/basicauth"
	cors "github.com/gofiber/fiber/v2/middleware/cors"
)

func Server() {
	app := fiber.New(fiber.Config{
		// Поддерживаем многопоточность
		Prefork: true,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	app.Get("/", SendPageIndex)

	app.Get("/meme", GetRandomMeme)
	app.Get("/meme/:id", GetMeme)
	app.Post("/meme/:id", SetMemeLike)

	app.Route("/admin", func(router fiber.Router) {
		// Самый простой способ без пользовательского интерфейса
		router.Use(basicauth.New(basicauth.Config{
			Users: map[string]string{
				"admin": "admin",
			},
		}))

		router.Get("/", SendPageDashboard)
		router.Get("/sse", GetDashboard)

		router.Post("/meme/:id/:prio", SetMemePrio)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Fatal(app.Listen(":" + port))
}
