package commands

import (
	"math/rand"
	"strings"
	"time"

	"github.com/mbags/df"
)

type Ask struct{}

func (Ask) Name() string {
	return "ask"
}

func (Ask) PreFlight(ctx *df.MessageContext) bool {
	return true
}

func (Ask) Handle(ctx *df.MessageContext) error {
	rand.Seed(time.Now().UnixNano())

	message := strings.Join(ctx.Args, " ")

	if strings.Contains(message, " or ") {
		substrs := strings.Split(message, " or ")
		_, err := ctx.Session.ChannelMessageSend(ctx.Msg.ChannelID, substrs[rand.Intn(len(substrs))])
		return err
	}

	_, err := ctx.Session.ChannelMessageSend(ctx.Msg.ChannelID, []string{"yes", "no"}[rand.Intn(2)])
	return err
}
