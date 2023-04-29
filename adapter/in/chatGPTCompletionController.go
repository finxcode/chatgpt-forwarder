package in

import (
	"chatgpt-forwarder/adapter/in/utils"
	"chatgpt-forwarder/application/port/in"
	"chatgpt-forwarder/application/port/in/common"
	"errors"
	"github.com/gofiber/fiber/v2"
	"log"
)

type ChatGPTCompletionController struct {
	getChatGPTCompletionUseCase in.GetChatGPTCompletionUseCase
}

func NewChatGPTCompletionController(getChatGPTCompletionUseCase in.GetChatGPTCompletionUseCase) *ChatGPTCompletionController {
	return &ChatGPTCompletionController{
		getChatGPTCompletionUseCase: getChatGPTCompletionUseCase,
	}
}

func (ctl *ChatGPTCompletionController) GetCompletion() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var mr *utils.MalformedRequest
		command := new(in.CompletionCommand)
		if err := utils.DecodeJSONBody(c, command); err != nil {
			if errors.As(err, &mr) {
				resp := common.Response{
					ErrCode: mr.Status,
					Message: mr.Msg,
					Data:    nil,
				}
				return c.JSON(resp)
			} else {
				resp := common.Response{
					ErrCode: fiber.StatusServiceUnavailable,
					Message: "service unavailable, please try later",
					Data:    nil,
				}
				return c.JSON(resp)
			}
		}
		respBody, err := ctl.getChatGPTCompletionUseCase.GetChatGPTCompletion(command)
		if err != nil || respBody == nil {
			log.Println(err.Error())
			resp := common.Response{
				ErrCode: fiber.StatusServiceUnavailable,
				Message: "service unavailable, please try later",
				Data:    nil,
			}
			return c.JSON(resp)
		} else {
			resp := common.Response{
				ErrCode: 0,
				Message: "ok",
				Data:    respBody,
			}
			return c.JSON(resp)
		}
	}
}
