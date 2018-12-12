package svn

import (
	"encoding/xml"
	"fmt"
	"strings"
	"testing"
)

var xmlStr string =
	`
<?xml version="1.0" encoding="UTF-8"?>
<log>
    <logentry revision="4900">
        <author>yanghua</author>
        <date>2018-12-04T02:38:09.577087Z</date>
        <msg>it test msg-m12</msg>
    </logentry>
    <logentry revision="4901">
        <author>wjq</author>
        <date>2018-12-04T02:41:59.577087Z</date>
        <msg></msg>
    </logentry>
    <logentry revision="4902">
        <author>yanghua</author>
        <date>2018-12-04T02:43:07.186462Z</date>
        <msg>it test msg-m18</msg>
    </logentry>
    <logentry revision="4903">
        <author>dzj</author>
        <date>2018-12-04T02:55:23.952087Z</date>
        <msg>it test msg-m19</msg>
    </logentry>
</log>
`
type Enter struct {
	XMLName  xml.Name `xml:"logentry"`
	Revision string   `xml:"revision,attr"`
	Author   string   `xml:"author"`
	Date     string   `xml:"date"`
	Msg      string   `xml:"msg"`
}

// log --xml 输出格式
type log struct {
	XMLName xml.Name `xml:"log"`
	Enter   []Enter  `xml:"logentry"`
}
// test to get the format string
func (lg log) format()  {
	sQue := []string{}
	for _, d := range lg.Enter {
		s := `{"revision": "` + d.Revision + `", "author": "` + d.Author + `", "date": "` + d.Date + `", "msg": "` + d.Msg + `"}`
		sQue = append(sQue, s)
		//fmt.Println(d)
	}
	fmt.Println(strings.Join(sQue, ",\n"))
	// should print string:
	// 				{"revision": "4900", "author": "yanghua", "date": "2018-12-04T02:38:09.577087Z", "msg": "it test msg-m12"},
	// 				....
	// 				{"revision": "4903", "author": "dzj", "date": "2018-12-04T02:55:23.952087Z", "msg": "it test msg-m19"}
	// have the value

	// more test the data format
	fmt.Println(lg)
	fmt.Println(lg.Enter)
}

func TestBridge_Log2(t *testing.T) {
	var dd log
	err := xml.Unmarshal([]byte(xmlStr), &dd)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	dd.format()
}