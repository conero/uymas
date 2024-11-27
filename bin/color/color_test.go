package color

import "testing"

func TestAnsiClear(t *testing.T) {
	raw := "It's best."
	text := Styles(raw, TextWhiteBr, BkgBlue, DimFont)
	t.Log(text)

	toTestFn := func() {
		ansiClear := Clear(raw)
		if ansiClear != raw {
			t.Errorf("%#v 清理失败", text)
		}
	}
	// case
	toTestFn()

	// case
	// 惜秦皇汉武，略输文采；唐宗宋祖，稍逊风骚。
	raw = "略输文采；唐宗宋祖，"
	text = Style(PurpleBr, raw)
	text = "惜秦皇汉武，" + text + "稍逊风骚。"
	raw = "惜秦皇汉武，" + raw + "稍逊风骚。"
	t.Log(text)
	toTestFn()

	// case
	//  一代天骄，成吉思汗，只识弯弓射大雕。
	raw = "成吉思汗，只识弯弓射大雕。"
	text = Styles(raw, TextGreenBr, BkgBlackBr, TwinkleFont)
	text = "一代天骄，" + text
	raw = "一代天骄，" + raw
	t.Log(text)
	toTestFn()
}
