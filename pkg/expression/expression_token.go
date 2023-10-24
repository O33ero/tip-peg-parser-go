package expression

import "regexp"

type Pattern string

const (
	ExpPlusExp  Pattern = "(.*) \\+ (.*)"
	ExpMinusExp         = "(.*) \\- (.*)"
	ExpMultiExp         = "(.*) \\* (.*)"
	ExpDivExp           = "(.*) \\/ (.*)"
	ExpGrtExp           = "(.*) \\> (.*)"
	ExpEqExp            = "(.*) == (.*)"
	ExpBox              = "\\( (.*) \\)"
)

type ExpToken struct {
	pattern Pattern
}

func (token ExpToken) GetPattern() (*regexp.Regexp, error) {
	return regexp.Compile(string(token.pattern))
}
