package command

import (
	"log"

	"github.com/Excute/botchi-go/internal/logger"
	"github.com/bwmarrin/discordgo"
)

var cmdFirst *discordgo.ApplicationCommand = &discordgo.ApplicationCommand{
	Name: "first",
	NameLocalizations: &map[discordgo.Locale]string{
		discordgo.Korean: "첫-커맨드",
	},

	Description: "Description of command",
	DescriptionLocalizations: &map[discordgo.Locale]string{
		discordgo.Korean: "커맨드 설명",
	},

	// TODO: Implement

	// Type: discordgo.ApplicationCommandType(discordgo.ApplicationCommandOptionString),

	// Options: []*discordgo.ApplicationCommandOption{
	// 	{
	// 		Name: "option-name",
	// 		NameLocalizations: map[discordgo.Locale]string{
	// 			discordgo.Korean: "옵션-이름",
	// 		},

	// 		Description: "Description of option",
	// 		DescriptionLocalizations: map[discordgo.Locale]string{
	// 			discordgo.Korean: "옵션 설명",
	// 		},
	// 	},
	// },
}

func cmdFirstHandler(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
	logger.Debug(session, "cmdFirstHandler() called", interaction.Interaction)

	if err := session.InteractionRespond(
		interaction.Interaction,
		&discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "This is response of first command",
			},
		},
	); err != nil {
		log.Panic(err)
	}
}

func setFirst() {
	commands = append(commands, cmdFirst)
	handlers[cmdFirst.Name] = cmdFirstHandler
}
