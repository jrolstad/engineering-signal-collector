package main

import (
	"github.com/jrolstad/engineering-signal-collector/internal/pkg/messaging"
	"github.com/jrolstad/engineering-signal-collector/internal/pkg/orchestration"
	"io/ioutil"
	"net/http"
)

func PollForLatestChangeLogs(messageHub messaging.MessageHub) error {
	endpoint := "https://changelogs.md/api/github/microsoft/vscode-go/"

	result, getError := getDatafromChangelog(endpoint)
	if getError != nil {
		return getError
	}

	message := MapToSignalMessage("changelog", endpoint, result)

	sendError := orchestration.SendSignal(messageHub, message)

	return sendError
}

func getDatafromChangelog(endpoint string) (string, error) {
	data, getError := executeGet(endpoint)

	return data, getError
}

func executeGet(endpoint string) (string, error) {
	client := &http.Client{}

	request, requestError := http.NewRequest("GET", endpoint, nil)
	if requestError != nil {
		return "", requestError
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/json")

	response, responseError := client.Do(request)
	if responseError != nil {
		return "", responseError
	}

	defer response.Body.Close()
	body, readError := ioutil.ReadAll(response.Body)
	if readError != nil {
		return "", readError
	}

	bodyContent := string(body)

	return bodyContent, nil
}
