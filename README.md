# wargame_bot_v2

A complete rewrite of my original discord bot.

## Setup

Before starting the application you need to create a `conf.yaml` file in the root directory. Fill in the the following template.

```yaml
# conf.yaml

# Your discord bot token and your own discord id.
discord:
    bot_token: ""
    owner_id: ""

# Your rcon port and password used by the server.
rcon:
    ip: ""
    port: ""
    pword: ""
```

## Discord

### Commands

All commands are stored in the [`commands.json`](./discord/commands.json) file. You can find more information on how they work on the [discord developer documentaion site](https://discord.com/developers/docs/interactions/application-commands).

## TODO
- Move settings to a single file
- Enable live chat reading
- Have server states.

## Repo design pattern

```
main.go
domain/
    models.go           # data types
store/
    store.go
        sqlite/
            sqlite.go
        mock/
            mock.go
service/
    service.go

```
