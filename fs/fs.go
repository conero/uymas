//Package fs support facilitate for handler file and direction.
package fs

import (
	"fmt"
	"github.com/conero/uymas/number"
	"math"
	"runtime"
)

//ByteSize is a example like `time.Duration`
// Deprecated: Use number.ByteSize replace.
type ByteSize int64

//1 Bytes = 8Bit					byte
//1 B 	  = 8Bit					byte			字节
//1 KB    = 1024 Bytes			    Kilobyte		千字节
//1 MB	  = 1024 KB				    Megabyte		百万字节		   兆
//1 GB    = 1024 MB				    Gigabyte		千兆			   吉
//1 TB	  = 1024 GB				    Terabyte		万亿字节		   太
//1 PB	  = 1024 TB				    Petabyte		千万亿字节	   拍
//1 EB	  = 1024 PB				    Exabyte			百亿亿字节	   艾
//1 ZB	  = 1024 EB				    Zettabyte		十万亿亿字节	   泽
//1 YB	  = 1024 ZB				    Yottabyte		一亿亿亿字节	   尧
//1 BB	  = 1024 YB				    Brontobyte
//1 NB	  = 1024 BB				    NonaByte
//1 DB	  = 1024 NB				    DoggaByte
const (
	//Bit   BitSize = 1
	Bytes ByteSize = 1
	KB             = 1024 * Bytes
	MB             = 1024 * KB
	GB             = 1024 * MB
	TB             = 1024 * GB
	PB             = 1024 * TB
	EB             = 1024 * PB
	//ZB,YB is too big for int64
	//ZB             = 1024 * EB
	//YB             = 1024 * ZB
)

//get the format of byte size
func (b ByteSize) Format() (float64, string) {
	if b == 0 {
		return 0, "Bytes"
	}
	var sizes = []string{"Bytes", "KB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB"}
	var i = math.Floor(math.Log(float64(b)) / math.Log(1024))
	//the max data unit is to `G`
	var sizesLen = float64(len(sizes))
	if i > sizesLen {
		i = sizesLen - 1
	}
	return float64(b) / math.Pow(1024, i), sizes[int(i)]
}

func (b ByteSize) String() string {
	v, unit := b.Format()
	return fmt.Sprintf("%.4f %v", v, unit)
}

// get memory stats range callback
func GetMemStatsRange() func() (runtime.MemStats, runtime.MemStats) {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	return func() (runtime.MemStats, runtime.MemStats) {
		var mem2 runtime.MemStats
		runtime.ReadMemStats(&mem2)
		return mem, mem2
	}
}

//memory usage
type MemUsage struct {
}

//get system memory range sub bytes
func (m *MemUsage) GetSysMemSub() func() number.BitSize {
	mRangeCall := GetMemStatsRange()
	return func() number.BitSize {
		m1, m2 := mRangeCall()
		return number.Byte * number.BitSize(m2.Sys-m1.Sys)
	}
}
