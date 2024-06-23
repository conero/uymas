package str

import (
	"testing"
	"time"
)

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

	// case
	ipt = "1949/10/01 09:30"
	ref = "2006/01/02 15:04"
	actl, err = TimeParseLayout(ipt)
	if err != nil {
		t.Error(err)
	} else if actl != ref {
		t.Fatalf("值 [%v], 转换后 [%v]; 与参考值 [%v] 不符合！", ipt, actl, ref)
	}

	// case
	ipt = "19491001 093000"
	ref = "20060102 150405"
	actl, err = TimeParseLayout(ipt)
	if err != nil {
		t.Error(err)
	} else if actl != ref {
		t.Fatalf("值 [%v], 转换后 [%v]; 与参考值 [%v] 不符合！", ipt, actl, ref)
	}

	// case
	ipt = "201109 0830"
	ref = "200601 1504"
	actl, err = TimeParseLayout(ipt)
	if err != nil {
		t.Error(err)
	} else if actl != ref {
		t.Fatalf("值 [%v], 转换后 [%v]; 与参考值 [%v] 不符合！", ipt, actl, ref)
	}

	// case
	ipt = "201109 08"
	ref = "200601 15"
	actl, err = TimeParseLayout(ipt)
	if err != nil {
		t.Error(err)
	} else if actl != ref {
		t.Fatalf("值 [%v], 转换后 [%v]; 与参考值 [%v] 不符合！", ipt, actl, ref)
	}

	// case
	ipt = "2023年08月25日"
	ref = "2006年01月02日"
	actl, err = TimeParseLayout(ipt)
	if err != nil {
		t.Error(err)
	} else if actl != ref {
		t.Fatalf("值 [%v], 转换后 [%v]; 与参考值 [%v] 不符合！", ipt, actl, ref)
	}

	// case
	ipt = "2023年08月25"
	ref = "2006年01月02"
	actl, err = TimeParseLayout(ipt)
	if err != nil {
		t.Error(err)
	} else if actl != ref {
		t.Fatalf("值 [%v], 转换后 [%v]; 与参考值 [%v] 不符合！", ipt, actl, ref)
	}

	// case
	ipt = "2023年08月"
	ref = "2006年01月"
	actl, err = TimeParseLayout(ipt)
	if err != nil {
		t.Error(err)
	} else if actl != ref {
		t.Fatalf("值 [%v], 转换后 [%v]; 与参考值 [%v] 不符合！", ipt, actl, ref)
	}

	// case
	ipt = "2023年"
	ref = "2006年"
	actl, err = TimeParseLayout(ipt)
	if err != nil {
		t.Error(err)
	} else if actl != ref {
		t.Fatalf("值 [%v], 转换后 [%v]; 与参考值 [%v] 不符合！", ipt, actl, ref)
	}
}

func TestParseDuration(t *testing.T) {
	duraStr := "10天"
	dura, err := ParseDuration(duraStr)
	ref := 10 * 24 * time.Hour

	toDiffFn := func() {
		if err != nil {
			t.Errorf("%s 解析错误，%v", duraStr, err)
		} else if ref != dura {
			t.Errorf("%s(%s) ≠ %s", duraStr, dura, ref)
		}
	}
	toDiffFn()

	// case
	duraStr = "time spend: 6hours, 10.3 minute"
	dura, err = ParseDuration(duraStr)
	ref = 6*time.Hour + 10*time.Minute + (0.3*60)*time.Second
	toDiffFn()
}
