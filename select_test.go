package qg

import "testing"

func TestSelectBasic(t *testing.T) {

	assertSQL(t,
		"select",
		[]interface{}{},
		Select(),
	)

	assertSQL(t,
		"select a",
		[]interface{}{},
		Select("a"),
	)

	assertSQL(t,
		"select ?",
		[]interface{}{1},
		Select(1),
	)

	assertSQL(t,
		"select a, b, c",
		[]interface{}{},
		Select("a", "b", "c"),
	)

	assertSQL(t,
		"select ?, ?, ?",
		[]interface{}{1, 2, 3},
		Select(1, 2, 3),
	)

	assertSQL(
		t,
		"select id from table",
		[]interface{}{},
		Select("id").From("table"),
	)

	assertSQL(t,
		"select id from table t",
		[]interface{}{},
		Select("id").From("table t"),
	)
}

func TestSelectWhere(t *testing.T) {
	assertSQL(t,
		"select id from table where ?",
		[]interface{}{true},
		Select("id").From("table").Where(true),
	)
}

func TestSelectSubSelect(t *testing.T) {
	assertSQL(t,
		"select id from table where id = (select ?)",
		[]interface{}{1},
		Select("id").From("table").Where(Eq("id", Select(1))),
	)

	assertSQL(t,
		"select (select (select ?))",
		[]interface{}{1},
		Select((Select(Select(1)))),
	)
}

func TestSelectTypes(t *testing.T) {
	assertSQL(t,
		"select id, ?, ?, ?, ?, ?",
		[]interface{}{1, 1.2, true, 'a', uint32(1)},
		Select("id", 1, 1.2, true, 'a', uint32(1)),
	)

	assertSQL(t,
		"select name from table where id = ?",
		[]interface{}{"blah"},
		Select("name").From("table").Where(Eq("id", Param("blah"))))
}
