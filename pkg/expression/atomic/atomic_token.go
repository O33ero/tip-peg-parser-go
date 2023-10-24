package atomic

import "regexp"

type Pattern string

const (
	ID    Pattern = "([a-zA-Z]+)"
	INT           = "([\\-0-9]+)"
	INPUT         = "input"
)

type AtmToken struct {
	pattern Pattern
}

func (token AtmToken) GetPattern() (*regexp.Regexp, error) {
	return regexp.Compile(string(token.pattern))
}
