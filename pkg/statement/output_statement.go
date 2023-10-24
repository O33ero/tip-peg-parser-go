package statement

type OutputStatement struct {
}

func (s OutputStatement) GetToken() StmToken {
	return StmToken{Pattern: OutputExp}
}

func (s OutputStatement) String() string {
	return "output ?"
}
