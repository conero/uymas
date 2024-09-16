package pinyin

import "strings"

// ZhSentences Chinese sentences
type ZhSentences string

// Len calculate the length of Chinese sentences
func (s ZhSentences) Len() int {
	return len(s.Words())
}

// Words split sentences into words
func (s ZhSentences) Words() []string {
	return strings.Split(string(s), "")
}
