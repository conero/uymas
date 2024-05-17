package svn

import (
	"encoding/xml"
	"gitee.com/conero/uymas/fs"
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

// XIRepo repository
type XIRepo struct {
	XMLName xml.Name `xml:"repository"`
	Root    string   `xml:"root"`
	Uuid    string   `xml:"uuid"`
}

// XIWc wc-info
type XIWc struct {
	XMLName xml.Name `xml:"wc-info"`
	Path    string   `xml:"wcroot-abspath"`
}

// XICommit commit
type XICommit struct {
	XMLName  xml.Name `xml:"commit"`
	Author   string   `xml:"author"`
	Date     string   `xml:"date"`
	Revision string   `xml:"revision,attr"`
}

// XIEnter enter
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

// XmlInfo info xml format
type XmlInfo struct {
	XMLName xml.Name `xml:"info"`
	Enter   XIEnter
}

// Url the the repository url
func (x *XmlInfo) Url() string {
	return x.Enter.Url
}

// Author get author
func (x *XmlInfo) Author() string {
	return x.Enter.Commit.Author
}

func (x *XmlInfo) Date() string {
	return x.Enter.Commit.Date
}

// Revision get the svn version/revision
func (x *XmlInfo) Revision() string {
	return x.Enter.Commit.Revision
}

func (x *XmlInfo) Uuid() string {
	return x.Enter.Repo.Uuid
}

// Bridge cLI command bridge between SVN and SVN
type Bridge struct {
	Path string // svn path
}

func (b *Bridge) GetArgs(args ...string) []string {
	if b.Path != "" {
		b.Path = fs.StdDir(b.Path)
		args = append(args, b.Path)
	}
	return args
}

// Info svn info --xml
func (b *Bridge) Info(pArgs ...string) (XmlInfo, error) {
	args := b.GetArgs("info", "--xml")
	// 附加参数
	if len(pArgs) > 1 {
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

// XmlLog log --xml output format
type XmlLog struct {
	XMLName xml.Name  `xml:"log"`
	Enter   []XLEnter `xml:"logentry"`
}

// Log svn log --xml
func (b *Bridge) Log(pArgs ...string) (XmlLog, error) {
	args := b.GetArgs("log", "--xml")
	// 附加参数
	if len(pArgs) > 1 {
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
