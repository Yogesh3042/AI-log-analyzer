package models

type LogEntry struct {
	Timestamp string
	Level     string
	Message   string
	Metadata  map[string]interface{}
}
