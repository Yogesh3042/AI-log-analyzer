package ai

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type OllamaRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
}

type OllamaResponse struct {
	Response string `json:"response"`
	Done     bool   `json:"done"`
}

func SummarizeLogsOllama(logs []string) (string, error) {
	prompt := "Summarize these logs and highlight key issues:\n\n"
	for _, l := range logs {
		prompt += l + "\n"
	}

	reqBody := OllamaRequest{
		Model:  "llama3.2",
		Prompt: prompt,
	}

	bodyBytes, _ := json.Marshal(reqBody)

	resp, err := http.Post("http://localhost:11434/api/generate", "application/json", bytes.NewBuffer(bodyBytes))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", errors.New("ollama error: " + resp.Status)
	}

	var result OllamaResponse
	responseText := ""

	for {
		decoder := json.NewDecoder(resp.Body)
		err := decoder.Decode(&result)
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}

		responseText += result.Response
		if result.Done {
			break
		}
	}

	return responseText, nil
}
