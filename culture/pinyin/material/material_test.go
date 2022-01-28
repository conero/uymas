package material

import "testing"

func TestGetDickRaw(t *testing.T) {
	t.Logf("Raw: \n%v", string(GetDickRaw()))
}

func TestGetCommonRaw(t *testing.T) {
	t.Logf("Raw: \n%v", string(GetCommonRaw()))
}

func TestNewPinyin(t *testing.T) {
	py := NewPinyin()

	words := "古丞秋"
	tone := py.GetPyTone(words)
	t.Logf("GetPyTone: %v -> %v", words, tone)
	tone = py.GetPyToneNumber(words)
	t.Logf("GetPyToneNumber: %v -> %v", words, tone)
	tone = py.GetPyToneAlpha(words)
	t.Logf("GetPyToneAlpha: %v -> %v", words, tone)
}

func TestNewCommonList(t *testing.T) {
	cl := NewCommonList()
	//t.Logf("CL: %#v", cl)
	//t.Logf("CL.dicks: %#v", cl.dicks)
	t.Logf("CL.StrokesList: %#v", cl.StrokesList())
	t.Logf("CL.WordList(7): %#v", cl.WordList(7))
}
