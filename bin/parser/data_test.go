package parser

import "testing"

func TestJsonReceiver_Receiver(t *testing.T) {
	var rd DataReceiver = &JsonReceiver{}
	var content string
	content = `{"name":"Joshua Conero", "birth_year": 1992, "user_id": 10210625}`

	// case - json
	rd.Receiver(ReceiverContent, content)
	t.Logf("Content: %#v", rd.GetData())

	// case - url
	content = "https://httpbin.org/get"
	rd.Receiver(ReceiverUrl, content)
	t.Logf("Url: %#v", rd.GetData())
}

func TestUrlReceiver_Receiver(t *testing.T) {
	var rd DataReceiver = &UrlReceiver{}
	var content string
	// case - url
	content = "tn=monline_4_dg&ie=utf-8&wd=httpbin"
	rd.Receiver(ReceiverContent, content)
	t.Logf("Content: %#v", rd.GetData())

	// case - json - url .bad
	content = `{"name":"Joshua Conero", "birth_year": 1992, "user_id": 10210625}`
	rd.Receiver(ReceiverContent, content)
	t.Logf("Content-json: %#v", rd.GetData())
}
