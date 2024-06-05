package util

import (
	"fmt"
	"strings"
	"testing"
)

func TestListIndex(t *testing.T) {
	var ss = []string{"I", "am", "Joshua", "Conero", "."}
	var idx, rfIdx int
	rfIdx = 3
	idx = ListIndex(ss, "Conero")
	if idx != rfIdx {
		t.Errorf("Search []string Index failure: %v != %v", idx, rfIdx)
	}

	//
	var its = []uint8{1, 9, 9, 2, 1, 9, 4, 9}
	idx = ListIndex(its, 1)
	rfIdx = 0
	if idx != rfIdx {
		t.Errorf("Search []uint8 Index failure: %v != %v", idx, rfIdx)
	}
}

func TestListNoRepeat(t *testing.T) {
	var list = []string{"a", "a", "b", "c", "c", "f"}
	rel := ListNoRepeat(list)
	ref := "abcf"
	toTestFn := func() {
		matchRef := strings.Join(rel, "")
		if ref != matchRef {
			t.Errorf("%v 去重失败，%s ≠ %s", list, matchRef, ref)
		}
	}

	// case
	toTestFn()
}

func TestMapKeys(t *testing.T) {
	var vmStr = map[string]string{
		"name": "Joshua conero",
		"age":  "24",
	}

	var keys = MapKeys(vmStr)
	if fmt.Sprintf("%#v", keys) != fmt.Sprintf("%#v", []string{"name", "age"}) {
		t.Errorf("Get map keys failue, like: %#v", keys)
	}
}
