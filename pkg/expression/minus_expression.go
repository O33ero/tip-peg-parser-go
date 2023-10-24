package expression

import "fmt"

type MinusExpression struct {
	BiExpression
}

func (e MinusExpression) String() string {
	return fmt.Sprintf("%s - %s", e.left, e.right)
}
