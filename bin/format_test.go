package bin

import "testing"

func TestFormatKv(t *testing.T) {
	tdd := map[string]interface{}{
		"author":                   "Joshua Conero",
		"email":                    "conero@163",
		"a":                        "TestFormatKv for beautify string.",
		"canBeALongStringTestAlso": 2,
	}
	t.Log("\n" + FormatKv(tdd))
	t.Log("\n" + FormatKv(tdd, ". "))
	t.Log("\n" + FormatKv(tdd, ". ", "*"))

	// support more type.
	tdd2 := map[interface{}]interface{}{
		"author":                   "Joshua Conero",
		210609:                     "conero@163",
		true:                       "TestFormatKv for beautify string.",
		"canBeALongStringTestAlso": 2,
	}
	t.Log("\n" + FormatKv(tdd2))
	t.Log("\n" + FormatKv(tdd2, ". "))
	t.Log("\n" + FormatKv(tdd2, ". ", "*"))
}
