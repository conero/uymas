package svn

import (
	"encoding/xml"
	"github.com/conero/uymas/fs"
)

// @Date：   2018/12/5 0005 22:57
// @Author:  Joshua Conero
// @Name:    命令桥
// [encoding/xml] 解析过程中struct一定要是可到处类型的字段(由于其是不同包之间的库)，以及不可遗忘xml标注

/*
<?xml version="1.0" encoding="UTF-8"?>
<info>
   <entry path="." revision="4946" kind="dir">
      <url>https://121.40.183.166/svn/mci600a</url>
      <relative-url>^/</relative-url>
      <repository>
         <root>https://121.40.183.166/svn/mci600a</root>
         <uuid>24ad63d0-2e50-4f47-9934-a14da7c91441</uuid>
      </repository>
      <wc-info>
         <wcroot-abspath>D:/server/zmapp/mci600a</wcroot-abspath>
         <schedule>normal</schedule>
         <depth>infinity</depth>
      </wc-info>
      <commit revision="4946">
         <author>daiqi</author>
         <date>2018-12-06T01:10:25.592712Z</date>
      </commit>
   </entry>
</info>
*/
// repository
type XIRepo struct {
	XMLName xml.Name `xml:"repository"`
	Root    string   `xml:"root"`
	Uuid    string   `xml:"uuid"`
}

// wc-info
type XIWc struct {
	XMLName xml.Name `xml:"wc-info"`
	Path    string   `xml:"wcroot-abspath"`
}

// commit
type XICommit struct {
	XMLName  xml.Name `xml:"commit"`
	Author   string   `xml:"author"`
	Date     string   `xml:"date"`
	Revision string   `xml:"revision,attr"`
}

// enter
type XIEnter struct {
	XMLName  xml.Name `xml:"entry"`
	Path     string   `xml:"path,attr"`
	Revision string   `xml:"revision,attr"`
	Kind     string   `xml:"kind,attr"`
	Url      string   `xml:"url"`
	Repo     XIRepo
	Wc       XIWc
	Commit   XICommit
}

// info xml 格式
type XmlInfo struct {
	XMLName xml.Name `xml:"info"`
	Enter   XIEnter
}

// 地址
func (x *XmlInfo) Url() string {
	return x.Enter.Url
}

// author
func (x *XmlInfo) Author() string {
	return x.Enter.Commit.Author
}

// 日期
func (x *XmlInfo) Date() string {
	return x.Enter.Commit.Date
}

// 版本信息
func (x *XmlInfo) Revision() string {
	return x.Enter.Commit.Revision
}

// uuid
func (x *XmlInfo) Uuid() string {
	return x.Enter.Repo.Uuid
}

// 与 svn 之间的cli-命令桥
type Bridge struct {
	Path string // svn path
}

// 获取命令
func (b *Bridge) GetArgs(args ...string) []string {
	if b.Path != "" {
		b.Path = fs.StdDir(b.Path)
		args = append(args, b.Path)
	}
	return args
}

// svn info --xml
func (b *Bridge) Info(pArgs ...string) (XmlInfo, error) {
	args := b.GetArgs("info", "--xml")
	// 附加参数
	if pArgs != nil && len(pArgs) > 1 {
		args = append(args, pArgs...)
	}
	out, err := Call(args...)
	var dd XmlInfo
	if err != nil {
		return dd, err
	}
	//println(out)
	// 解析XML
	//dd := XmlInfo{}
	err = xml.Unmarshal([]byte(out), &dd)
	if err != nil {
		return dd, err
	}
	return dd, nil
}

/*
<?xml version="1.0" encoding="UTF-8"?>
<log>
   <logentry revision="4907">
      <author>joshua</author>
      <date>2018-12-04T03:52:28.702087Z</date>
      <msg></msg>
   </logentry>
   <logentry revision="4904">
      <author>conero</author>
      <date>2018-12-04T03:01:11.280212Z</date>
      <msg></msg>
   </logentry>
</log>
*/

type XLEnter struct {
	XMLName  xml.Name `xml:"logentry"`
	Revision string   `xml:"revision,attr"`
	Author   string   `xml:"author"`
	Date     string   `xml:"date"`
	Msg      string   `xml:"msg"`
}

// log --xml 输出格式
type XmlLog struct {
	XMLName xml.Name `xml:"log"`
	Enter   []XLEnter `xml:"logentry"`
}

// svn log --xml
func (b *Bridge) Log(pArgs ...string) (XmlLog, error) {
	args := b.GetArgs("log", "--xml")
	// 附加参数
	if pArgs != nil && len(pArgs) > 1 {
		args = append(args, pArgs...)
	}
	//fmt.Println(args)
	out, err := Call(args...)
	var dd XmlLog
	if err != nil {
		return dd, err
	}

	// 解析XML
	err = xml.Unmarshal([]byte(out), &dd)
	if err != nil {
		return dd, err
	}
	return dd, nil
}
