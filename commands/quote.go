package commands

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/shppr/df"
	"github.com/shppr/df/db"
)

type Quote struct{}

func (Quote) Name() string {
	return "quote"
}

func (Quote) PreFlight(ctx *df.MessageContext) bool {
	return true
}

func (Quote) Handle(ctx *df.MessageContext) error {
	var qt db.Quote
	var err error
	rand.Seed(time.Now().UnixNano())

	if len(ctx.Args) == 0 {
		qt.Id = db.SafeId(rand.Int63())
	} else {
		id, err := strconv.ParseInt(ctx.Args[0], 0, 64)
		if err != nil {
			ctx.Session.ChannelMessageSend(ctx.Msg.ChannelID, "I can't find that quote.")
			return err
		}
		qt.Id = db.SafeId(id)
	}
	err = qt.Get()
	if err != nil {
		return err
	}

	resp := fmt.Sprintf("Quote %d of %d: %s", qt.Id, db.QuoteCount(), qt.Quote)

	_, err = ctx.Session.ChannelMessageSend(ctx.Msg.ChannelID, resp)
	return err
}
