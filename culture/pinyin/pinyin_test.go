package pinyin

import "testing"

func TestPyinNumber(t *testing.T) {
	word, ref := `guó`, `guo2`
	rslt := PyinNumber(word)

	runCaseFn := func() {
		if ref != rslt {
			t.Errorf("f(%s) -> %s ≠ %s", word, rslt, ref)
		}
	}
	// case
	runCaseFn()

	// case
	word, ref = `yòng`, `yong4`
	rslt = PyinNumber(word)
	runCaseFn()

	// case
	word, ref = `quán,shēn,ér,tuì`, `quan2,shen1,er2,tui4`
	rslt = PyinNumber(word)
	runCaseFn()

}
