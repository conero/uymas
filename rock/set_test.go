package rock

import (
	"testing"

	"gitee.com/conero/uymas/v2/str"
)

func TestListIndex(t *testing.T) {
	var ss = []string{"I", "am", "Joshua", "Conero", "."}
	var idx, rfIdx int
	rfIdx = 3
	idx = ListIndex(ss, "Conero")
	if idx != rfIdx {
		t.Errorf("Search []string Index failure: %v != %v", idx, rfIdx)
	}

	//
	var its = []uint8{1, 9, 9, 2, 1, 9, 4, 9}
	idx = ListIndex(its, 1)
	rfIdx = 0
	if idx != rfIdx {
		t.Errorf("Search []uint8 Index failure: %v != %v", idx, rfIdx)
	}
}

func TestListEq(t *testing.T) {
	intArr1 := []int{1, 9, 49, 1001, 24, 903}
	intArr2 := []int{903, 24, 49}

	// case
	if ListEq(intArr1, intArr2) {
		t.Errorf("%#v = %#v，次判别错误", intArr1, intArr2)
	}

	// case
	intArr2 = []int{903, 24, 1001, 49, 9, 1}
	if !ListEq(intArr1, intArr2) {
		t.Errorf("%#v ≠ %#v，次判别错误", intArr1, intArr2)
	}

	intArr1, intArr2 = nil, nil
	if !ListEq(intArr1, intArr2) {
		t.Errorf("%#v ≠ %#v，次判别错误", intArr1, intArr2)
	}

}

func TestListSubset(t *testing.T) {
	intArr1 := []int{1, 9, 49, 1001, 24, 903}
	intArr2 := []int{903, 24, 49}

	// case
	if !ListSubset(intArr1, intArr2) {
		t.Errorf("%#v 应该为 %#v 的子数组，次判别错误", intArr1, intArr2)
	}

	strArr1 := `I am Jc, Coder.`
	strArr2 := "Coder."

	// case
	if !ListSubset([]byte(strArr1), []byte(strArr2)) {
		t.Errorf("%#v 应该为 %#v 的子数组，次判别错误", intArr1, intArr2)
	}
}

func TestListNext(t *testing.T) {
	// case
	intArr := []int{1, 9, 49, 1001, 24, 903}
	intNext, index := ListNext(intArr, 49)
	if intNext != 1001 || index != 3 {
		t.Errorf("ListNext(%#v, 49) 错误，返回值：%#v, %#v", intArr, intNext, index)
	}
	intNext, index = ListNext(intArr, 49, 2)
	if intNext != 24 || index != 4 {
		t.Errorf("ListNext(%#v, 49) 错误，返回值：%#v, %#v", intArr, intNext, index)
	}

	// case
	strArr := []string{"中国", "首都", "是", "北京", "火星", "real", "make"}
	strNext, index := ListNext(strArr, "北京")
	if strNext != "火星" || index != 4 {
		t.Errorf("ListNext(%#v, %#v) 错误，返回值：%#v, %#v", strArr, "北京", strNext, index)
	}
	strNext, index = ListNext(strArr, "make")
	if strNext != "" || index != -1 {
		t.Errorf("ListNext(%#v, %#v) 错误，返回值：%#v, %#v", strArr, "make", strNext, index)
	}
}

func TestListReverse(t *testing.T) {
	intArr := []int{1, 9, 49, 1001, 24, 903}
	var newArr = make([]int, len(intArr))
	copy(newArr, intArr)
	//t.Logf("newArr: %v", newArr)
	ListReverse(newArr)
	if intArr[0] != newArr[len(newArr)-1] {
		t.Errorf("ListReverse(%#v) 错误，返回值：%#v", intArr, newArr)
	} else {
		t.Logf("ListReverse(%#v) 正确，返回值：%#v", intArr, newArr)
	}

	// case
	strArr := []string{"中国", "首都", "是", "北京", "火星", "real", "make"}
	var newStrArr = make([]string, len(strArr))
	copy(newStrArr, strArr)
	ListReverse(newStrArr)
	if strArr[0] != newStrArr[len(newStrArr)-1] {
		t.Errorf("ListReverse(%#v) 错误，返回值：%#v", strArr, newStrArr)
	} else {
		t.Logf("ListReverse(%#v) 正确，返回值：%#v", strArr, newStrArr)
	}

	// case  空值测试
	newStrArr = nil
	ListReverse(newStrArr)
	if len(newStrArr) > 0 {
		t.Errorf("ListReverse(%#v) 错误，返回值：%#v", strArr, newStrArr)
	}

	// case 字符串测试
	vStr := "I am Jc, Coder. 贵州贵阳"
	// []byte() 会导致乱码
	// []rune 安全的utf8字符串分割
	vStrAsBytes := []byte(vStr)
	ListReverse(vStrAsBytes)
	newStr := string(vStrAsBytes)
	if len(newStr) != len(vStr) {
		t.Errorf("ListReverse(%#v) 错误，返回值：%v", vStr, newStr)
	} else {
		t.Logf("ListReverse(%#v) 正确，返回值：%s", vStr, newStr)
	}
	t.Logf("ListReverseString: %s", ListReverseString("中国人民万岁！世界人民万岁！"))
}

// 测试样本
const testTxtUtf8 = `In today's digital era, technology has become an integral part of our daily lives. 从智能手机到人工智能，科技正在重塑世界。For instance, the use of AI models like GPT has revolutionized natural language processing, enabling machines to understand and generate human-like text. 这不仅提升了效率，也带来了新的挑战，例如数据隐私和伦理问题。💡Consider the equation: E = mc² — a cornerstone of modern physics. 爱因斯坦的这一发现改变了我们对宇宙的理解。Meanwhile, in the field of computer science, algorithms such as quicksort (平均时间复杂度: O(n log n)) are fundamental to data processing. 编程语言如 Python 和 Go 因其简洁性而广受欢迎。🚀Moreover, globalization has led to increased cross-cultural communication. 人们通过社交媒体、视频会议和即时通讯工具保持联系。Platforms like WeChat, WhatsApp, and Slack allow users to send messages, make calls, and share files across borders. 一个简单的问候“Hello! 你好！안녕하세요！👋”就能连接不同文化背景的人们。In education, online learning platforms have made knowledge more accessible. 学生可以通过Coursera、edX或国内的慕课网学习世界顶尖大学的课程。This democratization of information empowers individuals to acquire new skills, whether it's coding, design, or data analysis. 数字化学习不仅打破了地理限制，也促进了终身学习的理念。However, with great power comes great responsibility. 随着技术的发展，我们也必须关注其社会影响。Issues such as misinformation, cyberbullying, and digital addiction require careful attention and regulation. 我们需要建立更安全、更包容的网络环境。🌐In conclusion, the integration of technology and global connectivity offers immense opportunities. 只要我们合理利用，科技将成为推动社会进步的重要力量。Let us embrace innovation while upholding ethical values and human dignity. 🌱 Together, we can build a smarter, more connected, and sustainable future.`

// 单位 ns/op  表示每次操作需要花费纳秒数
func BenchmarkListReverseString(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ListReverseString(testTxtUtf8)
	}
	b.StopTimer()
}

// 测试样本
func BenchmarkListReverseStringVs(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		str.Str(testTxtUtf8).Reverse()
	}
	b.StopTimer()
}
