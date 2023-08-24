package str

import "testing"

func TestTimeParseLayout(t *testing.T) {
	var ipt, ref, actl string
	var err error

	// case
	ipt = "2023-08-24"
	ref = "2006-01-02"
	actl, err = TimeParseLayout(ipt)
	if err != nil {
		t.Error(err)
	} else if actl != ref {
		t.Fatalf("值 [%v], 转换后 [%v]; 与参考值 [%v] 不符合！", ipt, actl, ref)
	}

	// case
	ipt = "2023-6-23"
	ref = "2006-1-02"
	actl, err = TimeParseLayout(ipt)
	if err != nil {
		t.Error(err)
	} else if actl != ref {
		t.Fatalf("值 [%v], 转换后 [%v]; 与参考值 [%v] 不符合！", ipt, actl, ref)
	}

	// case
	ipt = "2023-06"
	ref = "2006-01"
	actl, err = TimeParseLayout(ipt)
	if err != nil {
		t.Error(err)
	} else if actl != ref {
		t.Fatalf("值 [%v], 转换后 [%v]; 与参考值 [%v] 不符合！", ipt, actl, ref)
	}

	// case
	ipt = "2023-6-08 18:9:7"
	ref = "2006-1-02 15:4:5"
	actl, err = TimeParseLayout(ipt)
	if err != nil {
		t.Error(err)
	} else if actl != ref {
		t.Fatalf("值 [%v], 转换后 [%v]; 与参考值 [%v] 不符合！", ipt, actl, ref)
	}

	// case
	ipt = "08-09-08 20:20"
	ref = "06-01-02 15:04"
	actl, err = TimeParseLayout(ipt)
	if err != nil {
		t.Error(err)
	} else if actl != ref {
		t.Fatalf("值 [%v], 转换后 [%v]; 与参考值 [%v] 不符合！", ipt, actl, ref)
	}
}
