package cmdparser

import (
	"errors"
	"strings"
)

var (
	ErrNotCommand = errors.New("cmdparser: Not a command (no prefix)")
)

func Parse(s string) (
	name string,
	positionals []string,
	err error,
) {
	if !strings.HasPrefix(s, "#") {
		err = ErrNotCommand
		return
	}

	// remove prefix from s
	s = s[1:]

	var nextSpace int
	name, nextSpace = nextToken(s)
	if nextSpace == -1 {
		// no more tokens (no flags, no positionals)
		return
	}
	s = s[nextSpace:]

	var token string
	for {
		token, nextSpace = nextToken(s)

		if len(token) != 0 {
			// if token is not blank
			positionals = append(positionals, token)
		} else {
			token = " "
		}

		if nextSpace == -1 {
			// no more tokens (no flags, no positionals)
			return
		}

		s = s[len(token):]
	}
}

func nextToken(s string) (string, int) {
	nextSpace := strings.Index(s, " ")
	if nextSpace == -1 {
		return s, -1
	}

	return s[:nextSpace], nextSpace
}
