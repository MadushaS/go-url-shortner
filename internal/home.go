// internal/homepage.go
package internal

import (
	"html/template"
	"net/http"
)

var tpl = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>URL Shortener</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 50px; }
        form { display: flex; flex-direction: column; width: 300px; }
        input, button { padding: 10px; margin: 5px 0; }
        .short-url { margin-top: 20px; }
    </style>
</head>
<body>
    <h1>URL Shortener</h1>
    <form id="shortenForm">
        <input type="url" id="url" placeholder="Enter URL" required>
        <input type="text" id="custom_id" placeholder="Custom ID (optional)">
        <input type="datetime-local" id="expires" placeholder="Expiration Date (optional)">
        <button type="submit">Shorten</button>
    </form>
    <div class="short-url" id="shortURL"></div>

    <script>
        document.getElementById('shortenForm').addEventListener('submit', async function(event) {
            event.preventDefault();
            const url = document.getElementById('url').value;
            const customID = document.getElementById('custom_id').value;
            const expires = document.getElementById('expires').value;
            const response = await fetch('/shorten', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ url, custom_id: customID, expires: new Date(expires).toISOString() })
            });
            const result = await response.json();
            document.getElementById('shortURL').textContent = 'Short URL: ' + result.short_url;
        });
    </script>
</body>
</html>
`

func HomePage(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("homepage").Parse(tpl)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}
