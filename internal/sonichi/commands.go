package sonichi

import (
	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
)

const (
	_commandPrefix = "!"
)

var (
	commandRegistry = []string{
		"hello",
		"help",
	}
)

// Help mentions the author with a list of this bot's commands
func (op *Operator) Help(s *discordgo.Session, m *discordgo.Event) {
	return
}

// SayHello mentions the author with a "nice" greeting
func (op *Operator) SayHello(s *discordgo.Session, m *discordgo.MessageCreate) {
	if ignoreSelf(m.Author.ID, s.State.User.ID) {
		return
	}

	regIdx, valid := isValidCommand(m.Content)
	if !valid {
		return
	}

	msg := addMention(m.Author.ID, "Hello! :yum:")

	if commandRegistry[regIdx] == "hello" {
		if _, err := s.ChannelMessageSend(m.ChannelID, msg); err != nil {
			log.Error().Err(err).Msg("failed to send message")
		}
	}
}

// Welcome mentions a new user when joining the server
func (op *Operator) Welcome(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
	for _, g := range s.State.Guilds {
		if g.Name == "Green Vipers" {
			for _, c := range g.Channels {
				if c.Name == "welcome" {
					log.Debug().Msg("found welcome")
					msg := addMention(m.User.ID, "Welcome to the Green Vipers :sunglasses:")

					if _, err := s.ChannelMessageSend(c.ID, msg); err != nil {
						log.Error().Err(err).Msg("failed to send message")
					}
				}
			}
		}
	}
}
