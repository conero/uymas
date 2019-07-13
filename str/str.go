// 字符串处理工具包
package str

import (
	"bytes"
	"fmt"
	"html/template"
	"math"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

// @Date：   2018/10/30 0030 15:14
// @Author:  Joshua Conero
// @Name:    字符串

const (
	NumberStr = "0123456789"
	LowerStr  = "abcdefghijklmnopqrstuvwxyz"
	UpperStr  = "ABCDEFGHJIKLMNOPQRSTUVWXYZ"
)

// 写入器导出为内容
type WriterToContent struct {
	content string
}

// 实现写入器语法
func (wr *WriterToContent) Write(p []byte) (n int, err error) {
	wr.content += string(p)
	fmt.Println(wr.content, "l")
	return 0, nil
}

// 获取值
func (wr *WriterToContent) Content() string {
	return wr.content
}

// 首字母大写
func Ucfirst(str string) string {
	idx := strings.Index(str, " ")
	if idx > -1 {
		newStr := []string{}
		for _, s := range strings.Split(str, " ") {
			newStr = append(newStr, Ucfirst(s))
		}
		str = strings.Join(newStr, "")
	} else {
		if len(str) > 0 {
			str = strings.ToUpper(str[0:1]) + str[1:]
		}
	}
	return str
}

// 首字母小写
func Lcfirst(str string) string {
	idx := strings.Index(str, " ")
	if idx > -1 {
		newStr := []string{}
		for _, s := range strings.Split(str, " ") {
			newStr = append(newStr, Lcfirst(s))
		}
		str = strings.Join(newStr, "")
	} else {
		if len(str) > 0 {
			str = strings.ToLower(str[0:1]) + str[1:]
		}
	}
	return str
}

// 安全字符串分割
func SplitSafe(s, sep string) []string {
	var dd []string
	s = ClearSpace(s)
	dd = strings.Split(s, sep)
	return dd
}

// 清除空格
func ClearSpace(s string) string {
	s = strings.TrimSpace(s)
	if strings.Index(s, " ") > -1 {
		spaceReg := regexp.MustCompile("\\s")
		s = spaceReg.ReplaceAllString(s, "")
	}
	return s
}

// 根据 go template 模板编译后返回数据
// 支持 template 模板语法
func Render(tpl string, data interface{}) (string, error) {
	var value string
	temp, err := template.New("Render").Parse(tpl)
	if err != nil {
		return "", err
	}
	var bf bytes.Buffer
	err2 := temp.Execute(&bf, data)
	if err2 == nil {
		return bf.String(), nil
	}
	return value, err2
}

// 字符串反转
func Reverse(s string) string {
	sQue := strings.Split(s, "")
	sQueLen := len(sQue)
	sNewQue := []string{}
	for i := sQueLen - 1; i > -1; i-- {
		sNewQue = append(sNewQue, sQue[i])
	}
	return strings.Join(sNewQue, "")
}

// 向左填充
func PadLeft(s string, pad string, le int) string {
	sLen := len(s)
	if sLen < le {
		le -= sLen
		padLen := len(pad)
		n := math.Ceil(float64(le) / float64(padLen))
		pref := strings.Repeat(pad, int(n))
		prefLen := len(pref)
		if prefLen > le {
			pref = pref[prefLen-le:]
		}
		s = pref + s
	}
	return s
}

// 向右填充
func PadRight(s string, pad string, le int) string {
	sLen := len(s)
	if sLen < le {
		le -= sLen
		padLen := len(pad)
		n := math.Ceil(float64(le) / float64(padLen))
		end := strings.Repeat(pad, int(n))
		prefLen := len(end)
		if prefLen > le {
			end = end[:le]
		}
		s = s + end
	}
	return s
}

// 数据字符串生成基函数
func RandStrBase(base string, length int) string {
	var s string
	vlen := len(base)

	if vlen > 0 {
		ss := []string{}
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		for i := 0; i < length; i++ {
			x := r.Intn(vlen)
			ss = append(ss, base[x:x+1])
		}
		s = strings.Join(ss, "")
	}
	return s
}

// 随机字符串生成器
type RandString struct {
}

// 随机数字
func (rs RandString) Number(length int) string {
	return RandStrBase(NumberStr, length)
}

// 随机小写字母
func (rs RandString) Lower(length int) string {
	return RandStrBase(LowerStr, length)
}

// 随机大写字母
func (rs RandString) Upper(length int) string {
	return RandStrBase(UpperStr, length)
}

// 随机字母
func (rs RandString) Letter(length int) string {
	return RandStrBase(LowerStr+UpperStr, length)
}

// 随机字符串
// 包含： +_.空格/
func (rs RandString) String(length int) string {
	base := NumberStr + LowerStr + UpperStr + "-_./ $!#%&:;@^|{}[]~`"
	return RandStrBase(base, length)
}

// 随机安全字符，没有特殊符号
func (rs RandString) SafeStr(length int) string {
	base := NumberStr + LowerStr + UpperStr + "-_"
	return RandStrBase(base, length)
}

// 随机字符串
var RandStr RandString
