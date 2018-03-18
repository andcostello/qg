package qg

type JoinType int

const (
	LeftJoin JoinType = iota
	RightJoin
	InnerJoin
	OuterJoin
)

func (j JoinType) String() string {
	return []string{"left", "right", "inner", "outer"}[j]
}

type JoinStmt struct {
	Type      JoinType
	Right     interface{}
	Condition interface{}
}

func (s *SelectStmt) LeftJoin(right interface{}, condition interface{}) *SelectStmt {
	s.Joins = append(s.Joins, JoinStmt{Type: LeftJoin, Right: right, Condition: condition})
	return s
}

func (s *SelectStmt) RightJoin(right interface{}, condition interface{}) *SelectStmt {
	s.Joins = append(s.Joins, JoinStmt{Type: RightJoin, Right: right, Condition: condition})
	return s
}

func (s *SelectStmt) InnerJoin(right interface{}, condition interface{}) *SelectStmt {
	s.Joins = append(s.Joins, JoinStmt{Type: InnerJoin, Right: right, Condition: condition})
	return s
}

func (s *SelectStmt) OuterJoin(right interface{}, condition interface{}) *SelectStmt {
	s.Joins = append(s.Joins, JoinStmt{Type: OuterJoin, Right: right, Condition: condition})
	return s
}

func (j JoinStmt) ToSQL() (string, []interface{}) {
	rs, rb := toSQL(j.Right)
	cs, cb := toSQL(j.Condition)

	if _, ok := j.Condition.(UsingStmt); ok {
		return j.Type.String() + " join " + rs + " using (" + cs + ")", append(rb, cb...)
	}

	return j.Type.String() + " join " + rs + " on " + cs, append(rb, cb...)
}

type UsingStmt struct {
	Condition interface{}
}

func Using(v interface{}) UsingStmt {
	return UsingStmt{Condition: v}
}

func (u UsingStmt) ToSQL() (string, []interface{}) {
	return toSQL(u.Condition)
}
