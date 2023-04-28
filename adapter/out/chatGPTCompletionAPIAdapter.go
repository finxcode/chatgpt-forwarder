package out

import (
	"bytes"
	"chatgpt-forwarder/adapter/out/mapper"
	"chatgpt-forwarder/application/port/in"
	"chatgpt-forwarder/application/port/out"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const (
	HOST        = "http://34.228.23.28"
	PORT        = ":3000"
	API         = "/api/v1/getCompletion"
	MODEL       = "gpt-3.5-turbo"
	TEMPERATURE = 0.7
)

type chatGPTCompletionAPIAdapter struct {
}

func NewChatGPTCompletionAPIAdapter() *chatGPTCompletionAPIAdapter {
	return &chatGPTCompletionAPIAdapter{}
}

func (c *chatGPTCompletionAPIAdapter) GetChatGPTCompletionOutgoing(command *in.CompletionCommand) (*in.CompletionResponse, error) {
	chatGPTCommand := mapper.MapUserCommandToChatGPTCommand(command)
	chatGPTCommand.Model = MODEL
	chatGPTCommand.Temperature = TEMPERATURE

	respBody := out.CompletionResponse{}
	marshal, err := json.Marshal(chatGPTCommand)
	if err != nil {
		return nil, err
	}
	reqBody := bytes.NewReader(marshal)

	r, err := http.NewRequest("GET", fmt.Sprintf("%s%s%s", HOST, PORT, API), reqBody)
	log.Println(fmt.Sprintf("%s%s%s", HOST, PORT, API))
	if err != nil {
		return nil, err
	}
	r.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&respBody)
	log.Println(respBody.Model)
	if err != nil {
		return nil, err
	}

	respMessage := mapper.MapChatGPTResponseToUserResponse(&respBody)

	return respMessage, nil

}
