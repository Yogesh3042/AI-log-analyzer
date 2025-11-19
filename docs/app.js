const API_URL = "https://ai-log-analyzer-wcpb.onrender.com/analyze";

async function analyzeLogs() {
    const logs = document.getElementById("logInput").value.trim();
    const resultBox = document.getElementById("result");

    if (!logs) {
        resultBox.innerHTML = "<p class='error'>Please paste logs first.</p>";
        return;
    }

    resultBox.innerHTML = "<p class='loading'>⏳ Analyzing logs with AI...</p>";

    try {
        const response = await fetch(API_URL, {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ logs })
        });

        const text = await response.text();

        try {
            const json = JSON.parse(text);
            resultBox.innerHTML = `<pre>${json.result}</pre>`;
        } catch {
            resultBox.innerHTML = `<pre class='error'>Server Error: ${text}</pre>`;
        }

    } catch (error) {
        resultBox.innerHTML = `<p class='error'>Client Error: ${error}</p>`;
    }
}
