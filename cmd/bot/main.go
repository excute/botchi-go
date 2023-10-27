package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Excute/botchi-go/internal/discord"
	"github.com/Excute/botchi-go/internal/handler"
	"github.com/bwmarrin/discordgo"
)

func main() {
	session := discord.Session

	// Register the messageCreate func as a callback for MessageCreate events.
	session.AddHandler(handler.MessageCreate)

	// Only care about receiving message events.
	session.Identify.Intents = discordgo.IntentGuildMessages

	// Open a websocket connection to Discord and begin listening.
	if err := session.Open(); err != nil {
		panic(err)
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	session.Close()
}
