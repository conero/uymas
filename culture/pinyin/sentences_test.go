package pinyin

import (
	"testing"
)

func TestZhSentences_Len(t *testing.T) {
	s := ZhSentences("中华人民共和国")
	refLen := 7

	rslt := s.Len()
	if refLen != rslt {
		t.Errorf("%v 字符串长度错误：%d ≠ %d", s, rslt, refLen)
	}
}
