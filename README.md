# Discord Quote Bot

A simple Discord bot written in Go used by me and my friends, that sends a random quote daily and responds to user commands with categorized quotes. Uses `discordgo` for Discord API integration and `godotenv` for environment variable management.

## Features

- Sends a daily quote to a specified Discord channel.
- Supports commands for retrieving random, motivational, and funny quotes.
- Simple and lightweight, using only a single text file (`quotes.txt`) for quotes.
- Hosted on GitHub and deployed on Railway for easy deployment and maintenance.

## Commands

| Command             | Description              |
| ------------------- | ------------------------ |
| `!quote`            | Get a random quote       |
| `!quote random`     | Get a random quote       |
| `!quote motivation` | Get a motivational quote |
| `!quote funny`      | Get a funny quote        |
| `!quote help`       | Show available commands  |

## Installation

### Prerequisites

- Go 1.23 or later
- A Discord bot token (create one at [Discord Developer Portal](https://discord.com/developers/applications))

### Clone the Repository

```sh
git clone https://github.com/robinho46/discord-bot.git
cd discord-bot
```

### Set Up Environment Variables

Create a `.env` file in the project root and add your bot token and channel ID:

```
TOKEN=your-discord-bot-token
CHANNEL_ID=your-discord-channel-id
```

### Build and Run

```sh
go build -o discord-bot
./discord-bot
```

## Deployment

The bot is deployed on [Railway](https://railway.app/). To deploy it yourself:

1. Push your code to GitHub.
2. Connect your repository to Railway.
3. Set up environment variables (`TOKEN` and `CHANNEL_ID`).
4. Deploy and start the bot.

## Future Improvements

- Add more quote categories.
- Improve error handling and logging.
- Support database storage for quotes.

## License

This project is licensed under the MIT License.

