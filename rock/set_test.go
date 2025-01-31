package rock

import (
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

func TestListEq(t *testing.T) {
	intArr1 := []int{1, 9, 49, 1001, 24, 903}
	intArr2 := []int{903, 24, 49}

	// case
	if ListEq(intArr1, intArr2) {
		t.Errorf("%#v = %#v，次判别错误", intArr1, intArr2)
	}

	// case
	intArr2 = []int{903, 24, 1001, 49, 9, 1}
	if !ListEq(intArr1, intArr2) {
		t.Errorf("%#v ≠ %#v，次判别错误", intArr1, intArr2)
	}

	intArr1, intArr2 = nil, nil
	if !ListEq(intArr1, intArr2) {
		t.Errorf("%#v ≠ %#v，次判别错误", intArr1, intArr2)
	}

}

func TestListSubset(t *testing.T) {
	intArr1 := []int{1, 9, 49, 1001, 24, 903}
	intArr2 := []int{903, 24, 49}

	// case
	if !ListSubset(intArr1, intArr2) {
		t.Errorf("%#v 应该为 %#v 的子数组，次判别错误", intArr1, intArr2)
	}

	strArr1 := `I am Jc, Coder.`
	strArr2 := "Coder."

	// case
	if !ListSubset([]byte(strArr1), []byte(strArr2)) {
		t.Errorf("%#v 应该为 %#v 的子数组，次判别错误", intArr1, intArr2)
	}
}
