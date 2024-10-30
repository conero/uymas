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
