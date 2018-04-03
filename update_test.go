package qg

import "testing"

func TestUpdateBasic(t *testing.T) {
	assertSQL(t,
		"update table set id = ?",
		[]interface{}{1},
		Update("table").Set("id", 1))

	assertSQL(t,
		"update table set (a, b) = (select ?, ?)",
		[]interface{}{1, 2},
		Update("table").Set([]string{"a", "b"}, Select(1, 2)))

	assertSQL(t,
		"update table set (a, b) = (?, ?)",
		[]interface{}{1, 2},
		Update("table").Set([]string{"a", "b"}, []int{1, 2}))
}
