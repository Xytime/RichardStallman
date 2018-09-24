# RichardStallmanBot
Discord bot that corrects people when they accidentally say linux instead of GNU/Linux (or as I've taken to calling it, GNU+Linux)
Bot link: https://discordapp.com/oauth2/authorize?client_id=493584300336480259&scope=bot&permissions=3072

[img]https://github.com/Xytime/RichardStallman/raw/master/demo.png[/img]

## How 2 Build

Install Dependancies
```
go get github.com/bwmarrin/discordgo
```
and build! 

```go build richard.go```

## How 2 Run

Make sure richard is executable (if on linux)

```chmod +x richard```

Run discord bot

```DISCORD_BOT_TOKEN=<YourTokenHere> richard```

