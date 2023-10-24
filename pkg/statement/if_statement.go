package statement

type IfStatement struct {
}

func (s IfStatement) GetToken() StmToken {
	return StmToken{Pattern: IfExp}
}

func (s IfStatement) String() string {
	return "If ( ? ) "
}
