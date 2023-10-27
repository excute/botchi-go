package discord

import (
	"os"

	"github.com/bwmarrin/discordgo"
)

var (
	Session *discordgo.Session
)

func init() {
	session, err := discordgo.New("Bot " + os.Getenv("DISCORD_BOT_TOKEN"))
	if err != nil {
		panic(err)
	}

	Session = session
}
