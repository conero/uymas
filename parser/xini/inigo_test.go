package xini

import (
	"fmt"
	"testing"
)

// @Date：   2018/8/19 0019 15:03
// @Author:  Joshua Conero
// @Name:    ini test
func TestNewParser(t *testing.T) {
	p := NewParser()

	// int
	p.Set("test", 5)
	//fmt.Println(p.GetData())
	has, value := p.Get("test")
	if !has || value.(int) != 5 {
		t.Fatal("[\"test\"=5] 设置值无效")
	}

	// bool
	p.Set("bool", true)
	//fmt.Println(p.GetData())
	has, value = p.Get("bool")
	if !has || value.(bool) != true {
		t.Fatal("[\"bool\"=true] 设置值无效")
	}

}

func TestNewParserRong(t *testing.T) {
	rong := NewParser(nil, SupportNameRong)
	fmt.Println(rong)
	fmt.Println(rong.Driver())

	if rong.Driver() != SupportNameRong {
		t.Fatal("Driver 默认生成无效！")
	}
}

func TestNewParserIni(t *testing.T) {
	ini := NewParser(nil, SupportNameIni)
	ini.Set("test", 8).
		Set("name", "Full")
	fmt.Println(ini)
	if ini.Driver() != SupportNameIni {
		t.Fatal("Driver 默认生成无效！")
	}
}

func TestParseValue(t *testing.T) {
	// bool
	if !ParseValue("true").(bool) ||
		ParseValue("false").(bool) ||
		!ParseValue("TRUE").(bool) {
		t.Error("bool 类型解析失败")
	}

	// int64
	if ParseValue("100").(int64) != 100 {
		t.Error("整形解析错误： 100")
	}
	if ParseValue("0").(int64) != 0 {
		t.Error("整形解析错误： 0")
	}
	if ParseValue("-9672").(int64) != -9672 {
		t.Error("整形解析错误： -9672")
	}

	// float64
	if v, has := ParseValue("3.1475856").(float64); !has || v != 3.1475856 {
		t.Error("float解析失败： 3.1475856. ", v)
	}
	if v, has := ParseValue("-0.014854").(float64); !has || v != -0.014854 {
		t.Error("float解析失败： -0.014854. ", v)
	}

	// []int
	if v, has := ParseValue("2,0,0,9,0,6,1,7").([]int); !has ||
		fmt.Sprintf("%v", []int{2, 0, 0, 9, 0, 6, 1, 7}) != fmt.Sprintf("%v", v) {
		pv := ParseValue("2,0,0,9,0,6,1,7")
		t.Errorf("[]int 解析失败，(%T| %v), raw: %T, %v", v, v, pv, pv)
	}
	if v, has := ParseValue("-2,0,0,9,0,-6,-1,7").([]int); !has ||
		fmt.Sprintf("%v", []int{-2, 0, 0, 9, 0, -6, -1, 7}) != fmt.Sprintf("%v", v) {
		pv := ParseValue("-2,0,0,9,0,-6,-1,7")
		t.Errorf("[]int 解析失败，(%T| %v), raw: %T, %v", v, v, pv, pv)
	}

	// []float64
	if v, has := ParseValue("-2.0019,0.617,21.41").([]float64); !has ||
		fmt.Sprintf("%v", []float64{-2.0019, 0.617, 21.41}) != fmt.Sprintf("%v", v) {
		pv := ParseValue("-2.0019,0.617,21.41")
		t.Errorf("[]int 解析失败，(%T| %v), raw: %T, %v", v, v, pv, pv)
	}
}

func TestStrClear(t *testing.T) {
	var input, want, get string
	const format = "%v => %v VS %v"

	// case
	input, want = `'Joshua Conero'`, "Joshua Conero"
	get = StrClear(input)
	if want != get {
		t.Fatalf(format, input, get, want)
	}

	// case
	input, want = `"Joshua Conero"`, "Joshua Conero"
	get = StrClear(input)
	if want != get {
		t.Fatalf(format, input, get, want)
	}

	// case
	input, want = `"'Joshua Conero'"`, "'Joshua Conero'"
	get = StrClear(input)
	if want != get {
		t.Fatalf(format, input, get, want)
	}

	// case
	input, want = `"Joshua", "Conero"`, `"Joshua", "Conero"`
	get = StrClear(input)
	if want != get {
		t.Fatalf(format, input, get, want)
	}

	// case
	input, want = `'Joshua', 'Conero'`, `'Joshua', 'Conero'`
	get = StrClear(input)
	if want != get {
		t.Fatalf(format, input, get, want)
	}

	// case
	input, want = `"Joshua Conero'`, `"Joshua Conero'`
	get = StrClear(input)
	if want != get {
		t.Fatalf(format, input, get, want)
	}
}
