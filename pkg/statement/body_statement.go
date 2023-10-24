package statement

import "fmt"

type BodyStatement struct {
	statements []Statement
}

func (s BodyStatement) String() string {
	return fmt.Sprintf("{ %s }", s.statements)
}

func (s BodyStatement) Put(statement Statement) {
	s.statements = append(s.statements, statement)
}

func (s BodyStatement) GetBodyStatement() []Statement {
	return s.statements
}
