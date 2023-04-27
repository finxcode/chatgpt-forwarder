package out

type getChatGPTCompletionPort interface {
	GetChatGPTCompletion(command *CompletionCommand) *CompletionResponse
}
