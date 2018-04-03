package qg

import (
	"strings"
)

// UpdateStmt represents the components of an update statement.
type UpdateStmt struct {
	Table    interface{}
	SetStmts []setStmt
}

type setStmt struct {
	left  interface{}
	right interface{}
}

// Update builds a new update statement for the given table. Can be a string or Alias.
func Update(v interface{}) *UpdateStmt {
	return &UpdateStmt{Table: v}
}

// Set adds a set statment to the query, with the left side column(s) and the right side expression.
func (u *UpdateStmt) Set(left interface{}, right interface{}) *UpdateStmt {
	u.SetStmts = append(u.SetStmts, setStmt{left, right})
	return u
}

// ToSQL implements the SQLable interface for UpdateStmt.
func (u *UpdateStmt) ToSQL() (string, []interface{}) {
	sql := []string{"update"}
	var bindings []interface{}

	if u.Table != nil {
		s, b := toSQL(u.Table)
		sql = append(sql, s)
		bindings = append(bindings, b...)
	}

	sql = append(sql, "set")

	for _, setStmt := range u.SetStmts {
		s, b := wrapIfComplex(setStmt.right)
		switch left := setStmt.left.(type) {
		case string:
			sql = append(sql, left+" = "+s)
			bindings = append(bindings, b...)
		case []string:
			sql = append(sql, "("+strings.Join(left, ", ")+") = "+s)
			bindings = append(bindings, b...)
		}
	}

	return strings.Join(sql, " "), bindings
}
