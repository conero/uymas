package pinyin

// Element the data dictionary enter
type Element struct {
	Unicode string
	// possible existence of polyphonic characters
	pinyin  string
	Chinese string
}
