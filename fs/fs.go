//Package fs support facilitate for handler file and direction.
package fs

import (
	"github.com/conero/uymas/number"
	"runtime"
)

// GetMemStatsRange get memory stats range callback
func GetMemStatsRange() func() (runtime.MemStats, runtime.MemStats) {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	return func() (runtime.MemStats, runtime.MemStats) {
		var mem2 runtime.MemStats
		runtime.ReadMemStats(&mem2)
		return mem, mem2
	}
}

// MemUsage memory usage
type MemUsage struct {
}

// GetSysMemSub get system memory range sub bytes
func (m *MemUsage) GetSysMemSub() func() number.BitSize {
	mRangeCall := GetMemStatsRange()
	return func() number.BitSize {
		m1, m2 := mRangeCall()
		return number.Byte * number.BitSize(m2.Sys-m1.Sys)
	}
}
