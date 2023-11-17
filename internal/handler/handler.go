package handler

import (
	"log"

	"github.com/Excute/botchi-go/internal/command"
	"github.com/Excute/botchi-go/internal/logger"
	"github.com/bwmarrin/discordgo"
)

func Handlers() []interface{} {
	return []interface{}{
		// List of handlers
		command.Handler,
	}
}

// return []func(){
// 	command.Handler,
// 	HandlePingPong,
// }
// return nil

// HandlePingPong responses to ping and pong
func HandlePingPong(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	logger.Debug(s, "got message", &discordgo.Interaction{
		Message: m.Message,
	})

	// If the message is "ping" reply with "Pong!"
	if m.Content == "ping" {
		if _, err := s.ChannelMessageSend(m.ChannelID, "Pong!"); err != nil {
			log.Panic(err)
		}
	}

	// If the message is "pong" reply with "Ping!"
	if m.Content == "pong" {
		if _, err := s.ChannelMessageSend(m.ChannelID, "Ping!"); err != nil {
			log.Panic(err)
		}
	}
}
