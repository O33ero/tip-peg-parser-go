package statement

type EqualStatement struct {
}

func (s EqualStatement) GetToken() StmToken {
	return StmToken{Pattern: IdEqExp}
}

func (s EqualStatement) String() string {
	return "eq ? = ?"
}
