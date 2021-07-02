package number

import (
	"fmt"
	"math"
)

type Unit int64

// 1k = 1000 thousand/kilo
// 1M = 1000_000 million
// 1G = 1000_000_000 billion
const (
	A Unit = 1
	K      = 1000 * A
	W      = 10 * K
	M      = 1000 * K
	G      = 1000 * M
)

// Format get the format of byte size
func (b Unit) Format() (float64, string) {
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

func (b Unit) String() string {
	v, unit := b.Format()
	return fmt.Sprintf("%.4f %v", v, unit)
}

func (b Unit) Unit() float64 {
	return float64(b)
}

func (b Unit) K() float64 {
	return float64(b / K)
}

func (b Unit) W() float64 {
	return float64(b / W)
}

func (b Unit) M() float64 {
	return float64(b / M)
}

func (b Unit) G() float64 {
	return float64(b / G)
}
