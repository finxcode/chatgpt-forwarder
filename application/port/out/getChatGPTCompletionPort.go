package out

import "chatgpt-forwarder/application/port/in"

type GetChatGPTCompletionPort interface {
	GetChatGPTCompletionOutgoing(command *in.CompletionCommand) (*in.CompletionResponse, error)
}
