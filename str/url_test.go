package str

import "testing"

/**
 * @DATE        2019/6/3
 * @NAME        Joshua Conero
 * @DESCRIPIT   测试文档
 **/

func TestUrl_AbsHref(t *testing.T) {
	var vp, vu string
	var want, out string

	var u Url

	// case
	vp, vu = "about/joshua", "https://conero.cn/visual/link?a=www#test"
	want = "https://conero.cn/visual/link/about/joshua"
	out = u.AbsHref(vp, vu)
	if out != want {
		t.Error(out, " != ", want)
	}

	// case
	vp, vu = "./about/joshua", "https://conero.cn/visual/link?a=www#test"
	want = "https://conero.cn/visual/link/about/joshua"
	out = u.AbsHref(vp, vu)
	if out != want {
		t.Error(out, " != ", want)
	}

	// case
	vp, vu = "../about/joshua", "https://conero.cn/visual/link?a=www#test"
	want = "https://conero.cn/visual/about/joshua"
	out = u.AbsHref(vp, vu)
	if out != want {
		t.Error(out, " != ", want)
	}

	// case
	vp, vu = "../..//about/joshua", "https://conero.cn/visual/link?a=www#test"
	want = "https://conero.cn/about/joshua"
	out = u.AbsHref(vp, vu)
	if out != want {
		t.Error(out, " != ", want)
	}

	// case
	vp, vu = "../..//about/joshua", "https://conero.cn/visual/link"
	want = "https://conero.cn/about/joshua"
	out = u.AbsHref(vp, vu)
	if out != want {
		t.Error(out, " != ", want)
	}
}

func TestUrl_AbsHref2(t *testing.T) {
	var U Url

	var u, c, w string
	u = "https://www.bigdata-expo.cn/"

	// cs
	w = u
	c = U.AbsHref("", u)
	if c != w {
		t.Error(c)
	}

	// news/1487902393
	// cs
	w = "https://www.bigdata-expo.cn/news/1487902393"
	c = U.AbsHref("news/1487902393", u)
	if c != w {
		t.Error(c)
	}

	// news/1487902393
	// cs
	w = "https://www.bigdata-expo.cn/expo/exhireco"
	c = U.AbsHref("../../expo/exhireco", "https://www.bigdata-expo.cn/news/1487902393")
	if c != w {
		t.Error(c)
	}

	// news/1487902393
	// cs
	w = "https://www.bigdata-expo.cn/news/1467358512"
	c = U.AbsHref("./1467358512", "https://www.bigdata-expo.cn/news/1487902393")
	if c != w {
		t.Error(c)
	}
}
