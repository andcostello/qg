package qg

// Condition represents a comparison operation.
type Condition struct {
	Operator string
	Left     interface{}
	Right    interface{}
}

// Eq implements the equals operator.
func Eq(a, b interface{}) Condition {
	return Condition{
		Operator: "=",
		Left:     a,
		Right:    b,
	}
}

// ToSQL implements the SQLable interface for Condition.
func (c Condition) ToSQL() (string, []interface{}) {
	sl, bl := wrapIfComplex(c.Left)
	sr, br := wrapIfComplex(c.Right)
	return sl + " " + c.Operator + " " + sr, append(bl, br...)
}
