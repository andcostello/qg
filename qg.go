package qg

import (
	"reflect"
)

// SQLable is an interface for anything that can be serialized into SQL.
type SQLable interface {
	ToSQL() (string, []interface{})
}

// StringParam is used box a string in order to mark it as a binding parameter.
// All other string arguments are considered raw sql.
type StringParam struct {
	Value string
}

// ToSQL implementation of the SQLable interface.
func (p StringParam) ToSQL() (string, []interface{}) {
	return "?", []interface{}{p.Value}
}

// Param creates a new StringParam.
func Param(v string) StringParam {
	return StringParam{v}
}

// AliasClause is used to build aliases. E.g. <exp> as <name>.
type AliasClause struct {
	Name  string
	Value interface{}
}

// ToSQL implements to the SQLable interface for AliasClause.
func (a AliasClause) ToSQL() (string, []interface{}) {
	s, b := wrapIfComplex(a.Value)
	return s + " as " + a.Name, b
}

// Alias creates a new AliasClause.
func Alias(v interface{}, name string) AliasClause {
	return AliasClause{name, v}
}

// Internal utility function for converting anything to sql.
func toSQL(v interface{}) (string, []interface{}) {
	switch t := v.(type) {
	case string:
		return t, nil
	case SQLable:
		return t.ToSQL()
	}

	t := reflect.ValueOf(v)
	switch t.Kind() {
	case reflect.Array, reflect.Slice:
		t.
	}
	return "?", []interface{}{v}
}

// Internal utility for dynamically wrapping sql in parentehsis if needed. (e.g. sub-select)
func wrapIfComplex(v interface{}) (string, []interface{}) {
	if t, ok := v.(StringParam); ok {
		return t.ToSQL()
	}
	if t, ok := v.(SQLable); ok {
		s, b := t.ToSQL()
		return "(" + s + ")", b
	}
	return toSQL(v)
}
