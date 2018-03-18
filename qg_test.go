package qg

import "testing"

func assertSQL(t *testing.T, expectedSQL string, expectedBindings []interface{}, query SQLable) {
	sql, bindings := query.ToSQL()
	if expectedSQL != sql {
		t.Errorf("Expected '%s', got '%s'", expectedSQL, sql)
	}

	if len(expectedBindings) != len(bindings) {
		t.Errorf("Expected %d bindings, got %d\n%#v\n%#v", len(expectedBindings), len(bindings), expectedBindings, bindings)
		return
	}

	for i := range expectedBindings {
		if expectedBindings[i] != bindings[i] {
			t.Errorf("Expected binding (%v) at position %d, got %v\n%s", expectedBindings[i], i, bindings[i], expectedSQL)
		}
	}
}
