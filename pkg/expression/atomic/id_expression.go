package atomic

type IdExpression struct {
	id string
}

func (e IdExpression) String() string {
	return e.id
}
