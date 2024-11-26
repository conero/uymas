package butil

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// InputRequire required input data from stdin
func InputRequire(title string, validFunc func(string) bool) string {
	var input = bufio.NewScanner(os.Stdin)
	fmt.Print(title)
	var text string
	for input.Scan() {
		text = input.Text()
		text = strings.TrimSpace(text)
		if validFunc != nil && validFunc(text) {
			break
		} else if text != "" {
			break
		}
		fmt.Print(title)
	}
	return text
}

// InputOption need input data from stdin optional
func InputOption(title, def string) string {
	var input = bufio.NewScanner(os.Stdin)
	fmt.Print(title)
	var text string
	for input.Scan() {
		text = input.Text()
		text = strings.TrimSpace(text)
		if text == "" {
			text = def
		}
		break
	}
	return text
}
