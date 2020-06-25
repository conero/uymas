package fs

import "testing"

func TestByteSize_String(t *testing.T) {
	var km ByteSize = 144 * KB
	t.Log(km)

	var bytes ByteSize = 27_488_109
	t.Log(bytes)

	km = 54387431
	t.Log(km)

	km = 5438743157233
	t.Log(km)
}
