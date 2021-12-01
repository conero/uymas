package util

import (
	"math"
	"testing"
)

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

func TestObject_AssignMap(t *testing.T) {
	obj := Object{}
	type dog struct {
		Weight   float64
		Describe string
		Age      int
		IsMale   bool
	}

	var tgt = map[string]interface{}{}
	srcMap := map[string]interface{}{
		"name":  "Joshua Conero",
		"age":   18,
		"score": 89.3,
	}

	obj.AssignMap(tgt, srcMap)
	t.Logf("%#v", tgt)

	//case: map->map
	srcMap = map[string]interface{}{
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
	tgt = map[string]interface{}{}
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
