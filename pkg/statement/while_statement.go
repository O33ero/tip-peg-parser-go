package statement

import (
	"fmt"
	"tip-peg-parser-go/pkg/expression"
)

type WhileStatement struct {
	Condition expression.Expression
	Body      BodyStatement
}

func (s WhileStatement) GetToken() StmToken {
	return StmToken{Pattern: WhileExp}
}

func (s WhileStatement) String() string {
	return fmt.Sprintf("while ( %s ) %s", s.Condition, s.Body)
}
