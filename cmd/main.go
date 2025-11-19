package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/you/logAnalyser/internal/analyzer"
)

type RequestBody struct {
	Logs string `json:"logs"`
}

func main() {
	fmt.Println("🟢 AI Log Analyzer UI running on http://localhost:8080")

	// Serve static files
	fs := http.FileServer(http.Dir("./web"))
	http.Handle("/", fs)

	// API endpoint
	http.HandleFunc("/analyze", func(w http.ResponseWriter, r *http.Request) {
		// CORS
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")

		// Preflight
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		if r.Method != http.MethodPost {
			json.NewEncoder(w).Encode(map[string]string{
				"summary": "Only POST supported",
			})
			return
		}

		// Parse input
		var req struct {
			Logs string `json:"logs"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			json.NewEncoder(w).Encode(map[string]string{
				"summary": "Invalid JSON: " + err.Error(),
			})
			return
		}

		// Run AI
		summary, err := analyzer.Analyze(req.Logs)
		if err != nil {
			json.NewEncoder(w).Encode(map[string]string{
				"summary": "AI Error: " + err.Error(),
			})
			return
		}

		// **ALWAYS RETURN JSON**
		json.NewEncoder(w).Encode(map[string]string{
			"summary": summary,
		})
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
