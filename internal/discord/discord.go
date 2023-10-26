package discord

import "github.com/bwmarrin/discordgo"

var (
	Session *discordgo.Session
)

func init() {
	session, err := discordgo.New("Bot TOKEN")
	if err != nil {
		panic(err)
	}

	Session = session
}
