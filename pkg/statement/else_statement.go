package statement

type ElseStatement struct {
}

func (s ElseStatement) GetToken() StmToken {
	return StmToken{Pattern: IfElse}
}

func (s ElseStatement) String() string {
	return "Else "
}
