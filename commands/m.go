package commands

import (
	"errors"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/shppr/df"
)

type M struct{}

func (M) Name() string {
	return "m"
}

func (M) Handle(ctx *df.MessageContext) error {
	if len(ctx.Args) != 1 || len(ctx.Msg.Mentions) != 1 {
		e := errors.New("Invalid mention from user")
		backPat(ctx, ctx.Msg.Author)
		return e
	}
	backPat(ctx, ctx.Msg.Mentions[0])

	return nil
}

func (M) PreFlight(ctx *df.MessageContext) bool {
	return true
}

func backPat(ctx *df.MessageContext, user *discordgo.User) {
	replyMsg := fmt.Sprintf("You're doing good work, %s!", user.Mention())
	ctx.Session.ChannelMessageSend(ctx.Msg.ChannelID, replyMsg)
	return
}
