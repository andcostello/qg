package qg

// SQLable is an interface for anything that can be serialized into SQL.
type SQLable interface {
	ToSQL() (string, []interface{})
}

// Param is used box a string in order to mark it as a binding parameter.
// All other string arguments are considered raw sql.
type param struct {
	Value string
}

// ToSQL implementation of the SQLable interface
func (p param) ToSQL() (string, []interface{}) {
	return "?", []interface{}{p.Value}
}

func Param(v string) param {
	return param{v}
}

type alias struct {
	name string
	v    interface{}
}

func (a alias) ToSQL() (string, []interface{}) {
	s, b := wrapIfComplex(a.v)
	return s + " as " + a.name, b
}

func Alias(v interface{}, name string) alias {
	return alias{name, v}
}

// Core utility function for converting anything to sql
func toSQL(v interface{}) (string, []interface{}) {
	if t, ok := v.(string); ok {
		return t, nil
	}
	if t, ok := v.(SQLable); ok {
		return t.ToSQL()
	}
	return "?", []interface{}{v}
}

// Core utility for dynamically wrapping sql in parentehsis if needed. (e.g. sub-select)
func wrapIfComplex(v interface{}) (string, []interface{}) {
	if t, ok := v.(param); ok {
		return t.ToSQL()
	}
	if t, ok := v.(SQLable); ok {
		s, b := t.ToSQL()
		return "(" + s + ")", b
	}
	return toSQL(v)
}
