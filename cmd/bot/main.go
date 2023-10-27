package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Excute/botchi-go/internal/handler"
	"github.com/bwmarrin/discordgo"
)

var (
	session *discordgo.Session
)

func init() {
	setSession()

	setIntents(
		discordgo.IntentGuildMessages,
	)
}

func setSession() {
	s, err := discordgo.New("Bot " + os.Getenv("DISCORD_BOT_TOKEN"))
	if err != nil {
		panic(err)
	}

	session = s
}

func setIntents(intents ...discordgo.Intent) {
	for _, intent := range intents {
		session.Identify.Intents |= intent
	}
}

func main() {
	// Register the messageCreate func as a callback for MessageCreate events.
	session.AddHandler(handler.MessageCreate)

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
