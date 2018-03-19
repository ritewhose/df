package df

import (
	"log"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// Bot Root object for handling commands.
type Bot struct {
	*discordgo.Session
	prefix     string
	commandMap map[string]Command
}

// NewBotFromEnv Construct a bot from environment variables: token, prefix
func NewBotFromEnv() (*Bot, error) {
	token := os.Getenv("testtoken")

	pre := "." // get from env
	sess, err := discordgo.New("Bot " + token)

	sess.Open()
	if err != nil {
		return nil, err
	}

	b := &Bot{
		Session:    sess,
		prefix:     pre,
		commandMap: make(map[string]Command),
	}

	b.AddRawHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		b.HandleCommand(m.Message)
	})

	return b, nil
}

func (b *Bot) RegisterCommand(cmd Command) {
	b.commandMap[cmd.Name()] = cmd
}

func (b *Bot) AddRawHandler(handler interface{}) {
	b.Session.AddHandler(handler)
}

func (b *Bot) HandleCommand(msg *discordgo.Message) {
	tokens := strings.Split(msg.Content, " ")
	cmdName, validCommand := b.isCommand(tokens[0])

	if !validCommand {
		return
	}

	ctx := &MessageContext{
		Msg:     msg,
		Session: b.Session,
	}

	if len(tokens) > 1 {
		ctx.Args = tokens[1:]
	}

	cmd := b.commandMap[cmdName]
	if !cmd.PreFlight(ctx) {
		return
	}

	e := cmd.Handle(ctx)
	if e != nil {
		log.Printf("[HandleCommand] %s\n", e)
	}
}

func (b *Bot) isCommand(s string) (string, bool) {
	if len(s) == 1 {
		return "", false
	}
	if strings.HasPrefix(s, b.prefix) {
		if _, ok := b.commandMap[s[1:]]; ok {
			return s[1:], true
		}
	}
	return "", false
}
