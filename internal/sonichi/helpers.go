package sonichi

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/bwmarrin/discordgo"
)

const (
	_mentionsRegex = `^/(\<\@\![A-Za-z0-9]\>)+/g`
)

func addMention(authorID, message string) string {
	return fmt.Sprintf("<@%s> %s", authorID, message)
}

// Ignore all messages created by the bot itself
func ignoreSelf(authorID, userID string) bool {
	return authorID == userID
}

// validate incoming commands against the command registry
func isValidCommand(rec string) (int, bool) {
	rec = strings.TrimSpace(rec)

	if len(rec) <= 0 {
		return -1, false
	}

	cmd := strings.Split(rec, " ")[0]

	if string(cmd[0]) == _commandPrefix {
		cmd = strings.TrimPrefix(cmd, _commandPrefix)
		for n, c := range commandRegistry {
			if c == cmd {
				return n, true
			}
		}
	}

	return -1, false
}

func trimMentions(message string) string {
	re := regexp.MustCompile(_mentionsRegex)
	return re.ReplaceAllString(message, "")
}

// verify if bot is mentioned
func isMentioned(m []*discordgo.User, botID string) bool {
	for _, u := range m {
		if u.ID == botID {
			return true
		}
	}

	return false
}
