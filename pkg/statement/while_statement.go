package statement

type WhileStatement struct {
}

func (s WhileStatement) GetToken() StmToken {
	return StmToken{Pattern: WhileExp}
}

func (s WhileStatement) String() string {
	return "While ( ? ) "
}
