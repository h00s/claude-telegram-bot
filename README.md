# Telegram bot with Claude AI

This is a Telegram bot that uses Claude AI to generate responses to messages.

## How to use

1. Create a Telegram bot using [@BotFather](https://t.me/BotFather)
2. Apply for a Anthropic API key.
3. Run the bot (preferably as a Docker container) with the following environment variables:
- `ANTHROPIC_KEY`: Your Anthropic API key
- `TELEGRAM_TOKEN`: Your Telegram bot token

## How to run as a Docker container

Here's an Docker compose example:

```yaml
version: '3.5'
services:
  claudegptbot:
    image: h00s/claudegpt-telegram-bot
    container_name: claudegptbot
    stop_grace_period: 15s
    restart: unless-stopped
    environment:
      - TELEGRAM_TOKEN=123456:abc
      - ANTHROPIC_KEY=12345
```

## Commands and usage

- `/start`: Start the bot
- `/hello`: Test the bot if it's working
- `/new`: Start a new conversation (Claude will forget the previous conversation)

Any message sent to the bot will be sent to Claude and the response will be sent back to the user.