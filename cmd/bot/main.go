package main

import (
	"github.com/mbags/df"
	"github.com/mbags/df/commands"
)

func main() {
	b, err := df.NewBotFromEnv()
	must(err)

	b.RegisterCommand(commands.In{})
	b.RegisterCommand(commands.Ask{})

	select {}
}

func must(e error) {
	if e != nil {
		panic(e)
	}
}