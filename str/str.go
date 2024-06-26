// Package str string handler method.
package str

import (
	"bytes"
	"encoding/base64"
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

// Lcfirst converts the first character of each word in a string to lowercase.
func Lcfirst(str string) string {
	idx := strings.Index(str, " ")
	if idx > -1 {
		var newStr []string
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

// Ucfirst converts the first character of each word in a string to uppercase.
func Ucfirst(str string) string {
	idx := strings.Index(str, " ")
	if idx > -1 {
		var newStr []string
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

func IsLatinAlpha(alpha string) bool {
	return strings.Contains(LowerStr, strings.ToLower(alpha))
}

// LowerStyle camelcase --> snake case
// covert string to be lower style, like:
//
//	`FirstName` 			-> `first_name`,
//	`getHeightWidthRate` 	-> `get_height_width_rate`
func LowerStyle(vStr string) string {
	vLen := len(vStr)
	if vLen > 0 {
		bys := []byte(vStr)
		var upperQueue []int
		for i := 0; i < vLen; i++ {
			alpha := vStr[i : i+1]
			if IsLatinAlpha(alpha) && alpha == strings.ToUpper(alpha) {
				upperQueue = append(upperQueue, i)
			}
		}

		var valueQueue []string
		var lastIndex = 0
		var uQLen = len(upperQueue)
		for j, v := range upperQueue {
			if v == 0 {
				continue
			}
			valueQueue = append(valueQueue, string(bys[lastIndex:v]))
			lastIndex = v
			//Last
			if j == (uQLen - 1) {
				valueQueue = append(valueQueue, string(bys[lastIndex:]))
			}
		}

		if len(valueQueue) == 0 {
			return strings.ToLower(vStr)
		}
		return strings.ToLower(strings.Join(valueQueue, "_"))
	}
	return vStr
}

// CamelCase camelcase --> snake case
// covert string to be lower style, like:
//
//	`first_name` 			-> `FirstName`,
//	`get_height_width_rate` 	-> `GetHeightWidthRate`
//
// snake case --> camelcase
func CamelCase(vStr string) string {
	if vLen := len(vStr); vLen > 0 {
		vQueue := strings.Split(vStr, "_")
		var newQue []string
		for _, vQ := range vQueue {
			vQ = strings.TrimSpace(vQ)
			if vQ == "" {
				continue
			}
			newQue = append(newQue, Ucfirst(vQ))
		}
		vStr = strings.Join(newQue, "")
	}
	return vStr
}

// SplitSafe split safe string
func SplitSafe(s, sep string) []string {
	var dd []string
	s = ClearSpace(s)
	dd = strings.Split(s, sep)
	return dd
}

// ClearSpace clear string space
func ClearSpace(s string) string {
	s = strings.TrimSpace(s)
	if strings.Contains(s, " ") {
		spaceReg := regexp.MustCompile(`\s`)
		s = spaceReg.ReplaceAllString(s, "")
	}
	return s
}

// Render 根据 go template 模板编译后返回数据
// 支持 template 模板语法
func Render(tpl string, data any) (string, error) {
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

// Reverse string reverse
func Reverse(s string) string {
	sQue := strings.Split(s, "")
	sQueLen := len(sQue)
	var sNewQue []string
	for i := sQueLen - 1; i > -1; i-- {
		sNewQue = append(sNewQue, sQue[i])
	}
	return strings.Join(sNewQue, "")
}

// PadLeft string pad substring from left.
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

// PadRight string pad substring from right.
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

// RandStrBase make rand string by base string.
func RandStrBase(base string, length int) string {
	var s string
	vLen := len(base)

	if vLen > 0 {
		var ss []string
		for i := 0; i < length; i++ {
			r := rand.New(rand.NewSource(time.Now().UnixNano() + int64(i)))
			x := r.Intn(vLen)
			ss = append(ss, base[x:x+1])
		}
		s = strings.Join(ss, "")
	}
	return s
}

// RandString rand string creator.
type RandString struct {
}

// Number get random number by length
func (rs RandString) Number(length int) string {
	return RandStrBase(NumberStr, length)
}

// Lower get lower string
func (rs RandString) Lower(length int) string {
	return RandStrBase(LowerStr, length)
}

// Upper get upper string
func (rs RandString) Upper(length int) string {
	return RandStrBase(UpperStr, length)
}

// Letter get random latin alpha.
func (rs RandString) Letter(length int) string {
	return RandStrBase(LowerStr+UpperStr, length)
}

// 随机字符串
// 包含： +_.空格/
func (rs RandString) String(length int) string {
	base := NumberStr + LowerStr + UpperStr + "-_./ $!#%&:;@^|{}[]~`"
	return RandStrBase(base, length)
}

// SafeStr get safe string not contain special symbol
func (rs RandString) SafeStr(length int) string {
	base := NumberStr + LowerStr + UpperStr + "-_"
	return RandStrBase(base, length)
}

// RandStr rand string instance
var RandStr RandString

// Base64Encode base64 string encode
func Base64Encode(origin string) string {
	return base64.StdEncoding.EncodeToString([]byte(origin))
}

// Base64Decode base64 string decode
func Base64Decode(code string) string {
	decode, err := base64.StdEncoding.DecodeString(code)
	if err != nil {
		return ""
	}
	return string(decode)
}

// GetNotEmpty get not empty by many strings.
func GetNotEmpty(strs ...string) string {
	for _, str := range strs {
		str = strings.TrimSpace(str)
		if str != "" {
			return str
		}
	}
	return ""
}
