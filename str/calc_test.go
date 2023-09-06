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

func TestCalc_Count(t *testing.T) {
	var ipt, ref, rsl string

	// case
	ipt, ref = "5**3", "125"
	calc := NewCalc(ipt)
	calc.Count()
	rsl = calc.String()
	if rsl != ref {
		t.Errorf("input -> rsl ≠ ref: %s -> %s ≠ %s", ipt, rsl, ref)
	}

	// case
	ipt, ref = "9^3", "729"
	calc.Count(ipt)
	rsl = calc.String()
	if rsl != ref {
		t.Errorf("input -> rsl ≠ ref: %s -> %s ≠ %s", ipt, rsl, ref)
	}

	// case
	ipt, ref = "(5*20+6*35)*6-11*3", "1827"
	calc.Count(ipt)
	rsl = calc.String()
	if rsl != ref {
		t.Errorf("input -> rsl ≠ ref: %s -> %s ≠ %s", ipt, rsl, ref)
	}

	// case
	ipt, ref = "538,000/15", "35866.6666667"
	calc.Count(ipt)
	rsl = calc.String()
	if rsl != ref {
		t.Errorf("input -> rsl ≠ ref: %s -> %s ≠ %s", ipt, rsl, ref)
	}

	// case
	ipt, ref = "356%30", "26"
	calc.Count(ipt)
	rsl = calc.String()
	if rsl != ref {
		t.Errorf("input -> rsl ≠ ref: %s -> %s ≠ %s", ipt, rsl, ref)
	}

	// case
	ipt, ref = "sqrt(3**2+4**2)", "5"
	calc.Count(ipt)
	rsl = calc.String()
	if rsl != ref {
		t.Errorf("input -> rsl ≠ ref: %s -> %s ≠ %s", ipt, rsl, ref)
	}

	// case
	ipt, ref = "sqrt(sqrt(625)) + sqrt(3**2+4**2)**2", "30"
	calc.Count(ipt)
	rsl = calc.String()
	if rsl != ref {
		t.Errorf("input -> rsl ≠ ref: %s -> %s ≠ %s", ipt, rsl, ref)
	}
}
