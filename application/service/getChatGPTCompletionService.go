package service

import (
	"chatgpt-forwarder/application/port/in"
	"chatgpt-forwarder/application/port/out"
)

type getChatGPTCompletionService struct {
	getChatGPTCompletionPort out.GetChatGPTCompletionPort
}

func NewGetChatGPTCompletionService(getChatGPTCompletionPort out.GetChatGPTCompletionPort) *getChatGPTCompletionService {
	return &getChatGPTCompletionService{
		getChatGPTCompletionPort: getChatGPTCompletionPort,
	}
}

func (g *getChatGPTCompletionService) GetChatGPTCompletion(command *in.CompletionCommand) (*in.CompletionResponse, error) {
	return g.getChatGPTCompletionPort.GetChatGPTCompletionOutgoing(command)
}
