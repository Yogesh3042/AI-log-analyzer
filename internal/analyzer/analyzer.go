package analyzer

import (
	"context"
	"fmt"

	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/you/logAnalyser/internal/parser"
	"github.com/you/logAnalyser/pkg/ai"
)

func RunBasicAnalysis(path string) string {
	lines, err := parser.ParseFile(path)
	if err != nil {
		return "Error reading file: " + err.Error()
	}

	summary, err := ai.SummarizeLogsOllama(lines)
	if err != nil {
		return "AI Summary Error: " + err.Error()
	}

	return fmt.Sprintf("Total log lines: %d\n\nAI Summary:\n%s", len(lines), summary)
}
func Analyze(logs string) (string, error) {
	llm, err := ollama.New(ollama.WithModel("llama3"))
	if err != nil {
		return "", fmt.Errorf("ollama init error: %w", err)
	}

	ctx := context.Background()
	resp, err := llm.Call(ctx, "Summarize these logs and highlight issues:\n\n"+logs)
	if err != nil {
		return "", fmt.Errorf("ollama error: %w", err)
	}

	return resp, nil
}
