package expression

import "fmt"

type MultiExpression struct {
	BiExpression
}

func (e MultiExpression) String() string {
	return fmt.Sprintf("%s * %s", e.left, e.right)
}
