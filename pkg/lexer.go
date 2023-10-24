package pkg

import (
	"errors"
	"fmt"
	"strings"
	"tip-peg-parser-go/pkg/statement"
)

func Parse(lines []string) {
	flatted := lineFlatter(lines)
	statements, _ := statementLexer(flatted)

	fmt.Printf("%s", statements)
}

func statementLexer(lines []string) ([]statement.Statement, error) {
	var (
		output    []statement.Statement
		stmTokens = statement.GetAllStatementTokens()
	)

	prev := ""
	for _, line := range lines {
		next := prev + " " + line

		statement, err := tryToGuessStatement(next, stmTokens)
		if err != nil {
			fmt.Printf("Cannot recognize token: %s\n", prev)
			prev = next
			continue
		}

		output = append(output, statement)

		fmt.Printf("Resolve token: %s\n", next)
		pattern, _ := statement.GetToken().GetPattern()
		replaced := pattern.ReplaceAllString(next, "")
		prev = strings.TrimSpace(replaced)
	}

	if strings.TrimSpace(prev) != "" {
		fmt.Printf("Cannot recognize statement: %s\n", prev)
		return []statement.Statement{}, errors.New("cannot recognize statement: " + prev)
	}

	return output, nil
}

func tryToGuessStatement(line string, tokens []statement.StmToken) (statement.Statement, error) {
	for _, token := range tokens {
		pattern, _ := token.GetPattern()
		match := pattern.FindString(line)
		if match != "" {
			switch token.Pattern {
			case statement.IfElse:
				return statement.IfStatement{}, nil
			case statement.WhileExp:
				return statement.WhileStatement{}, nil
			case statement.IfExp:
				return statement.ElseStatement{}, nil
			case statement.OutputExp:
				return statement.OutputStatement{}, nil
			case statement.EndStm:
				return statement.EndBodyStatement{}, nil
			case statement.IdEqExp:
				return statement.EqualStatement{}, nil
			}
		}
	}
	return nil, errors.New("cannot recognize token")
}

func lineFlatter(lines []string) []string {
	var output []string
	for _, line := range lines {
		line = strings.ReplaceAll(line, "(", " ( ")
		line = strings.ReplaceAll(line, ")", " ) ")
		line = strings.ReplaceAll(line, ";", " ; ")
		for _, split := range strings.Split(line, " ") {
			if split != "" {
				output = append(output, split)
			}
		}
	}

	fmt.Printf("Flatter: %s\n", output)
	return output
}
