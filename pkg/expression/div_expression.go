package expression

import "fmt"

type DivExpression struct {
	BiExpression
}

func (e DivExpression) String() string {
	return fmt.Sprintf("%s / %s", e.left, e.right)
}
