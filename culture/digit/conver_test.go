package digit

import "testing"

func TestCover_ToChnUpper(t *testing.T) {
	var test Cover = 9000
	test.ToChnUpper()

	test = 2391792.0872
	test.ToChnUpper()
}
