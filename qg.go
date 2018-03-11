package qg

// SQLable is the base interface for anything that can be serialized into SQL
type SQLable interface {
	ToSQL() (string, []interface{})
}

type Param struct {
	v string
}

func (p Param) ToSQL() (string, []interface{}) {
	return "?", []interface{}{p.v}
}

func toSQL(v interface{}) (string, []interface{}) {
	if t, ok := v.(string); ok {
		return t, nil
	}
	if t, ok := v.(SQLable); ok {
		return t.ToSQL()
	}
	return "?", []interface{}{v}
}

func wrapIfComplex(v interface{}) (string, []interface{}) {
	if t, ok := v.(Param); ok {
		return t.ToSQL()
	}
	if t, ok := v.(SQLable); ok {
		s, b := t.ToSQL()
		return "(" + s + ")", b
	}
	return toSQL(v)
}
