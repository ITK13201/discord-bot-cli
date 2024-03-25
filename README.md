# Discord Bot Cli

This is a CLI application for sending notifications to discord using a bot.

```
cli application for notifications to discord

Usage:
  discord-bot-cli [flags]

Flags:
  -c, --channel string       channel name defined in config file
      --config string        config file (default is $HOME/.discord-bot-cli.yaml)
  -d, --description string   description
  -h, --help                 help for discord-bot-cli
  -l, --level string         level (allowed: 'info', 'warn', 'error') (default "info")
  -t, --title string         title
```

## Configuration

The following is an example of a configuration file.

```yaml
token: "<your Discord Bot Token>"
channels:
  - name: "<your custom Discord Channel Name>"
    id: "<your Discord Channel ID>"
  - name: "<your custom Discord Channel Name>"
    id: "<your Discord Channel ID>"
```
