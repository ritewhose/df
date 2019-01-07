package commands

import (
	"fmt"
	"strings"

	"github.com/shppr/df"
	"github.com/shppr/df/db"
)

type AddQuote struct{}

func (AddQuote) Name() string {
	return "addquote"
}

func (AddQuote) PreFlight(ctx *df.MessageContext) bool {
	return true
}

func (AddQuote) Handle(ctx *df.MessageContext) error {
	quote := strings.Join(ctx.Args, " ")
	qt := db.Quote{Creator: ctx.Msg.Author.String(), Quote: quote}

	err := qt.Add()
	if err != nil {
		ctx.Session.ChannelMessageSend(ctx.Msg.ChannelID, "I couldn't add that quote!")
		return err
	}

	resp := fmt.Sprintf("Quote %d added.", qt.Id)
	_, err = ctx.Session.ChannelMessageSend(ctx.Msg.ChannelID, resp)
	return err
}
