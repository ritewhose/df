package commands

import (
	"strings"

	"github.com/mbags/df"
	"github.com/mbags/df/db"
)

type AddQuote struct {
	db.StorageDriver
}

func (AddQuote) Name() string {
	return "addquote"
}

func (AddQuote) PreFlight(ctx *df.MessageContext) bool {
	return true
}

func (aq AddQuote) Handle(ctx *df.MessageContext) error {
	exists, err := aq.Exists(ctx.Msg.Author.ID)
	if err != nil {
		return err
	}

	if !exists {
		u := &db.User{
			UserName: ctx.Msg.Author.ID,
		}

		err = aq.InsertUser(u)
		if err != nil {
			return err
		}
	}

	q := &db.Quote{
		Creator:  ctx.Msg.Author.ID,
		QuoteMsg: strings.Join(ctx.Args, " "),
	}

	err = aq.InsertQuote(q)
	if err != nil {
		return err
	}

	_, err = ctx.Session.ChannelMessageSend(ctx.Msg.ChannelID, "Quote added.")

	return err
}
