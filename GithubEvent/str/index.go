package str

import (
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, _ *http.Request) {
	tmpl, err := template.New("index").Parse(indexHtml)
	if err != nil {
		http.Error(w, "Error loading page", http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}

const indexHtml = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>GitHub Webhook Listener</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            background-color: #2c2f33;
            color: #ddd;
            margin: 0;
            padding: 0;
        }

        h1 {
            text-align: center;
            color: #57a6ff;
            margin-top: 50px;
        }

        p {
            text-align: center;
            font-size: 16px;
            color: #bbb;
            margin-bottom: 20px;
        }

        .container {
            max-width: 800px;
            margin: 0 auto;
            padding: 30px;
            background-color: #40444b;
            border-radius: 10px;
            box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
        }

        .repo-link {
            text-align: center;
            margin-top: 20px;
        }

        .repo-link a {
            color: #57a6ff;
            text-decoration: none;
            font-weight: bold;
        }

        .repo-link a:hover {
            text-decoration: underline;
        }

        footer {
            text-align: center;
            font-size: 12px;
            color: #888;
            margin-top: 40px;
        }

        footer a {
            color: #57a6ff;
            text-decoration: none;
        }

        footer a:hover {
            text-decoration: underline;
        }

        .endpoint-info {
            margin-top: 30px;
            text-align: center;
            font-size: 14px;
            color: #ddd;
        }

        .endpoint-info a {
            color: #57a6ff;
            text-decoration: none;
            font-weight: bold;
        }

        .endpoint-info a:hover {
            text-decoration: underline;
        }
    </style>
</head>
<body>

<div class="container">
    <h1>GitHub Webhook Listener</h1>
    <p>This is a simple GitHub webhook listener written in Go.</p>
    <p>It listens for incoming webhook requests from GitHub and handles them accordingly.</p>

    <!-- Endpoint Information -->
    <div class="endpoint-info">
        <p>To use the webhook listener, GitHub should send a POST request to the following endpoint:</p>
        <p><strong>/github</strong></p>
        <p>You can pass a <strong>chat_id</strong> as a query parameter to receive notifications for a specific chat. Example: <a href="/github?chat_id=123456789" target="_blank">/github?chat_id=123456789</a></p>
    </div>

    <!-- GitHub Repository Link -->
    <div class="repo-link">
        <p>Check out the source code on <a href="https://github.com/AshokShau/github-telegram-notify" target="_blank">GitHub</a>.</p>
    </div>
</div>

<footer>
    <p>&copy; 2024 <a href="https://github.com/AshokShau/github-telegram-notify" target="_blank">GitHub Webhook Listener (AshokShau)</a> - All Rights Reserved</p>
</footer>

</body>
</html>
`
