package number

import "testing"

func TestBitSize_Format(t *testing.T) {
	var bit BitSize = 8 * 1234
	t.Log(bit)
}
