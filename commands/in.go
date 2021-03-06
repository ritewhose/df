package commands

import (
	"github.com/mbags/df"
	"fmt"
	"time"
	"strings"
	"errors"
)

type In struct{}

func (In) Name() string {
	return "in"
}

func (In) Handle(ctx *df.MessageContext) error {
	if len(ctx.Args) == 0 {
		e := errors.New("Invalid duration from user")
		explainFail(ctx, e, "Provide a duration....")
		return e
	}

	replyMsg := fmt.Sprintf("<@%s> ", ctx.Msg.Author.ID)

	dur, err := time.ParseDuration(ctx.Args[0])
	if err != nil {
		explainFail(ctx, err, "Invalid duration.")
		return err
	}

	if len(ctx.Args) >= 2 {
		replyMsg += strings.Join(ctx.Args[1:], " ")
	}
	go func() {
		<-time.After(dur)
		ctx.Session.ChannelMessageSend(ctx.Msg.ChannelID, replyMsg)
	}()

	return nil
}

func (In) PreFlight(ctx *df.MessageContext) bool {
	return true
}

func explainFail(ctx *df.MessageContext, err error, msg string) bool {
	if err != nil {
		ctx.Session.ChannelMessageSend(ctx.Msg.ChannelID, msg)
		return true
	}
	return false
}
