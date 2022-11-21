package server

import (
	"github.com/Random7-JF/go-rcon/app/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html"
)

func Serve(App *config.App) {
	// Create a new engine
	htmlEngine := html.New("./views", ".html")

	htmlEngine.Debug(!App.Production)
	htmlEngine.Reload(!App.Production)

	App.WebServer = fiber.New(fiber.Config{
		Views: htmlEngine,
	})

	App.WebServer.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	SetupRoutes(App)

	url := App.WebSettings.Ip + ":" + App.WebSettings.Port
	App.WebServer.Listen(url)
}
