package expression

import "fmt"

type PlusExpression struct {
	BiExpression
}

func (e PlusExpression) String() string {
	return fmt.Sprintf("%s + %s", e.left, e.right)
}
