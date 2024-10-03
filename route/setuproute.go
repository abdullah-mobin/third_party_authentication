package route

import (
	"Google_sign_option/handler"
	"Google_sign_option/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoute(app *fiber.App) {
	login := app.Group("/login")
	login.Post("/", handler.Login).Name("login-route")
	login.Get("/google", handler.LoginWithGoogle).Name("google-route")
	app.Get("/oauth/redirect", handler.GoogleCallback).Name("google-callback")

	hlw := app.Group("/hlw-world", middleware.Protected())
	hlw.Get("/welcome", handler.Welcome)
}
