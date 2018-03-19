package qg

// JoinType enumerates the list of possible sql joins.
type JoinType int

// All supported join types.
const (
	LeftJoin JoinType = iota
	RightJoin
	InnerJoin
	OuterJoin
)

// String is a Stringer implementation for JoinType.
func (j JoinType) String() string {
	return []string{"left", "right", "inner", "outer"}[j]
}

// JoinStmt represents a sql join statement.
type JoinStmt struct {
	Type      JoinType
	Right     interface{}
	Condition interface{}
}

// LeftJoin builds a sql left join.
func (s *SelectStmt) LeftJoin(right interface{}, condition interface{}) *SelectStmt {
	s.Joins = append(s.Joins, JoinStmt{Type: LeftJoin, Right: right, Condition: condition})
	return s
}

// RightJoin builds a sql right join.
func (s *SelectStmt) RightJoin(right interface{}, condition interface{}) *SelectStmt {
	s.Joins = append(s.Joins, JoinStmt{Type: RightJoin, Right: right, Condition: condition})
	return s
}

// InnerJoin builds a sql inner join.
func (s *SelectStmt) InnerJoin(right interface{}, condition interface{}) *SelectStmt {
	s.Joins = append(s.Joins, JoinStmt{Type: InnerJoin, Right: right, Condition: condition})
	return s
}

// OuterJoin builds a sql outer join.
func (s *SelectStmt) OuterJoin(right interface{}, condition interface{}) *SelectStmt {
	s.Joins = append(s.Joins, JoinStmt{Type: OuterJoin, Right: right, Condition: condition})
	return s
}

// ToSQL implements the SQLable interface for JoinStmt.
func (j JoinStmt) ToSQL() (string, []interface{}) {
	rs, rb := toSQL(j.Right)
	cs, cb := toSQL(j.Condition)

	if _, ok := j.Condition.(UsingStmt); ok {
		return j.Type.String() + " join " + rs + " using (" + cs + ")", append(rb, cb...)
	}

	return j.Type.String() + " join " + rs + " on " + cs, append(rb, cb...)
}

// UsingStmt is a sql using condition for a join clause.
type UsingStmt struct {
	Condition interface{}
}

// Using constructs a new UsingStmt with the given value.
func Using(v interface{}) UsingStmt {
	return UsingStmt{Condition: v}
}

// ToSQL implements the SQLable interface for UsingStmt.
func (u UsingStmt) ToSQL() (string, []interface{}) {
	return toSQL(u.Condition)
}
