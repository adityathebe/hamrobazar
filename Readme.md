# Hamrobazar Monitor

Monitor new items on Hamrobazar and get alerted on Telegram.

## Prerequisite

1. Create a Telegram Bot and obtain the API token.

2. Find out the chat ID of the telegram user/channel/group where you want to receive the alerts. To get the chat ID for your personal account, message [@RawDataBot](https://t.me/RawDataBot)

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
export HAMROBAZAR_TG_TOKEN="412345678"
export HAMROBAZAR_TG_CHAT_ID="xxxxxxx:xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"

hamrobazar fixtures/flat.yaml
```

## Screenshots

![Telegram Chat Screenshot](screenshot.png)
