package expression

import "fmt"

type BoxExpression struct {
	exp Expression
}

func (e BoxExpression) String() string {
	return fmt.Sprintf("( %s )", e.exp)
}
