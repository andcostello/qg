package qg

type JoinType int

const (
	LeftJoin JoinType = iota
	RightJoin
	InnerJoin
	OuterJoin
)

type JoinStmt struct {
	Type      JoinType
	Table     string
	Condition interface{}
}
