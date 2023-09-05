package str

import "testing"

func TestFloatSimple(t *testing.T) {
	var ipt, ref, rsl string

	// case
	ipt, ref = "3.14", "3.14"
	rsl = FloatSimple(ipt)
	if rsl != ref {
		t.Errorf("input -> rsl ≠ ref: %s -> %s ≠ %s", ipt, rsl, ref)
	}

	// case
	ipt, ref = "0.100000", "0.1"
	rsl = FloatSimple(ipt)
	if rsl != ref {
		t.Errorf("input -> rsl ≠ ref: %s -> %s ≠ %s", ipt, rsl, ref)
	}

	// case
	ipt, ref = "0.2309050", "0.230905"
	rsl = FloatSimple(ipt)
	if rsl != ref {
		t.Errorf("input -> rsl ≠ ref: %s -> %s ≠ %s", ipt, rsl, ref)
	}

	// case
	ipt, ref = "99", "99"
	rsl = FloatSimple(ipt)
	if rsl != ref {
		t.Errorf("input -> rsl ≠ ref: %s -> %s ≠ %s", ipt, rsl, ref)
	}

	// case
	ipt, ref = "100.01", "100.01"
	rsl = FloatSimple(ipt)
	if rsl != ref {
		t.Errorf("input -> rsl ≠ ref: %s -> %s ≠ %s", ipt, rsl, ref)
	}

	// case
	ipt, ref = "3.1415926540000000000000", "3.141592654"
	rsl = FloatSimple(ipt)
	if rsl != ref {
		t.Errorf("input -> rsl ≠ ref: %s -> %s ≠ %s", ipt, rsl, ref)
	}
}
