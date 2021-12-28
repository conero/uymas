package bin

import "testing"

func TestCmd2StringMap(t *testing.T) {
	var i, o, r string

	// case1 get-child
	i = "get-child"
	r = "GetChild"
	o = Cmd2StringMap(i)
	if r != o {
		t.Fatalf("%v -> %v != %v", i, o, r)
	}

	// case2 get_child_item_and_showMe
	i = "get_child_item_and_showMe"
	r = "GetChildItemAndShowMe"
	o = Cmd2StringMap(i)
	if r != o {
		t.Fatalf("%v -> %v != %v", i, o, r)
	}

	// case3 version
	i = "version"
	r = "Version"
	o = Cmd2StringMap(i)
	if r != o {
		t.Fatalf("%v -> %v != %v", i, o, r)
	}

	// case4 version
	i = "yes or      no"
	r = "YesOrNo"
	o = Cmd2StringMap(i)
	if r != o {
		t.Fatalf("%v -> %v != %v", i, o, r)
	}

	// case test
	i = "test"
	r = "Test"
	o = Cmd2StringMap(i)
	if r != o {
		t.Fatalf("%v -> %v != %v", i, o, r)
	}
}
