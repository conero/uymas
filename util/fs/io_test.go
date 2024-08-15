package fs

import "testing"

func TestStdPathName(t *testing.T) {
	var vpath, vout, vwant string
	vpath = "C:/Users/11066/AppData/Local/Temp/runtime//caches/path_in_all20191105140208.tar"
	vwant = "C:/Users/11066/AppData/Local/Temp/runtime/caches/path_in_all20191105140208.tar"
	vout = StdPathName(vpath)

	if vout != vwant {
		t.Fatalf("结果与预期一致，%v => %v", vout, vwant)
	}
}
