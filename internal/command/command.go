package command

import (
	"github.com/Excute/botchi-go/internal/logger"
	"github.com/bwmarrin/discordgo"
)

var commands []command

type command interface {
	add(session *discordgo.Session)
	register(session *discordgo.Session) error
	deregister(session *discordgo.Session) error
}

type genericCommand struct {
	handler            interface{}
	applicationCommand *discordgo.ApplicationCommand
}

func (g *genericCommand) add(session *discordgo.Session) {
	session.AddHandler(g.handler)
}

func (g *genericCommand) register(session *discordgo.Session) (err error) {
	logger.Debug(session, "registering command", nil) // TODO: set message

	g.applicationCommand, err = session.ApplicationCommandCreate(session.State.User.ID, "", g.applicationCommand)
	if err != nil {
		return
	}

	logger.Debug(session, "registered command", nil) // TODO: set message

	return nil
}

func (g *genericCommand) deregister(session *discordgo.Session) error {
	if g.applicationCommand.ID == "" {
		return nil
	}

	return session.ApplicationCommandDelete(session.State.User.ID, "", g.applicationCommand.ID)
}

func init() {
	commands = []command{
		// Append all commands
		first,
	}
}

func Add(session *discordgo.Session) {
	for _, command := range commands {
		command.add(session)
	}
}

func Register(session *discordgo.Session) {
	for _, command := range commands {
		if err := command.register(session); err != nil {
			logger.Error(session, err, nil) // TODO: write embed message
		}
	}
}

func Deregister(session *discordgo.Session) {
	for _, command := range commands {
		if err := command.deregister(session); err != nil {
			logger.Error(session, err, nil) // TODO: write embed message
		}
	}
}
