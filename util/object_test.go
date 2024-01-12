package util

import (
	"fmt"
	"math"
	"testing"
	"time"
)

type testObjectTopMac struct {
	ID         int       `json:"id"`
	Mid        int       `json:"mid"`
	Name       string    `json:"name"`
	Label      string    `json:"label"`
	IsAdd      int       `json:"is_add"`
	IsUpd      int       `json:"is_upd"`
	IsList     int       `json:"is_list"`
	IsSearch   int       `json:"is_search"`
	EditType   string    `json:"edit_type"`
	EditOption string    `json:"edit_option"`
	IsRequire  int       `json:"is_require"`
	LsColWidth int       `json:"ls_col_width"`
	VarDefault string    `json:"var_default"`
	ColOrder   int       `json:"col_order"`
	CreateTime time.Time `json:"create_time"`
}

// object 测试大工具
type testObjectMax struct {
	ID         int       `json:"id"`
	Module     string    `json:"module"`
	Label      string    `json:"label"`
	IsOpen     int       `json:"is_open"`
	Author     string    `json:"author"`
	Version    string    `json:"version"`
	MainTable  string    `json:"main_table"`
	IsDel      int       `json:"is_del"`
	IsUpd      int       `json:"is_upd"`
	IsAdd      int       `json:"is_add"`
	OperMain   string    `json:"oper_main"`
	OperAdd    string    `json:"oper_add"`
	OperEdit   string    `json:"oper_edit"`
	OperDel    string    `json:"oper_del"`
	LsWidth    int       `json:"ls_width"`
	PageOption string    `json:"page_option"`
	RouterName string    `json:"router_name"`
	CreateTime time.Time `json:"create_time"`
	SubObject  testObjectTopMac
}

var tom = testObjectMax{
	ID:         1024,
	Module:     "Uymas",
	Label:      "Joshua Conero",
	IsOpen:     1,
	Author:     "古丞秋",
	Version:    "v1.1.2",
	MainTable:  "object_max",
	IsDel:      0,
	IsUpd:      1,
	IsAdd:      0,
	OperAdd:    "add",
	OperDel:    "del",
	PageOption: "遥望中原，荒烟外、许多城郭。想当年、花遮柳护，凤楼龙阁。万岁山前珠翠绕，蓬壶殿里笙歌作。到而今、铁骑满郊畿，风尘恶。兵安在，膏锋锷；民安在，填沟壑。叹江山如故，千村寥落。何日请缨提锐旅，一鞭直渡清河洛？却归来、再续汉阳游，骑黄鹤。",
	CreateTime: time.Now(),
	SubObject: testObjectTopMac{
		ID:         1024,
		Mid:        1024,
		Name:       "name",
		Label:      "姓名",
		IsAdd:      0,
		IsUpd:      1,
		IsList:     0,
		IsSearch:   2,
		EditType:   "C",
		EditOption: "君不见黄河之水天上来，奔流到海不复回。君不见高堂明镜悲白发，朝如青丝暮成雪。人生得意须尽欢，莫使金樽空对月。天生我材必有用，千金散尽还复来。烹羊宰牛且为乐，会须一饮三百杯。岑夫子，丹丘生，将进酒，君莫停。与君歌一曲，请君为我侧耳听。钟鼓馔玉不足贵，但愿长醉不愿醒。古来圣贤皆寂寞，惟有饮者留其名。陈王昔时宴平乐，斗酒十千恣欢谑。主人何为言少钱，径须沽取对君酌。五花马，千金裘，呼儿将出换美酒，与尔同销万古愁。",
		IsRequire:  4,
		LsColWidth: 5,
		VarDefault: "txt",
		ColOrder:   7,
		CreateTime: time.Now(),
	},
}

func TestObject_Assign(t *testing.T) {
	o := Object{}
	//simple struct
	type ta struct {
		Name      string
		Age       int
		IsMale    bool
		Score     float64
		ScoreList []float64
	}

	//complex struct
	type tc struct {
		Ta        ta
		Name      string
		IsGroup   bool
		VtypeMark int32
	}

	taDefualt := ta{
		Name: "Emma",
		Age:  18,
	}

	//case 1: sample
	ta1 := taDefualt
	srs1 := ta{
		Name:      "Charlie",
		IsMale:    true,
		Score:     0.1244,
		ScoreList: []float64{0.124, 0.124, 0.881},
	}
	o.Assign(&ta1, srs1)

	t.Logf("taDefualt: %#v , \r\n source: %#v, \n\r targat: %#v\r\n", taDefualt, srs1, ta1)

	//case 2: different struct
	var tc1 tc
	o.Assign(&tc1, srs1)
	t.Logf("zore tc assian after: %#v , \r\n use by last source", tc1)

	o.Assign(&tc1.Ta, srs1)
	t.Logf("tc target chage use field: tc.ta: %#v , \r\n use by last source", tc1)

	srs2 := tc{
		Ta: ta{
			Age:  25,
			Name: "Child Struct",
		},
		Name:      "Nesting",
		VtypeMark: 12,
	}
	o.Assign(&tc1, srs2)
	t.Logf("tc target chage use field: tc.ta: %#v , \r\n use by last source", tc1)
}

func TestObject_Assign2(t *testing.T) {
	type subConfig struct {
		Level   string
		Data    map[string]any
		Score   float64
		Charset string
		IsTcp   bool
	}
	type config struct {
		Port      int
		Type      string
		Username  string
		Pswd      string
		DevAble   bool
		SubConfig subConfig
	}

	//var sbCf = subConfig{
	//	Score: 1.4,
	//	Data:  map[string]any{"name": "joshua"},
	//}
	//var testSbCf = func(sc *subConfig) {
	//	sc.Score = 5.2
	//	sc.Data = map[string]any{"age": 31}
	//}
	//testSbCf(&sbCf)
	//fmt.Printf("sbCf: %#v\n", sbCf)

	// case
	var defCfg = config{
		Port:    3308,
		Type:    "mysql",
		DevAble: true,
		// @bug 二级存在问题
		// @todo
		SubConfig: subConfig{
			Data: map[string]any{
				"weight": 45.87,
			},
			Score: -0.19,
		},
	}
	var vCfg = config{
		Username: "sys-mng",
		Pswd:     "3gtwfrb6i.k-1/z*9'hd4x8e2p",
		SubConfig: subConfig{
			IsTcp: true,
		},
	}

	var obj Object
	obj.Assign(&vCfg, defCfg)
	if vCfg.Port != defCfg.Port {
		t.Errorf("Assign 数据合并无效，%#v", vCfg)
	}
	t.Logf("vConf: %#v", vCfg)

}

func TestObject_AssignMap(t *testing.T) {
	obj := Object{}
	type dog struct {
		Weight   float64
		Describe string
		Age      int
		IsMale   bool
	}

	var tgt = map[string]any{}
	srcMap := map[string]any{
		"name":  "Joshua Conero",
		"age":   18,
		"score": 89.3,
	}

	obj.AssignMap(tgt, srcMap)
	t.Logf("%#v", tgt)

	//case: map->map
	srcMap = map[string]any{
		"score": 93.14,
		"class": "A+",
	}
	obj.AssignMap(tgt, srcMap)
	t.Logf("%#v", tgt)

	//case: struct->map
	dg := dog{
		Weight:   14.23,
		Describe: "Chinese Cat.",
		Age:      2,
		IsMale:   true,
	}
	tgt = map[string]any{}
	obj.AssignMap(tgt, dg)
	t.Logf("%#v", tgt)

}

func TestStructToMap(t *testing.T) {
	type Ty struct {
		Name          string
		Age           int
		HeightWidthLv float64
		EmptyInt      int
		EmptyString   string
	}

	tt := Ty{}
	t.Logf("%v", StructToMap(tt))

	tt.Name = "Joshua Conero"
	tt.HeightWidthLv = math.Pi
	tt.Age = 58
	t.Logf("%v", StructToMap(tt))
	t.Logf("StructToMapLStyle: %#v", StructToMapLStyle(tt))
	t.Logf("StructToMapLStyleIgnoreEmpty: %#v", ToMapLStyleIgnoreEmpty(tt))

	// reflect.Ptr
	ty := &Ty{}
	ty = &tt
	t.Logf("Ty -> %#v", ty)
	t.Logf("%v", StructToMap(ty))

}

func BenchmarkStructToMapViaJson(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StructToMapViaJson(tom)
	}
}

func BenchmarkStructToMapLStyle(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StructToMapLStyle(tom)
	}
}

func TestStructToMapLStyle(t *testing.T) {
	t.Logf("StructToMapLStyle => %#v", StructToMapLStyle(tom))
	t.Logf("StructToMapLStyle => %#v", StructToMapLStyle(tom, "CreateTime", "ID", "SubObject"))
}

func TestObject_Keys(t *testing.T) {
	// map-case1
	var v any
	v = map[string]string{"name": "Joshua", "age": "23", "country": "cn"}

	var obj Object
	var rf = []string{"name", "age", "country"}

	keys := obj.Keys(v)
	vmStr := fmt.Sprintf("%#v", keys)
	rfStr := fmt.Sprintf("%#v", rf)
	if vmStr != rfStr {
		t.Errorf("Fail by map[string]string => %v != %v", vmStr, rfStr)
	}

	// map-case2
	// delete for map
	delete(v.(map[string]string), "country")
	keys = obj.Keys(v)
	vmStr = fmt.Sprintf("%#v", keys)

	rf = []string{"name", "age"}
	rfStr = fmt.Sprintf("%#v", rf)

	if vmStr != rfStr {
		t.Errorf("Fail by map[string]string => %v != %v", vmStr, rfStr)
	}

	// struct-case2
	type Ta struct {
		Name    string
		Age     int
		Country string
	}

	var ta Ta
	keys = obj.Keys(ta)
	vmStr = fmt.Sprintf("%#v", keys)

	rf = []string{"Name", "Age", "Country"}
	rfStr = fmt.Sprintf("%#v", rf)
	if vmStr != rfStr {
		t.Errorf("Fail by map[string]string => %v != %v", vmStr, rfStr)
	}

	// struct-case3
	var tPtr = &Ta{}
	keys = obj.Keys(tPtr)
	vmStr = fmt.Sprintf("%#v", keys)

	rf = []string{"Name", "Age", "Country"}
	rfStr = fmt.Sprintf("%#v", rf)
	if vmStr != rfStr {
		t.Errorf("Fail by map[string]string => %v != %v", vmStr, rfStr)
	}
}

func TestObject_AssignCovert(t *testing.T) {
	var obj Object
	var tgt, src any

	// case
	tgt = "Joshua"
	src = "Joshua Doeeking Conero"
	obj.AssignCovert(&tgt, src)
	ref := fmt.Sprintf("%T/%#v", src, src)
	rslt := fmt.Sprintf("%T/%#v", tgt, tgt)
	if ref != rslt {
		t.Errorf("至覆盖错误.\n  IN=>  %s\n  OUT=> %s\n", rslt, ref)
	}

	// case
	tgt = "Joshua"
	src = ""

	ref = fmt.Sprintf("%T/%#v", tgt, tgt)
	obj.AssignCovert(&tgt, src)
	rslt = fmt.Sprintf("%T/%#v", tgt, tgt)
	if ref != rslt {
		t.Errorf("至覆盖错误.\n  IN=>  %s\n  OUT=> %s\n", rslt, ref)
	}

	// case
	tgt = int64(64)
	src = 202_403
	obj.AssignCovert(&tgt, src)
	ref = fmt.Sprintf("%T/%#v", tgt, src)
	rslt = fmt.Sprintf("%T/%#v", tgt, tgt)
	if ref != rslt {
		t.Errorf("至覆盖错误.\n  IN=>  %s\n  OUT=> %s\n", rslt, ref)
	}

	// case
	tgt = 3.141592654
	src = 11_000
	obj.AssignCovert(&tgt, src)
	ref = fmt.Sprintf("%T/%#v", tgt, src)
	rslt = fmt.Sprintf("%T/%#v", tgt, tgt)
	if ref != rslt {
		t.Errorf("至覆盖错误.\n  IN=>  %s\n  OUT=> %s\n", rslt, ref)
	}

	// case
	tgt = 1949
	src = "1949"
	obj.AssignCovert(&tgt, src)
	ref = fmt.Sprintf("%T/%#v", tgt, 1949)
	rslt = fmt.Sprintf("%T/%#v", tgt, tgt)
	if ref != rslt {
		t.Errorf("至覆盖错误.\n  IN=>  %s\n  OUT=> %s\n", rslt, ref)
	}

	// case
	tgt = false
	src = "1949"
	obj.AssignCovert(&tgt, src)
	ref = fmt.Sprintf("%T/%#v", true, true)
	rslt = fmt.Sprintf("%T/%#v", tgt, tgt)
	if ref != rslt {
		t.Errorf("至覆盖错误.\n  IN=>  %s\n  OUT=> %s\n", rslt, ref)
	}

	// case
	tgt = 1949
	src = "1949"
	obj.AssignCovert(&tgt, src)
	ref = fmt.Sprintf("%T/%#v", tgt, 1949)
	rslt = fmt.Sprintf("%T/%#v", tgt, tgt)
	if ref != rslt {
		t.Errorf("至覆盖错误.\n  IN=>  %s\n  OUT=> %s\n", rslt, ref)
	}

	// case
	tgt = 2024.0328
	src = "2024.0328"
	obj.AssignCovert(&tgt, src)
	ref = fmt.Sprintf("%T/%#v", tgt, 2024.0328)
	rslt = fmt.Sprintf("%T/%#v", tgt, tgt)
	if ref != rslt {
		t.Errorf("至覆盖错误.\n  IN=>  %s\n  OUT=> %s\n", rslt, ref)
	}
}

func TestMapToStructViaJson(t *testing.T) {
	type test1 struct {
		Age    int64  `json:"age"`
		Name   string `json:"name"`
		IsMale bool
		Kg     float32
	}

	t1 := test1{}
	err := MapToStructViaJson(map[string]string{
		"age":    "15",
		"name":   "Joshua Conero",
		"IsMale": "true",
		"Kg":     "23.189",
	}, &t1)
	if err != nil {
		t.Errorf("MapToStructViaJson 错误，%v", err)
	} else {
		t.Logf("test1: %#v\n", t1)
		if t1.Age != 15 {
			t.Errorf("MapToStructViaJson 赋值 string -> int64 不匹配")
		}
		if t1.Name != "Joshua Conero" {
			t.Errorf("MapToStructViaJson 赋值 string -> string 不匹配")
		}
		if t1.IsMale != true {
			t.Errorf("MapToStructViaJson 赋值 string -> true 不匹配")
		}
	}

}
