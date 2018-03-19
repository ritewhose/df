package commands

import (
	"fmt"
	"strconv"

	"github.com/mbags/df"
	"github.com/mbags/df/db"
)

type Quote struct {
	db.StorageDriver
}

func (Quote) Name() string {
	return "quote"
}

func (Quote) PreFlight() bool {
	return true
}

func (q Quote) Handle(ctx *df.MessageContext) error {
	var quote *db.Quote

	quoteNum, err := strconv.Atoi(ctx.Args[0])
	if err != nil {
		return err
	}

	if len(ctx.Args) > 0 {
		quote, err = q.SelectQuoteByNumber(quoteNum)
	} else {
		quote, err = q.SelectRandomQuote()
	}
	if err != nil {
		return err
	}

	_, err = ctx.Session.ChannelMessageSend(ctx.Msg.ChannelID, fmt.Sprint(quote))

	return err
}
