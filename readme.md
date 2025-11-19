Great — here is a **professional, polished README.md** for your project.

You can copy-paste this entire content into your `readme.md` file in the root of your repo.

---

# 📘 **AI Log Analyzer**

AI Log Analyzer is a lightweight, fast, developer-friendly tool that analyzes system/application logs using **local AI models (via Ollama)**.
It reads log files, detects issues, summarizes patterns, and provides actionable insights — all without sending any data outside your machine.

This project includes:

✅ Go backend for log ingestion
✅ Ollama AI integration (local LLM—no API key needed)
✅ Simple HTML/JS UI dashboard
✅ Sample log files
✅ Clean modular project structure

---

## 🚀 Features

### 🔍 **AI-Powered Log Analysis**

* Detects anomalies, failures, spikes, latency issues.
* Summarizes logs into human-friendly insights.
* Generates recommendations for fixes.

### 🖥️ **Web Dashboard**

* Upload logs
* Run “AI Analyze” with one click
* Get structured summary instantly

### ⚡ **Fast & Local**

* Runs entirely with **Ollama** (no external API, no billing).
* Works in Codespaces, Windows (no admin needed via Codespaces), Linux, macOS.

---

## 📂 **Project Structure**

```
AI-log-analyzer/
│
├── cmd/
│   └── main.go                 # Entry point for HTTP server
│
├── internal/
│   └── analyzer/
│       ├── analyzer.go         # Log reading + AI call
│       └── ollama_client.go    # Ollama API wrapper
│
├── pkg/
│   └── utils/
│       └── file_reader.go      # File reading helpers
│
├── web/
│   ├── index.html              # UI Dashboard
│   ├── app.js                  # Frontend logic
│   └── styles.css              # UI Styling
│
├── sample_logs/
│   └── sample.log              # Test logs
│
├── go.mod
├── go.sum
└── README.md
```

---

## 🛠️ **Tech Stack**

### **Backend**

* Go 1.22+
* Ollama (local AI)
* net/http

### **Frontend**

* HTML5
* Vanilla JS
* Fetch API
* Minimal CSS

---

## ▶️ **Running Locally**

### **1️⃣ Install Ollama**

(Already pre-installed in GitHub Codespaces)

Otherwise:

```bash
curl -fsSL https://ollama.com/install.sh | sh
```

Run the server:

```bash
ollama serve
```

### **2️⃣ Pull a model**

Recommended:

```bash
ollama pull llama3.1
```

Check installed models:

```bash
ollama list
```

### **3️⃣ Run the Go server**

From project root:

```bash
go run cmd/main.go
```

You should see:

```
🟢 AI Log Analyzer Started
```

### **4️⃣ Open the UI**

If running locally:

```
http://localhost:8080
```

If using GitHub Codespaces → click "Open in Browser".

---

## 📤 **Usage**

### **Upload logs → Analyze → Get AI Summary**

The output includes:

* Total log lines
* Summary of events
* Key issues
* Recommendations

---

## 🧪 Sample Output

```
Total log lines: 10

AI Summary:
- High API latency detected
- Redis timeout issues
- Payment gateway failure (502)
- High disk usage alert
```

---


## 🤝 Contributing

Pull Requests are welcome!

---

## 📝 License

This project is licensed under **MIT License**.

---

## ⭐ Support

If you like this project, star the repo ⭐ — it helps grow visibility & motivates more improvements.

---

