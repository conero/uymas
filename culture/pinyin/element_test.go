package pinyin

import (
	"gitee.com/conero/uymas/bin/butil"
	"testing"
)

func TestList_Polyphony(t *testing.T) {
	word := `友谊长存，铁拳重击`
	fl := butil.Pwd("material/mt_pinyin.txt")
	py := NewPinyin(fl)
	if py.IsEmpty() {
		t.Errorf("未加载拼音资源，请在根目录执行测试。\n 资源目录： %s", fl)
	}
	ls := py.SearchByGroup(word)
	ls.Polyphony()
	//ply := ls.Polyphony()

	//t.Logf("List: %v", ls.Alpha())
	//t.Logf("Polyphony: %#v", ply)
}
