package commands

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"github.com/shppr/df"
)

var conversionFactors = map[string]float64{
	"y":  365.25 * 24 * 60 * 60,
	"mo": 29.53059 * 24 * 60 * 60,
	"d":  24 * 60 * 60,
	"h":  60 * 60,
	"m":  60,
	"s":  1,
}

type In struct {}

func (In) Name() string {
	return "in"
}

func (In) PreFlight(ctx *df.MessageContext) bool {
	return true
}

func (In) Handle(ctx *df.MessageContext) error {
	if len(ctx.Args) == 0 {
		ctx.Session.ChannelMessageSend(ctx.Msg.ChannelID, "Usage: `.in [quantity][y,mo,d,h,m,s] [message]`")
		return nil
	}
	unit := strings.TrimLeft(ctx.Args[0], "0123456789")
	durationString := strings.TrimSuffix(ctx.Args[0], unit)
	conversion, ok := conversionFactors[unit]
	if !ok {
		ctx.Session.ChannelMessageSend(ctx.Msg.ChannelID, "Invalid duration.")
		return nil
	}

	durationNum, err := strconv.ParseFloat(durationString, 64)
	if err != nil {
		ctx.Session.ChannelMessageSend(ctx.Msg.ChannelID, "Invalid input.")
		return nil
	}

	durationConverted := int(conversion * durationNum)

	replyMsg := fmt.Sprintf("<@%s>: ", ctx.Msg.Author.ID)

	if len(ctx.Args) > 1 {
		replyMsg += strings.Join(ctx.Args[1:], " ")
	}

	dur, err := time.ParseDuration(fmt.Sprintf("%ds", durationConverted))
	if err != nil {
		ctx.Session.ChannelMessageSend(ctx.Msg.ChannelID, "Duration too long (max = 292 years).")
		return fmt.Errorf("[timer] %s", err)
	}

	replyTime := time.Now().Add(dur)
	ackMsg := fmt.Sprintf("See you at: %s!", replyTime.UTC().Format(time.UnixDate))
	ctx.Session.ChannelMessageSend(ctx.Msg.ChannelID, ackMsg)

	go func() {
		<-time.After(time.Duration(dur))
		ctx.Session.ChannelMessageSend(ctx.Msg.ChannelID, replyMsg)
	}()

	return nil
}
