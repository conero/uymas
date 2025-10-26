package chest

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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
		if validFunc != nil {
			if validFunc(text) {
				break
			}
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
	}
	return text
}

// LineAsArgs line as args, parse for command
func LineAsArgs(line string) []string {
	line = strings.TrimSpace(line)
	if line == "" {
		return nil
	}
	reg := regexp.MustCompile(`\s{2,}`)
	// "reg"
	findReg := regexp.MustCompile(`"[^"]+"`)
	rpl := line
	var mapValue = map[string]string{}
	index := 0
	for _, fd := range findReg.FindAllString(line, -1) {
		replValue := fmt.Sprintf("_|%d`", index)
		mapValue[replValue] = fd[1 : len(fd)-1]
		rpl = strings.Replace(rpl, fd, replValue, -1)
		index++
	}

	// 'reg'
	findReg = regexp.MustCompile(`'[^']+'`)
	for _, fd := range findReg.FindAllString(rpl, -1) {
		replValue := fmt.Sprintf("_|%d`", index)
		mapValue[replValue] = fd[1 : len(fd)-1]
		rpl = strings.Replace(rpl, fd, replValue, -1)
		index++
	}

	rpl = reg.ReplaceAllString(rpl, " ")
	newList := strings.Split(rpl, " ")
	if len(mapValue) > 0 {
		for i, v := range newList {
			if vRpl, ok := mapValue[v]; ok {
				newList[i] = vRpl
			}
		}
	}

	return newList
}
