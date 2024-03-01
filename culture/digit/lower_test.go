package digit

import "testing"

func TestLowerIndex(t *testing.T) {
	idx, ref := 0, "零"
	rel := LowerIndex(idx)
	todoFn := func() {
		if rel != ref {
			t.Errorf("%d -> %s ≠ %s", idx, rel, ref)
		}
	}

	// case
	todoFn()

	// case
	idx, ref = 9, "九"
	rel = LowerIndex(idx)
	todoFn()

	// case
	idx, ref = 10, "十"
	rel = LowerIndex(idx)
	todoFn()

	// case
	idx, ref = 17, "十七"
	rel = LowerIndex(idx)
	todoFn()

	// case
	idx, ref = 50, "五十"
	rel = LowerIndex(idx)
	todoFn()

	// case
	idx, ref = 59, "五十九"
	rel = LowerIndex(idx)
	todoFn()

	// case
	idx, ref = 105, "一百零五"
	rel = LowerIndex(idx)
	todoFn()

	// case
	idx, ref = 436, "四百三十六"
	rel = LowerIndex(idx)
	todoFn()

	// case
	idx, ref = 1024, "一千零二十四"
	rel = LowerIndex(idx)
	todoFn()
}
