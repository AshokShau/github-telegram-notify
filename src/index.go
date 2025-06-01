package src

import (
	"html/template"
	"net/http"
	"time"
)

// Home handles the root endpoint and renders the homepage
func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	data := struct {
		CurrentYear int
	}{
		CurrentYear: time.Now().Year(),
	}

	tmpl, err := template.New("index").Parse(indexHTML)
	if err != nil {
		http.Error(w, "Error loading page", http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}

const indexHTML = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta name="description" content="GitHub Webhook Listener for processing GitHub events and forwarding notifications">
    <title>GitHub Webhook Listener</title>
    <style>
        :root {
            --primary-color: #57a6ff;
            --bg-color: #2c2f33;
            --container-color: #40444b;
            --text-color: #ddd;
            --text-secondary: #bbb;
            --link-hover: #7fbfff;
            --shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
        }

        body {
            font-family: 'Segoe UI', system-ui, -apple-system, sans-serif;
            background-color: var(--bg-color);
            color: var(--text-color);
            margin: 0;
            padding: 0;
            line-height: 1.6;
        }

        h1 {
            text-align: center;
            color: var(--primary-color);
            margin-top: 2rem;
            margin-bottom: 1rem;
            font-weight: 600;
        }

        p {
            text-align: center;
            font-size: 1rem;
            color: var(--text-secondary);
            margin-bottom: 1.25rem;
        }

        .container {
            max-width: 800px;
            margin: 0 auto;
            padding: 2rem;
            background-color: var(--container-color);
            border-radius: 0.625rem;
            box-shadow: var(--shadow);
        }

        .repo-link {
            text-align: center;
            margin-top: 1.25rem;
        }

        a {
            color: var(--primary-color);
            text-decoration: none;
            font-weight: 500;
            transition: color 0.2s ease;
        }

        a:hover {
            color: var(--link-hover);
            text-decoration: underline;
        }

        footer {
            text-align: center;
            font-size: 0.75rem;
            color: var(--text-secondary);
            margin-top: 2.5rem;
            padding-bottom: 1rem;
        }

        .endpoint-info {
            margin-top: 1.875rem;
            padding: 1rem;
            background-color: rgba(0, 0, 0, 0.1);
            border-radius: 0.5rem;
            font-size: 0.875rem;
        }

        .endpoint-info code {
            display: inline-block;
            background-color: rgba(0, 0, 0, 0.2);
            padding: 0.2rem 0.4rem;
            border-radius: 0.25rem;
            font-family: 'Courier New', monospace;
            margin: 0.2rem 0;
        }

        .features {
            margin: 1.5rem 0;
            text-align: left;
        }

        .features ul {
            padding-left: 1.5rem;
        }

        .features li {
            margin-bottom: 0.5rem;
        }

        @media (max-width: 768px) {
            .container {
                padding: 1.5rem;
                margin: 0 1rem;
            }
            
            h1 {
                font-size: 1.5rem;
            }
        }
    </style>
</head>
<body>
<div class="container">
    <h1>GitHub Webhook Listener</h1>
    <p>A lightweight service for processing GitHub webhooks and forwarding notifications</p>

    <div class="features">
        <h3 style="color: var(--primary-color); text-align: center;">Key Features</h3>
        <ul>
            <li>Process GitHub webhook events in real-time</li>
            <li>Forward notifications to Telegram chats</li>
            <li>Simple configuration with chat_id parameter</li>
            <li>Lightweight and fast Go implementation</li>
            <li>Easy to deploy and integrate</li>
        </ul>
    </div>

    <div class="endpoint-info">
        <h3 style="color: var(--primary-color); text-align: center; margin-top: 0;">Usage</h3>
        <p>Configure your GitHub webhook to send POST requests to:</p>
        <p><code>/github</code></p>
        <p>Optional query parameters:</p>
        <ul style="list-style-type: none; padding-left: 0;">
            <li><code>chat_id</code> - Specify Telegram chat ID for notifications</li>
            <li><code>secret</code> - Add webhook secret for verification</li>
        </ul>
        <p>Example: <a href="/github?chat_id=123456789" target="_blank" rel="noopener noreferrer">/github?chat_id=123456789</a></p>
    </div>

    <div class="repo-link">
        <p>View source code and contribute on <a href="https://github.com/AshokShau/github-telegram-notify" target="_blank" rel="noopener noreferrer">GitHub</a>.</p>
    </div>
</div>

<footer>
    <p>&copy; {{.CurrentYear}} <a href="https://github.com/AshokShau/github-telegram-notify" target="_blank" rel="noopener noreferrer">GitHub Webhook Listener</a> | AshokShau</p>
</footer>
</body>
</html>`
