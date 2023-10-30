package command

import "github.com/bwmarrin/discordgo"

var commands []*discordgo.ApplicationCommand
var handlers map[string]func(session *discordgo.Session, interaction *discordgo.InteractionCreate)

func setCommands() {
	// Call all command setters, maybe in order
	setFirst()
}

func init() {
	// Init handlers map
	handlers = make(map[string]func(session *discordgo.Session, interaction *discordgo.InteractionCreate))

	setCommands()
}

// Commands returns commands
func Commands() []*discordgo.ApplicationCommand {
	return commands
}

// Handler returns unified handler for commands
func Handler(session *discordgo.Session, interactionCreate *discordgo.InteractionCreate) {
	if handler, ok := handlers[interactionCreate.ApplicationCommandData().Name]; ok {
		handler(session, interactionCreate)
	}
}
