package atomic

type IntExpression struct {
	value string
}

func (e IntExpression) String() string {
	return e.value
}
