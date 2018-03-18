package qg

import "testing"

func TestJoinBasic(t *testing.T) {

	assertSQL(t,
		"select id from table a inner join table b on a.id = b.id",
		[]interface{}{},
		Select("id").From("table a").InnerJoin("table b", Eq("a.id", "b.id")),
	)

	assertSQL(t,
		"select id from table inner join table using (id)",
		[]interface{}{},
		Select("id").From("table").InnerJoin("table", Using("id")),
	)
}

func TestJoinDeep(t *testing.T) {

	assertSQL(t,
		"select id from table inner join table using (id) left join table using (id) outer join table using (id) right join table using (id)",
		[]interface{}{},
		Select("id").From("table").
			InnerJoin("table", Using("id")).
			LeftJoin("table", Using("id")).
			OuterJoin("table", Using("id")).
			RightJoin("table", Using("id")))
}

func TestJoinNested(t *testing.T) {

	assertSQL(t,
		"select id from table inner join (select id from table) as table using (id)",
		[]interface{}{},
		Select("id").From("table").InnerJoin(Alias(Select("id").From("table"), "table"), Using("id")))
}
