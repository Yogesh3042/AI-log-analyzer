async function analyzeLogs() {
    const logs = document.getElementById("logInput").value;
    const API_URL = "https://ubiquitous-fortnight-7wvqxjqxx763xw5v-8080.app.github.dev/analyze";

    try {
        const response = await fetch(API_URL, {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ logs })
        });

        const text = await response.text();
        console.log("Raw Response:", text); // <-- important debugging

        const data = JSON.parse(text);

        document.getElementById("result").innerText = data.summary;

    } catch (err) {
        document.getElementById("result").innerText =
            "Client Error: " + err.message;
    }
}
