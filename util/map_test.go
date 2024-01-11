package util

import (
	"fmt"
	"testing"
)

func TestMapAssign(t *testing.T) {
	var source map[string]any
	rel := MapAssign(source)
	if rel == nil {
		t.Error("MapAssign 应该不返回 nil")
	}

	// case
	source = map[string]any{
		"bValue": 100000,
	}
	rel = MapAssign(source, map[string]any{
		"bValue": 15,
	}, map[string]any{
		"aValue": true,
	}, map[string]any{
		"y": "Gz",
		"x": "China",
	})
	ref := map[string]any{
		"bValue": 15,
		"aValue": true,
		"y":      "Gz",
		"x":      "China",
	}
	if rel == nil {
		t.Error("MapAssign 应该不返回 nil")
	}
	if fmt.Sprintf("%#v", rel) != fmt.Sprintf("%#v", ref) {
		t.Errorf("map 合并错误， %#v ≠ %#v\n", rel, ref)
	}

	// case
	var srcStr map[string]string
	var refStr = map[string]string{
		"joshua":  "xxxx",
		"gender":  "male",
		"address": "B.C.M xxx C",
	}
	relStr := MapAssign(srcStr, map[string]string{"joshua": "xxxx"},
		map[string]string{"joshua": "A+"},
		map[string]string{"joshua": "xxxx"},
		map[string]string{"gender": "male"},
		map[string]string{"address": "B.C.M xxx C"},
	)
	if fmt.Sprintf("%#v", refStr) != fmt.Sprintf("%#v", relStr) {
		t.Errorf("map 合并错误，%#v ≠ %#v\n", relStr, refStr)
	}
}
