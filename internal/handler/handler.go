package handler

import (
	"github.com/bwmarrin/discordgo"
)

var handlers []Handler

type Handler interface {
	add(session *discordgo.Session)
}

type genericHandler struct {
	handler interface{}
}

func (g *genericHandler) add(session *discordgo.Session) {
	session.AddHandler(g.handler)
}

func init() {
	handlers = []Handler{
		// Append all custom handlers
		pingPong,
	}
}

func Add(session *discordgo.Session) {
	for _, custom := range handlers {
		custom.add(session)
	}
}
