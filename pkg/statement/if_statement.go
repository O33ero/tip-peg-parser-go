package statement

import (
	"tip-peg-parser-go/pkg/expression"
)

type IfStatement struct {
	condition expression.Expression
	body      BodyStatement
}

func (s IfStatement) GetToken() StmToken {
	return StmToken{Pattern: IfExp}
}

func (s IfStatement) String() string {
	return "If ( ? ) "
}
