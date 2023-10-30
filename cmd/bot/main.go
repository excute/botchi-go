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
	"github.com/Excute/botchi-go/internal/logger"
)

func setIntents(session *discordgo.Session) {
	// TODO: Add intents needed
	session.Identify.Intents |= discordgo.IntentGuildMessages
}

func addHandlers(session *discordgo.Session) {
	session.AddHandler(handler.HandlePingPong)
	session.AddHandler(command.Handler)
}

func registerCommands(session *discordgo.Session) (registeredCommands []*discordgo.ApplicationCommand, err error) {
	commands := command.Commands()
	registeredCommands = make([]*discordgo.ApplicationCommand, len(commands))

	for i, command := range commands {
		registeredCommand, err := session.ApplicationCommandCreate(session.State.User.ID, "", command)
		if err != nil {
			return nil, err
		}
		registeredCommands[i] = registeredCommand
	}

	return registeredCommands, nil
}

func main() {
	// Create session
	session, err := discordgo.New("Bot " + os.Getenv("DISCORD_BOT_TOKEN"))
	if err != nil {
		log.Fatalf("Cannot create discord session: %+v", err)
	}

	// session.ChannelMessageSend("531633010433458178")

	// Set intents needed
	setIntents(session)

	// Set handlers to session
	addHandlers(session)

	// Open a websocket connection to Discord and begin listening.
	if err := session.Open(); err != nil {
		log.Fatalf("Cannot connect to Discord websocket: %+v", err)
	}

	if _, err := registerCommands(session); err != nil {
		logger.Panic(session, err, nil)
	}

	// Cleanly close down the Discord session.
	defer func() {
		// Wait here until CTRL-C or other term signal is received.
		fmt.Println("Bot is now online. Press CTRL-C to exit.")
		stopChannel := make(chan os.Signal, 1)
		signal.Notify(stopChannel, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
		<-stopChannel

		session.Close()
	}()
}
