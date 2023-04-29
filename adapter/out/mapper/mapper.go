package mapper

import (
	"chatgpt-forwarder/application/port/in"
	"chatgpt-forwarder/application/port/out"
)

func MapUserCommandToChatGPTCommand(command *in.CompletionCommand) *out.CompletionCommand {
	var messages []out.Message
	m := out.Message{}
	for _, message := range command.Messages {
		m.Role = message.Role
		m.Content = message.Content
		messages = append(messages, m)
	}
	return &out.CompletionCommand{
		Messages: messages,
	}
}

func MapChatGPTResponseToUserResponse(response *out.CompletionResponse) *in.CompletionResponse {
	if len(response.Data.Choices) == 0 {
		return nil
	}
	m := response.Data.Choices[0].Message
	respM := in.Message{
		Role:    m.Role,
		Content: m.Content,
	}
	return &in.CompletionResponse{
		Message: respM,
	}
}
