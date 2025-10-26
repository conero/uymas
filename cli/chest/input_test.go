package chest

import "testing"

func TestLineAsArgs(t *testing.T) {
	line := " git  clone  https://gitee.com/conero/uymas.git "
	args := LineAsArgs(line)
	if len(args) != 3 {
		t.Errorf("参数解析错误，参数长度：%d", len(args))
	} else {
		t.Logf("参数：%#v", args)
	}

	// case
	line = `echo "use go to code"`
	args = LineAsArgs(line)
	if len(args) != 2 {
		t.Errorf("参数解析错误，参数长度：%d，参数：%#v", len(args), args)
	} else {
		t.Logf("参数：%#v", args)
	}
}
