package str

const (
	NumberStr = "0123456789"
	LowerStr  = "abcdefghijklmnopqrstuvwxyz"
	UpperStr  = "ABCDEFGHJIKLMNOPQRSTUVWXYZ"
)

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
