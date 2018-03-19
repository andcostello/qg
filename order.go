package qg

// OrderType is a sort order.
type OrderType int

// All supported order types.
const (
	NoOrder OrderType = iota
	AscendingOrder
	DescendingOrder
)

// String is a Stringer implementation for OrderType.
func (o OrderType) String() string {
	return []string{"", "asc", "desc"}[o]
}

// OrderStmt represents an ordering clause.
type OrderStmt struct {
	Type  OrderType
	Value interface{}
}

// ToSQL implements the SQLable interface for OrderStmt.
func (o OrderStmt) ToSQL() (string, []interface{}) {
	s, b := wrapIfComplex(o.Value)
	if o.Type == NoOrder {
		return s, b
	}
	return s + " " + o.Type.String(), b
}

// OrderBy appends a new order clause to the select.
func (s *SelectStmt) OrderBy(v interface{}) *SelectStmt {
	s.OrderStatements = append(s.OrderStatements, OrderStmt{NoOrder, v})
	return s
}

// OrderByAsc appends a new ascending order clause to the select.
func (s *SelectStmt) OrderByAsc(v interface{}) *SelectStmt {
	s.OrderStatements = append(s.OrderStatements, OrderStmt{AscendingOrder, v})
	return s
}

// OrderByDesc appends a new descending order clause to the select.
func (s *SelectStmt) OrderByDesc(v interface{}) *SelectStmt {
	s.OrderStatements = append(s.OrderStatements, OrderStmt{DescendingOrder, v})
	return s
}
