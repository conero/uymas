package number

import "testing"

func TestBitSize_Format(t *testing.T) {
	var bit BitSize = 8 * 1234
	t.Log(bit)

	// Size-turn test, Link: https://www.bejson.com/convert/filesize/
	// 4362076160->545259520->532480->520->0.5078125->0.00049591064453125
	m520 := 520 * MiB
	var test int64 = 4362076160
	// Bit
	if testy := BitSize(test); m520 != testy {
		t.Errorf("Bit/当前值与实际不同：%.f (%v) != %.f", m520.Bit(), m520, testy.Bit())
	}
	// Byte
	test = 545259520
	if testy := BitSize(test) * Byte; m520 != testy {
		t.Errorf("Byte/当前值与实际不同：%.f (%v) != %.f", m520.Byte(), m520, testy.Byte())
	}
	// KiB
	test = 532480
	if testy := BitSize(test) * KiB; m520 != testy {
		t.Errorf("KB/当前值与实际不同：%.f (%v) != %.f", m520.KiB(), m520, testy.KiB())
	}
	// MiB
	test = 520
	if testy := BitSize(test) * MiB; m520 != testy {
		t.Errorf("MB/当前值与实际不同：%.f (%v) != %.f", m520.MiB(), m520, testy.MiB())
	}
	// GiB
	f64 := 0.5078125
	if testy := BitSize(f64 * float64(GiB)); m520 != testy {
		t.Errorf("GB/当前值与实际不同：%f (%v) != %f", m520.GiB(), m520, testy.GiB())
	}
}
