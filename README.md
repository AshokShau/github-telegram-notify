# GitHub Telegram Notify

A powerful, lightweight Go application that forwards GitHub webhook events to Telegram chats with rich, customizable notifications.

[![Deploy with Vercel](https://vercel.com/button)](https://vercel.com/new/clone?repository-url=https%3A%2F%2Fgithub.com%2FAshokShau%2Fgithub-telegram-notify)

This tool acts as a bridge between your GitHub repositories and your Telegram chats. It listens for over 40 different GitHub events—from code pushes and pull requests to security alerts and deployments—and transforms them into well-formatted, easy-to-read Telegram messages.

It's designed for developers, DevOps engineers, and project managers who want to stay updated on repository activity without leaving their favorite messaging app.

## 🌟 Features

- **Comprehensive Event Support**: Get notified for over 40 GitHub event types.
- **Rich Formatting**: Messages are parsed into Telegram's MarkdownV2 format, with emojis and interactive buttons.
- **Easy to Deploy**: One-click deployment to Vercel, or run it anywhere Go is supported.
- **Lightweight & Fast**: Built with Go for minimal resource consumption.
- **Customizable**: Fork the repo to customize message formats and add your own logic.
- **Open Source**: Licensed under the MIT license.

## 🛠️ Supported Events

This webhook supports a wide range of GitHub events, including:

- **Code & Repository**: `push`, `release`, `fork`, `repository`
- **Pull Requests**: `pull_request`
- **Issues**: `issues`, `issue_comment`
- **Discussions**: `discussion`, `discussion_comment`
- **CI/CD & Deployments**: `deployment`, `deployment_status`, `workflow_job`, `workflow_run`, `workflow_dispatch`
- **Security**: `repository_vulnerability_alert`, `secret_scanning_alert`
- **And many more...**

A full list of supported events can be found in the `src/utils/githubEvents.go` file.

## 🚀 Getting Started

### Prerequisites

- **Go**: Version 1.20 or higher (for local development).
- **Telegram Bot Token**: You can get one from [@BotFather](https://core.telegram.org/bots#6-botfather).
- **GitHub Repository**: You'll need admin access to the repository you want to monitor.

### Local Development

1.  **Clone the repository**:
    ```bash
    git clone https://github.com/AshokShau/github-telegram-notify.git
    cd github-telegram-notify
    ```

2.  **Set up environment variables**:
    Create a `.env` file in the root of the project and add your Telegram bot token:
    ```
    BOT_TOKEN="YOUR_TELEGRAM_BOT_TOKEN"
    ```

3.  **Run the application**:
    ```bash
    go run main.go
    ```
    The server will start on port `3000` by default.

4.  **Expose your local server**:
    To receive webhooks from GitHub, you'll need to expose your local server to the internet. We recommend using a tool like [ngrok](https://ngrok.com/):
    ```bash
    ngrok http 3000
    ```
    This will give you a public URL that you can use for your GitHub webhook.

### GitHub Webhook Configuration

1.  Go to your GitHub repository's **Settings > Webhooks**.
2.  Click **Add webhook**.
3.  **Payload URL**: Enter the URL from ngrok, followed by `/github?chat_id=YOUR_CHAT_ID`. Replace `YOUR_CHAT_ID` with the ID of the Telegram chat where you want to receive notifications.
4.  **Content type**: Select `application/json`.
5.  **Secret**: (Optional) You can add a webhook secret for added security.
6.  **Which events would you like to trigger this webhook?**: Select the events you want to be notified about.
7.  Click **Add webhook**.

## 🌐 Deployment

### Vercel (Recommended)

The easiest way to deploy this application is with Vercel.

1.  **Fork this repository**.
2.  Click the **Deploy with Vercel** button at the top of this README.
3.  In your Vercel project settings, add your `BOT_TOKEN` as an environment variable.
4.  Deploy!

### Manual Deployment

You can also build the application and run it as a binary on any server:

```bash
go build -o gh-telegram
./gh-telegram
```

## 🤝 Contributing

Contributions are welcome! If you'd like to help improve this project, please:

1.  Fork the repository.
2.  Create a new branch for your feature or bug fix.
3.  Submit a pull request with a clear description of your changes.

## 📜 License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## 💬 Support

- **Demo Bot**: [@FallenAlertBot](https://t.me/FallenAlertBot)
- **Telegram Support**: [@AshokShau](https://t.me/AshokShau)
- **Updates Channel**: [@FallenProjects](https://t.me/FallenProjects)
