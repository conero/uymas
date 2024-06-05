package pinyin

import (
	"gitee.com/conero/uymas/bin/butil"
	"testing"
)

func TestList_Polyphony(t *testing.T) {
	fl := butil.Pwd("material/mt_pinyin.txt")
	py := NewPinyin(fl)
	if py.IsEmpty() {
		t.Fatalf("未加载拼音资源，请在根目录执行测试。\n 资源目录： %s", fl)
	}

	// 重庆、重庆长安
	word := `重庆长安`
	ls := py.SearchByGroup(word)
	pyList := ls.Polyphony(0)
	if len(pyList) < 2 {
		t.Errorf("多音字组合错误")
	}

	//ply := ls.Polyphony()

	//t.Logf("List: %v", ls.Alpha())
	//t.Logf("Polyphony: %#v", ply)

	//case
	word = `友谊长存，铁拳重击`
	ls = py.SearchByGroup(word)
	if len(pyList) < 2 {
		t.Errorf("多音字组合错误")
	}
}
