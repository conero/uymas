package cloud

import "testing"

func TestPortAddress(t *testing.T) {
	vIpt := "80"
	rel := PortAddress(vIpt)
	ref := ":80"

	testFn := func() {
		if rel != ref {
			t.Errorf("%s -> %s !≠ %s", vIpt, rel, rel)
		}
	}

	// case
	testFn()

	vIpt = "localhost:443"
	rel = PortAddress(vIpt)
	ref = "localhost:443"
}

func TestPortAsWeb(t *testing.T) {
	vIpt := "80"
	rel := PortAsWeb(vIpt)
	ref := "http://localhost"
	testFn := func() {
		if rel != ref {
			t.Errorf("%s -> %s !≠ %s", vIpt, rel, rel)
		}
	}
	testFn()

	// case
	vIpt = "443"
	rel = PortAsWeb(vIpt, true)
	ref = "https://localhost"
	testFn()

	// case
	vIpt = "443"
	rel = PortAsWeb([2]string{vIpt, "jcconero.cn"}, true)
	ref = "https://jcconero.cn"
	testFn()
}
