package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"

	"github.com/Excute/botchi-go/internal/command"
	"github.com/Excute/botchi-go/internal/handler"
)

func setIntents(session *discordgo.Session) {
	// TODO: Add intents needed
	session.Identify.Intents |= discordgo.IntentGuildMessages
}

func main() {
	// Create session
	session, err := discordgo.New("Bot " + os.Getenv("DISCORD_BOT_TOKEN"))
	if err != nil {
		log.Fatalf("Cannot create discord session: %+v", err)
	}

	// TODO: use logger instead
	fmt.Println("Bot is now online. Press CTRL-C to exit.")

	defer func() {
		// Cleanly close down the Discord session.
		session.Close()
	}()

	// Set intents needed
	setIntents(session)

	// Set handlers to session
	command.Add(session)
	handler.Add(session)

	// Open a websocket connection to Discord and begin listening.
	if err := session.Open(); err != nil {
		log.Fatalf("Cannot connect to Discord websocket: %+v", err)
	}

	// Register commands, deregister before session closing
	command.Register(session)
	defer command.Deregister(session)

	stopChannel := make(chan os.Signal, 1)
	signal.Notify(stopChannel, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stopChannel

	// Got stop signal, closing process
	// TODO: log
}
