package sonichi

import (
	"errors"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type Controller interface {
	Help(s *discordgo.Session, m *discordgo.Event)
	SayHello(s *discordgo.Session, m *discordgo.MessageCreate)
	Welcome(s *discordgo.Session, m *discordgo.GuildMemberAdd)
}

type Operator struct {
	config  Config
	session *discordgo.Session
}

type Config struct {
	Token    string `mapstructure:"token"`
	AuthType string `mapstructure:"authType"`
}

// New creates a pointer to a new Operator wrapping a discord session
func New(cfg Config) (*Operator, error) {
	// TODO: sanitize config values
	session, err := discordgo.New("Bot " + cfg.Token)
	if err != nil {
		return nil, fmt.Errorf("failed to create new discord session, %w", err)
	}

	op := &Operator{
		config:  cfg,
		session: session,
	}

	return op, nil
}

// BindHandlers will add any amount of valid handler funcs defined by the discordgo handler registry
func BindHandlers(op interface{}, handlers ...interface{}) error {
	if _, ok := op.(*Operator); !ok {
		return errors.New("can only bind handlers to a pointer to an operator")
	}

	for _, h := range handlers {
		op.(*Operator).session.AddHandler(h)
	}

	return nil
}

// Start wraps the discordgo session websocket connection
func Start(op interface{}) error {
	if _, ok := op.(*Operator); !ok {
		return errors.New("can only bind handlers to a pointer to an operator")
	}

	return op.(*Operator).session.Open()
}
