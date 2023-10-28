package statement

import "fmt"

type BodyStatement struct {
	Statements []Statement
}

func (s *BodyStatement) String() string {
	return fmt.Sprintf("{ %s }", s.Statements)
}

func (s *BodyStatement) Put(statement Statement) {
	s.Statements = append(s.Statements, statement)
}

func (s *BodyStatement) GetBodyStatement() []Statement {
	return s.Statements
}
