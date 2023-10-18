package number

import (
	"fmt"
	"testing"
)

func ExampleUnit_Unit() {
	n := Unit(5329742)
	fmt.Printf("5329742: %v\n", n)

	// Output:
	// 5329742: 5.3297 M
}

func TestFactorial(t *testing.T) {
	var ipt, ref, rsl uint64

	// case
	ipt, ref = 0, 1
	rsl = Factorial(ipt)
	if rsl != ref {
		t.Errorf("input -> rsl ≠ ref: %d! -> %d ≠ %d", ipt, rsl, ref)
	}

	// case
	ipt, ref = 1, 1
	rsl = Factorial(ipt)
	if rsl != ref {
		t.Errorf("input -> rsl ≠ ref: %d! -> %d ≠ %d", ipt, rsl, ref)
	}

	// case
	ipt, ref = 4, 24
	rsl = Factorial(ipt)
	if rsl != ref {
		t.Errorf("input -> rsl ≠ ref: %d! -> %d ≠ %d", ipt, rsl, ref)
	}

	// case
	ipt, ref = 9, 362880
	rsl = Factorial(ipt)
	if rsl != ref {
		t.Errorf("input -> rsl ≠ ref: %d! -> %d ≠ %d", ipt, rsl, ref)
	}
}
