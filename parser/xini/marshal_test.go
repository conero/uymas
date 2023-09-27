package xini

import "testing"

func TestMarshal(t *testing.T) {
	var v any

	// case 1
	v = map[string]any{
		"name":      "Joshua Conero",
		"score":     98.5,
		"max":       5601,
		"rate":      0.64256,
		"level":     "B",
		"is_finish": false,
		"list":      []float32{71, 80, 64, 99, 21.0},
		"signature": `
		怒发冲冠，凭栏处、潇潇雨歇。抬望眼、仰天长啸，壮怀激烈。三十功名尘与土，八千里路云和月。莫等闲、白了少年头，空悲切。
		靖康耻，犹未雪。臣子恨，何时灭。驾长车，踏破贺兰山缺。壮志饥餐胡虏肉，笑谈渴饮匈奴血。待从头、收拾旧山河，朝天阙。
		`,
	}

	by, er := Marshal(v)
	if er != nil {
		t.Errorf("xini encode error, %s", er.Error())
	}
	t.Logf("--------xini encode-----\n%s", string(by))

	// case 2
	v = struct {
		Manufacturer string
		CarName      string
		Wight        int
		IsOnline     bool
	}{
		Manufacturer: "比亚迪",
		CarName:      "宋L 2023",
		Wight:        2158,
	}

	by, er = Marshal(v)
	if er != nil {
		t.Errorf("xini encode error, %s", er.Error())
	}
	t.Logf("--------xini encode-----\n%s", string(by))
}
