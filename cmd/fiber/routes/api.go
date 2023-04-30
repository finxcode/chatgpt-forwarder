package routes

import (
	"chatgpt-forwarder/adapter/in"
	"github.com/gofiber/fiber/v2"
)

func SetApiV1Routes(router fiber.Router, completionAdapter *in.ChatGPTCompletionController) {
	router.Get("/getCompletion", completionAdapter.GetCompletion())
}
