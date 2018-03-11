package qg

import "testing"

func assertSQL(t *testing.T, expectedSQL string, expectedBindings []interface{}, query SQLable) {
	sql, bindings := query.ToSQL()
	if expectedSQL != sql {
		t.Errorf("Expected '%s', got '%s'", expectedSQL, sql)
	}

	if len(expectedBindings) != len(bindings) {
		t.Errorf("Expected %d bindings, got %d", len(expectedBindings), len(bindings))
		return
	}

	for i := range expectedBindings {
		if expectedBindings[i] != bindings[i] {
			t.Errorf("Expected binding (%v) at position %d, got %v\n%s", expectedBindings[i], i, bindings[i], expectedSQL)
		}
	}
}

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
		Select("name").From("table").Where(Eq("id", Param{"blah"})),
	)
}
