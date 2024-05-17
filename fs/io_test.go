package fs

import "testing"

func TestStdPathName(t *testing.T) {
	var vpath, vwant string
	vpath = "C:/Users/11066\\AppData\\Local/Temp/runtime//caches/path_in_all20191105140208.tar"
	vwant = "C:/Users/11066/AppData/Local/Temp/runtime/caches/path_in_all20191105140208.tar"
	toTestFn := func() {
		out := StdPathName(vpath)
		if out != vwant {
			t.Fatalf("结果与预期一致，%v => %v", out, vwant)
		}
	}
	toTestFn()

	// case
	vpath = "/root/jc//base.dddd////1223/2/333333333333333333/"
	vwant = "/root/jc/base.dddd/1223/2/333333333333333333/"
	toTestFn()

	// case
	vpath = "c:/joshua conero/x/\\\\\\\\\\\\/m.x"
	vwant = "c:/joshua conero/x/m.x"
	toTestFn()

	// case
	vpath = "c:\\joshua conero\\app\\\\\\\\install.js"
	vwant = "c:/joshua conero/app/install.js"
	toTestFn()

	// case
	vpath = "joshua conero\\common\\\\\\\\vue.js/"
	vwant = "joshua conero/common/vue.js/"
	toTestFn()
}
