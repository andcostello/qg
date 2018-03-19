package qg

import (
	"strings"
)

// SelectStmt represents the components of a select statement.
type SelectStmt struct {
	Columns         []interface{}
	IsDistinct      bool
	DistinctColumns []interface{}
	Condition       interface{}
	Table           interface{}
	Joins           []JoinStmt
	OrderStatements []OrderStmt
}

// Select builds a new select statement with the given columns.
func Select(v ...interface{}) *SelectStmt {
	return &SelectStmt{Columns: v}
}

// Distinct marks the select statement to only return distinct results.
func (s *SelectStmt) Distinct() *SelectStmt {
	s.IsDistinct = true
	return s
}

// DistinctOn marks the select statement to only return distinct results
// based on the provided expressions.
func (s *SelectStmt) DistinctOn(v ...interface{}) *SelectStmt {
	s.IsDistinct = true
	s.DistinctColumns = append(s.DistinctColumns, v...)
	return s
}

// From specifies the clause to select from.
func (s *SelectStmt) From(v interface{}) *SelectStmt {
	s.Table = v
	return s
}

// Where adds the select's condition clause.
func (s *SelectStmt) Where(v interface{}) *SelectStmt {
	s.Condition = v
	return s
}

// ToSQL implements the SQLable interface for SelectStmt.
func (s *SelectStmt) ToSQL() (string, []interface{}) {
	sql := []string{"select"}
	var bindings []interface{}

	if s.IsDistinct {
		sql = append(sql, "distinct")
		if len(s.DistinctColumns) > 0 {
			var columns []string
			for _, column := range s.DistinctColumns {
				s, b := wrapIfComplex(column)
				columns = append(columns, s)
				bindings = append(bindings, b...)
			}
			sql = append(sql, "on ("+strings.Join(columns, ", ")+")")
		}
	}

	if len(s.Columns) > 0 {
		var columns []string
		for _, column := range s.Columns {
			s, b := wrapIfComplex(column)
			columns = append(columns, s)
			bindings = append(bindings, b...)
		}
		sql = append(sql, strings.Join(columns, ", "))
	}

	if s.Table != nil {
		s, b := wrapIfComplex(s.Table)
		sql = append(sql, "from", s)
		bindings = append(bindings, b...)
	}

	for _, join := range s.Joins {
		s, b := toSQL(join)
		sql = append(sql, s)
		bindings = append(bindings, b...)
	}

	if s.Condition != nil {
		s, b := toSQL(s.Condition)
		sql = append(sql, "where", s)
		bindings = append(bindings, b...)
	}

	if len(s.OrderStatements) > 0 {
		orders := []string{"order by"}
		for _, order := range s.OrderStatements {
			s, b := toSQL(order)
			orders = append(orders, s)
			bindings = append(bindings, b...)
		}
		sql = append(sql, orders...)
	}

	return strings.Join(sql, " "), bindings
}
