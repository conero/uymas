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
