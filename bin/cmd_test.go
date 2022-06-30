package bin

import "testing"

func TestApp_GetDoc(t *testing.T) {
	a := &App{
		Title:       "uymas",
		Description: "the standard(beautiful) format cmd application",
		CmdList: []AppCmd{
			{
				Alias:    []string{"-n"},
				Name:     "name",
				Title:    "设置或查看 name 属性",
				Option:   nil,
				Register: nil,
			},
		},
	}

	t.Logf("\n" + a.GetDoc())
}

func TestAppOptionGroup_ParseEach(t *testing.T) {
	aop := &AppOptionGroup{}

	type testOption struct {
		Url string `cmd:"url, u; required"`
	}

	var vop testOption
	err := aop.ParseEach(&vop, func(opt *AppOption) {
	})
	if err != nil {
		t.Errorf("Parse Error: %v", err)
	}

	url := aop.Option("url")
	if url.Name != "url" {
		t.Errorf("the name of Url param parse error!")
	}
	if url.Validation != OptValidationRequire {
		t.Errorf("Url parse error，(%v != %v)!", url.Validation, OptValidationRequire)
	}
}
