package svn

import (
	"fmt"
	"testing"
)

var _testPath = "D:/server/zmapp/mci600a/"

// @Date：   2018/12/6 0006 13:38
// @Author:  Joshua Conero
// @Name:    名称描述
// @link:    https://golang.google.cn/pkg/encoding/xml/#example_Unmarshal

func TestBridge_Info(t *testing.T) {
	brd := &Bridge{Path: _testPath}
	dd, er := brd.Info()
	if er != nil {
		fmt.Println(" Error: " + er.Error())
		return
	}
	fmt.Println("Revision: " + dd.Revision())
	fmt.Println("Url: " + dd.Url())
	fmt.Println("Author: " + dd.Author())
	fmt.Println("Date: " + dd.Date())
	fmt.Println("Uuid: " + dd.Uuid())
}

func TestBridge_Log(t *testing.T) {
	brd := &Bridge{Path: _testPath}
	dd, er := brd.Log("-r", "4900:4903")
	if er != nil {
		fmt.Println(" Error: " + er.Error())
		return
	}
	//fmt.Println(dd)
	//fmt.Println(dd.Enter)
	for _, d := range dd.Enter {
		s := `{"revision": "` + d.Revision + `", "author": "` + d.Author + `", "date": "` + d.Date + `", "msg": "` + d.Msg + `"}`
		fmt.Println(" " + s)
		//fmt.Println(d)
	}
}
