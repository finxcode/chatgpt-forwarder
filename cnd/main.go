package main

import (
	"chatgpt-forwarder/adapter/in"
	"chatgpt-forwarder/adapter/out"
	"chatgpt-forwarder/application/service"
	"chatgpt-forwarder/cnd/fiber"
)

func main() {

	chatGPTCompletionAPIAdapter := out.NewChatGPTCompletionAPIAdapter()
	getChatGPTCompletionService := service.NewGetChatGPTCompletionService(chatGPTCompletionAPIAdapter)
	chatGPTCompletionController := in.NewChatGPTCompletionController(getChatGPTCompletionService)
	fiber.StartSrv(chatGPTCompletionController)

}
