package number

import "testing"

func TestOne_String(t *testing.T) {
	var m = 10 * M
	t.Log(m)

	var w = 328 * W
	t.Log(w)

	var g = 58392 * G
	t.Log(g)

	var n One

	n = 1234
	t.Log(n)

	n = 5678
	t.Log(n)

	n = 5329742
	t.Log(n)
}
