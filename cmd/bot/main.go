package main

import (
	"log"

	"github.com/mbags/df"
	"github.com/mbags/df/commands"
	"github.com/mbags/df/db"
)

func main() {
	b, err := df.NewBotFromEnv()
	must(err)

	db, err := db.DialDB("./gtb.db")
	if err != nil {
		log.Fatalf("Couldn't dial db: %s", err)
	}

	b.RegisterCommand(commands.In{})
	b.RegisterCommand(commands.Ask{})
	b.RegisterCommand(commands.AddQuote{db})

	select {}
}

func must(e error) {
	if e != nil {
		panic(e)
	}
}
