package fiber

import (
	"chatgpt-forwarder/adapter/in"
	"chatgpt-forwarder/cmd/fiber/middleware"
	"chatgpt-forwarder/cmd/fiber/routes"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)

const (
	PORT    = ":3000"
	API     = "api"
	Version = "v1"
)

func setRouteGroupApiV1(app *fiber.App) fiber.Router {
	prefix := fmt.Sprintf("/%s/%s", API, Version)
	return app.Group(prefix, middleware.ApiKeyAuth())
}

func StartSrv(completionController *in.ChatGPTCompletionController) {
	app := fiber.New()
	api := setRouteGroupApiV1(app)
	routes.SetApiV1Routes(api, completionController)
	err := app.Listen(PORT)
	log.Fatalf("server started failed with error:%s", err.Error())
}
