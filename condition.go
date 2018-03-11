package qg

type Condition struct {
	Operator string
	Left     interface{}
	Right    interface{}
}

func Eq(a, b interface{}) *Condition {
	return &Condition{
		Operator: "=",
		Left:     a,
		Right:    b,
	}
}

func (c *Condition) ToSQL() (string, []interface{}) {
	sl, bl := wrapIfComplex(c.Left)
	sr, br := wrapIfComplex(c.Right)
	return sl + " " + c.Operator + " " + sr, append(bl, br...)
}
