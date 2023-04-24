# Hamrobazar Monitor

Monitor new items on Hamrobazar and get alerts on Telegram.

## Prerequisite

1. Create a Telegram Bot and obtain the API token. Telegram has a unique way to create new bots. You'll need to use their official bot `@botfather` to create a new bot for yourself. Simply message [botfather](https://t.me/botfather) and it'll guide you through the process. This bot is needed to send you alerts about new items on Hamrobazar.

2. Now, find out the chat ID of the telegram user/channel/group where you want to receive the alerts. If you want the bot to send alerts to your personal Telegram account, you can send a message to [@RawDataBot](https://t.me/RawDataBot) and it'll give you the chat ID. For channels and group, take a look at the [Stackoverflow answer](https://stackoverflow.com/a/56546442)

## Installation

### 1. Build from source

This requires that you have Go v1.20+ installed locally.

```sh
git clone https://github.com/adityathebe/hamrobazar.git
cd hamrobazar
make install
```

### 2. Download prebuilt binary

Download the suitable binary from https://github.com/adityathebe/hamrobazar/releases/latest

## Usage

```sh
hamrobazar --chat-id <chat-id> --token <bot-api-token> fixtures/flat.yaml

# If env vars are set
export HAMROBAZAR_TG_TOKEN="xxxxxxxxxxx"
export HAMROBAZAR_TG_CHAT_ID="xxxxxxx:xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"

hamrobazar fixtures/flat.yaml
```

## Screenshots

![Telegram Chat Screenshot](screenshot.png)
