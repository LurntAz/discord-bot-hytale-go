# Bot Discord pour Commandes API (discord-bot-hytale-go)

A Go (Golang) Discord bot to send API commands to a local server.

---

## 🌟 Features

- **`!update`**: Sends an `update` command to the local server.
- **`!restart`**: Sends a `restart` command to the local server.
- **`!whitelist <player>`**: Adds a player to the whitelist by sending a command to the local server.

---

## 📋 Prerequisites

- [Go](https://golang.org/dl/) (v1.16 or higher)
- A Discord account and a server where you have admin permissions.
- A Discord bot token (create one at the [Discord Developer Portal](https://discord.com/developers/applications))
- A local server listening at `http://localhost:8080/execute` (or another configurable URL)

---

## 🛠 Installation

1. Clone this repository:
   ```bash
   git clone https://github.com/your-username/discord-bot.git
   cd discord-bot
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

---

## ⚙ Configuration

1. Create a `.env` file in the project root with the following variables:
   ```env
   DISCORD_BOT_TOKEN=your_discord_bot_token_here
   SERVER_URL=http://localhost:8080/execute
   ```

2. Replace `your_discord_bot_token_here` with your actual Discord bot token.

---

## 🚀 Running the Bot

1. Load environment variables (Linux/macOS):
   ```bash
   export $(grep -v '^#' .env | xargs)
   ```

   For Windows (PowerShell):
   ```powershell
   Get-Content .env | ForEach-Object { if ($_ -match '^\s*([^=]+)=(.*)') { [Environment]::SetEnvironmentVariable($matches[1], $matches[2]) } }
   ```

2. Start the bot:
   ```bash
   go run main.go
   ```

---

## 🤖 Inviting the Bot to Your Discord Server

1. Replace `CLIENT_ID` in the following link with your Discord application ID:
   ```
   https://discord.com/oauth2/authorize?client_id=CLIENT_ID&scope=bot&permissions=67584
   ```

2. Open the link in your browser and select the server where you want to add the bot.

---

## 📜 Available Commands

| Command | Description |
|---------|-------------|
| `!update` | Sends an `update` command to the local server. |
| `!restart` | Sends a `restart` command to the local server. |
| `!whitelist <player>` | Adds a player to the whitelist. |

---

## 📂 Project Structure

```
/discord-bot/
  ├── main.go          # Main bot source code
  ├── go.mod           # Go module file
  ├── go.sum           # Dependency checksums
  └── .env             # Configuration file (do not commit)
```

---

## ⚠ Common Issues

- **Bot not responding**: Ensure the token is correct and the bot is invited with the right permissions.
- **Connection error**: Verify your local server is running and accessible at the configured URL.
- **Unknown command**: Check the command prefix (`!`) and spelling.
