package main

import (
	"github.com/shppr/df"
	"github.com/shppr/df/commands"
)

func main() {
	b, err := df.NewBotFromEnv()
	must(err)

	b.RegisterCommand(commands.In{})
	b.RegisterCommand(commands.Ask{})
	b.RegisterCommand(commands.M{})

	select {}
}

func must(e error) {
	if e != nil {
		panic(e)
	}
}
