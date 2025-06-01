# GitHub Webhook to Telegram Bridge 🔗

A lightweight Go service that listens for GitHub webhooks and forwards notifications to Telegram chats with clean formatting.

[![Deploy with Vercel](https://vercel.com/button)](https://vercel.com/new/clone?repository-url=https%3A%2F%2Fgithub.com%2FAshokShau%2Fgithub-telegram-notify)

## 🌟 Features

- Real-time GitHub event notifications in Telegram
- Supports 20+ GitHub event types (pushes, PRs, issues, deployments, etc.)
- Clean, formatted messages with emoji visual cues
- Easy deployment to Vercel
- Lightweight

## 🚀 Quick Start

### Prerequisites
- Go 1.20+ (for local development)
- [Telegram bot token](https://core.telegram.org/bots#6-botfather)
- GitHub repository admin access

### Local Development
```bash
git clone https://github.com/AshokShau/github-telegram-notify.git
cd github-telegram-notify
go run main.go
```

For local testing, expose your port using:
```bash
ngrok http 3000
```

## ⚙️ Configuration

1. **Environment Variables**:
    - `BOT_TOKEN`: Your Telegram bot token
    - `PORT`: Server port (default: 3000)

2. **GitHub Webhook Setup**:
    - Payload URL: `https://your-domain.com/github?chat_id=YOUR_CHAT_ID`
    - Content type: `application/json`
    - Events: Select events to forward

## 🛠️ Supported Events

| Event Type          | Description                    |
|---------------------|--------------------------------|
| Push                | Code pushes to branches        |
| Pull Request        | PR opened/closed/merged        |
| Issues              | Issue created/commented/closed |
| Releases            | New version releases           |
| Deployments         | Code deployments               |
| Security Advisories | Vulnerability alerts           |
| And More            | ......                         |

## 🌐 Deployment

### Vercel (Recommended)
1. Fork this repository
2. Create new Vercel project
3. Import your forked repo
4. Add `BOT_TOKEN` in project settings
5. Deploy!

### Manual Deployment
Build and run the binary:
```bash
go build -o gh-telegram
./gh-telegram
```

## 📚 Documentation

- **Endpoint**: `/github` - Handles GitHub webhooks
- **Query Params**:
    - `chat_id`: Required Telegram chat ID

## 🤝 Contributing

Contributions welcome! Please:
1. Fork the repository
2. Create a feature branch
3. Submit a PR with clear description

## 📜 License

MIT License - See [LICENSE](LICENSE) for details.

## 💬 Support

- [Demo Bot](https://t.me/FallenAlertBot)
- [Telegram Support](https://t.me/AshokShau)
- [Update Channel](https://t.me/FallenProjects)
