package df

import (
	"github.com/bwmarrin/discordgo"
)

// MessageContext Contains The Raw message and session objects for use in a Command's Handle method.
type MessageContext struct {
	Session *discordgo.Session
	Msg     *discordgo.Message
	Args    []string
}

// Command Something the bot should do.
type Command interface {
	Name() string
	Handle(*MessageContext) error   // Handle the command.
	PreFlight(*MessageContext) bool // Should the command be executed?
}
