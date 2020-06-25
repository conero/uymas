package number

import (
	"fmt"
	"math"
)

type One int64

// 1k = 1000 thousand/kilo
// 1M = 1000_000 million
// 1G = 1000_000_000 billion
const (
	A One = 1
	K     = 1000 * A
	W     = 10 * K
	M     = 1000 * K
	G     = 1000 * M
)

//get the format of byte size
func (b One) Format() (float64, string) {
	if b == 0 {
		return 0, ""
	}
	var sizes = []string{"", "K", "M", "G"}
	var i = math.Floor(math.Log10(float64(b)) / math.Log10(1000))
	//the max data unit is to `G`
	var sizesLen = float64(len(sizes))
	if i > sizesLen {
		i = sizesLen - 1
	}
	return float64(b) / math.Pow(1000, i), sizes[int(i)]
}

func (b One) String() string {
	v, unit := b.Format()
	return fmt.Sprintf("%.4f %v", v, unit)
}
