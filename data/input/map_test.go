package input

import "testing"

func TestSplitMapBasic(t *testing.T) {
	vList := []string{"name.Joshua Conero", "gender.M"}
	m := SplitMapBasic(vList)
	t.Log(m)
}
