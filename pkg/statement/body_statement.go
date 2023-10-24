package statement

import "fmt"

type BodyStatement struct {
	statements []Statement
}

func (s BodyStatement) String() string {
	return fmt.Sprintf("{ %s }", s.statements)
}
