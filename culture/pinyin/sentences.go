package pinyin

import "strings"

// ZhSentences Chinese sentences
type ZhSentences string

// Len calculate the length of Chinese sentences
func (s ZhSentences) Len() int {
	queue := strings.Split(string(s), "")
	return len(queue)
}
