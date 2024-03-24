# Discord Bot Cli

This is a CLI application for sending notifications to discord using a bot.

```
Usage:
  discord-bot-cli [flags]

Flags:
  -c, --channel string       channel ID
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
```
