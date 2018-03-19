package qg

import "testing"

func TestJoinEq(t *testing.T) {
	assertSQL(t,
		"select (? = ?)",
		[]interface{}{1, 2},
		Select(Eq(1, 2)))

	assertSQL(t,
		"select id from table where id = ?",
		[]interface{}{1},
		Select("id").From("table").Where(Eq("id", 1)))

	assertSQL(t,
		"select ((id = ?) as result), created from table",
		[]interface{}{1},
		Select(Alias(Eq("id", 1), "result"), "created").From("table"))

	assertSQL(t,
		"select id from table where id = (select ?)",
		[]interface{}{1},
		Select("id").From("table").Where(Eq("id", Select(1))))
}
