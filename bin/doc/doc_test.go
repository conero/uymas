package doc

import "testing"

func TestParseKv(t *testing.T) {
	ln := ":lang = zh"
	refKey, refValue := "lang", "zh"
	rlKey, rlValue := ParseKv(ln)

	toMatch := func() {
		if refKey != rlKey {
			t.Errorf("Key 解析失败，%s ≠ %s", rlKey, refKey)
		}
		if refValue != rlValue {
			t.Errorf("Value 解析失败，%s ≠ %s", rlValue, refValue)
		}
	}
	toMatch()

	// case
	ln = " : lang-support = zh,en"
	refKey, refValue = "lang-support", "zh,en"
	rlKey, rlValue = ParseKv(ln)
	toMatch()

	// case
	ln = ":name = Joshua Conero"
	refKey, refValue = "name", "Joshua Conero"
	rlKey, rlValue = ParseKv(ln)
	toMatch()

}
