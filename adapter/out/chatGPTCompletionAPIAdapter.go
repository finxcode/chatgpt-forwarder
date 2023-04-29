package out

import (
	"bytes"
	"chatgpt-forwarder/adapter/out/mapper"
	"chatgpt-forwarder/application/port/in"
	"chatgpt-forwarder/application/port/out"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

type Request struct {
	Model       string    `json:"model"`
	Messages    []Message `json:"messages"`
	Temperature float32   `json:"temperature"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
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

	if err != nil {
		return nil, err
	}
	r.Header.Add("Content-Type", "application/json")

	var client = &http.Client{
		Transport: &http.Transport{},
	}

	res, err := client.Do(r)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	resp, err := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(resp, &respBody)

	if err != nil {
		return nil, err
	}

	respMessage := mapper.MapChatGPTResponseToUserResponse(&respBody)

	return respMessage, nil

}
