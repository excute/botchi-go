package logger

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Excute/botchi-go/internal/util"
	"github.com/bwmarrin/discordgo"
)

var (
	isDebugString string
	isDebug       bool

	logGuildID   string
	logChannelID string

	// Color reference: https://m2.material.io/design/color/the-color-system.html#tools-for-picking-colors
	logTemplatesMap = map[string]discordgo.MessageEmbed{
		"panic": {
			Title: "Panic",
			Color: 0xF44336,
		},
		"error": {
			Title: "Error",
			Color: 0x9C27B0,
		},
		"warning": {
			Title: "Warning",
			Color: 0xFFEB3B,
		},
		"info": {
			Title: "Info",
			Color: 0x4CAF50,
		},
		"debug": {
			Title: "Debug",
			Color: 0x2196F3,
		},
	}
)

func init() {
	isDebugString = os.Getenv("CMD_BOT_DEBUG")
	if isDebugString == "TRUE" {
		isDebug = true
		fmt.Println("Debug logging is enabled")
	}

	logGuildID = os.Getenv("DISCORD_BOT_LOG_GUILD_ID")
	logChannelID = os.Getenv("DISCORD_BOT_LOG_CHANNEL_ID")
}

func timestamp() (timestamp string) {
	return time.Now().Format(time.RFC3339)
}

func createEmbed(session *discordgo.Session, logType string, interaction *discordgo.Interaction) *discordgo.MessageEmbed {
	embed := logTemplatesMap[logType]
	embed.Timestamp = timestamp()

	if interaction == nil {
		return &embed
	}

	var (
		guildID   = interaction.GuildID
		channelID = interaction.ChannelID
		messageID = interaction.ID
	)

	if guildID == "" {
		guildID = interaction.Message.GuildID
	}
	if channelID == "" {
		channelID = interaction.Message.ChannelID
	}
	if messageID == "" {
		messageID = interaction.Message.ID
	}

	if guild, err := session.Guild(guildID); err == nil {
		var ownerName string
		if owner, err := session.User(guild.OwnerID); err == nil {
			ownerName = owner.String()
		}
		embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
			Name:   "Guild",
			Value:  fmt.Sprintf("%s (Owner: %s (ID: `%s`))", guild.Name, ownerName, guild.OwnerID),
			Inline: false,
		})
	}

	if channel, err := session.Channel(channelID); err != nil {
		embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
			Name:   "Channel",
			Value:  fmt.Sprintf("%s(ID: `%s`)", channel.Name, channel.ID),
			Inline: false,
		})
	}

	if guildID != "" && channelID != "" && messageID != "" {
		embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
			Name:   "Link",
			Value:  fmt.Sprintf("[link](%s)", util.MessageURL(guildID, channelID, messageID)),
			Inline: false,
		})
	}

	if triggerTime, err := discordgo.SnowflakeTimestamp(messageID); err == nil {
		embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
			Name:   "Time",
			Value:  triggerTime.String(),
			Inline: false,
		})
	}

	return &embed
}

// Panic logs as panic level, with clean exit
func Panic(session *discordgo.Session, incomingError error, interaction *discordgo.Interaction) {
	embed := createEmbed(session, "panic", interaction)

	indentedError, _ := json.MarshalIndent(incomingError, "", "  ")
	embed.Description = string(indentedError)

	sentMessage, err := session.ChannelMessageSendComplex(logChannelID,
		// TODO: Enhance message
		&discordgo.MessageSend{
			// Content: "string",
			Embeds: []*discordgo.MessageEmbed{embed},
			// TTS: ,
			// Components: []discordgo.MessageComponent{},
			// Files: []*discordgo.File{},
			// AllowedMentions: &discordgo.MessageAllowedMentions{},
			// Reference: &discordgo.MessageReference{},
		})
	if err != nil {
		log.Panicf("Cannot send panic log: %+v\nReason: %+v", incomingError, err)
		return
	}

	log.Panicf("Panic: %+v(%s)", incomingError, util.MessageURL(logGuildID, logChannelID, sentMessage.ID))
}

// Error logs as error level, without exiting program
func Error(session *discordgo.Session, incomingError error, interaction *discordgo.Interaction) {
	embed := createEmbed(session, "panic", interaction)

	indentedError, _ := json.MarshalIndent(incomingError, "", "  ")
	embed.Description = string(indentedError)

	sentMessage, err := session.ChannelMessageSendComplex(logChannelID,
		// TODO: Enhance message
		&discordgo.MessageSend{
			// Content: "string",
			Embeds: []*discordgo.MessageEmbed{embed},
			// TTS: ,
			// Components: []discordgo.MessageComponent{},
			// Files: []*discordgo.File{},
			// AllowedMentions: &discordgo.MessageAllowedMentions{},
			// Reference: &discordgo.MessageReference{},
		})
	if err != nil {
		log.Printf("Cannot send error log: %+v\nReason: %+v", incomingError, err)
		return
	}

	log.Printf("Error: %+v(%s)", incomingError, util.MessageURL(logGuildID, logChannelID, sentMessage.ID))
}

// Warning logs as warning level
func Warning() {
	// TODO: Implement
}

// Info logs as info level
func Info() {
	// TODO: Implement
}

// Debug logs as debug level
func Debug(session *discordgo.Session, debugMessage string, interaction *discordgo.Interaction) {
	if !isDebug {
		return
	}

	embed := createEmbed(session, "debug", interaction)
	embed.Description = debugMessage
	// TODO: implement embedded debug message

	sentMessage, err := session.ChannelMessageSendComplex(logChannelID,
		// TODO: Enhance message
		&discordgo.MessageSend{
			// Content: "string",
			Embeds: []*discordgo.MessageEmbed{embed},
			// TTS: ,
			// Components: []discordgo.MessageComponent{},
			// Files: []*discordgo.File{},
			// AllowedMentions: &discordgo.MessageAllowedMentions{},
			// Reference: &discordgo.MessageReference{},
		})
	if err != nil {
		log.Printf("Cannot send debug log: %s\nReason: %+v", debugMessage, err)
		return
	}

	log.Printf("Debug: %s(%s)", debugMessage, util.MessageURL(logGuildID, logChannelID, sentMessage.ID))
}
