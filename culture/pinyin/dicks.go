package pinyin

import (
	"bufio"
	"bytes"
	"os"
	"regexp"
	"strings"
)

// ReadDickFromIni read file and parse from ini files
func ReadDickFromIni(filename string) map[string]map[string]string {
	dick := map[string]map[string]string{}

	fh, err := os.Open(filename)
	if err == nil {
		return dick
	}

	buf := bufio.NewReader(fh)
	var mapKey string
	var regKey = regexp.MustCompile(`^\[[^\[\]]+]+`)
	var regStr = regexp.MustCompile(`\[|]`)
	var equalStr = "="
	for {
		line, err2 := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		// "#/;" 开头含忽略
		w := ""
		if line != "" {
			w = line[:1]
		}
		if w == "#" || w == ";" {
			line = ""
		}
		if line != "" {
			if regKey.MatchString(line) {
				mapKey = regStr.ReplaceAllString(line, "")
			} else if mapKey != "" {
				idx := strings.Index(line, equalStr)
				if idx > -1 {
					tKey := strings.TrimSpace(line[:idx])
					tValue := strings.TrimSpace(line[idx+1:])
					v, has := dick[mapKey]
					if has {
						v[tKey] = tValue
						dick[mapKey] = v
					} else {
						dick[mapKey] = map[string]string{
							tKey: tValue,
						}
					}
				}
			}
		}
		// 错误
		if err2 != nil {
			break
		}
	}

	return dick
}

// ReadDickFromByteKv read file and parse from ini files
func ReadDickFromByteKv(content []byte) map[string]string {
	dick := map[string]string{}

	bf := bytes.NewBuffer(content)
	buf := bufio.NewReader(bf)
	var equalStr = "="
	for {
		line, err2 := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		// "#/;" 开头含忽略
		w := ""
		if line != "" {
			w = line[:1]
		}
		if w == "#" || w == ";" {
			line = ""
		}
		if line != "" {
			idx := strings.Index(line, equalStr)
			if idx > -1 {
				tKey := strings.TrimSpace(line[:idx])
				tValue := strings.TrimSpace(line[idx+1:])
				dick[tKey] = tValue
			}
		}
		// 错误
		if err2 != nil {
			break
		}
	}

	return dick
}

// ReadIniLines read ini file as line array, support [key], and repeat parse line.
func ReadIniLines(lines []string) map[string]map[string]string {
	dick := map[string]map[string]string{}

	var mapKey string
	var regKey = regexp.MustCompile(`^\[[^\[\]]+]+`)
	var regStr = regexp.MustCompile(`\[|]`)
	var equalStr = "="
	for _, line := range lines {
		line = strings.TrimSpace(line)
		// "#/;" 开头含忽略
		w := ""
		if line != "" {
			w = line[:1]
		}
		if w == "#" || w == ";" {
			line = ""
		}
		if line != "" {
			if regKey.MatchString(line) {
				mapKey = regStr.ReplaceAllString(line, "")
			} else if mapKey != "" {
				idx := strings.Index(line, equalStr)
				if idx > -1 {
					tKey := strings.TrimSpace(line[:idx])
					tValue := strings.TrimSpace(line[idx+1:])
					v, has := dick[mapKey]
					if has {
						//同一队列，支持多次重复写如：
						//      key = 22222,sdsd,
						//      key = 3333,cdd
						//      ;等同 key = 22222,sdsd,3333,cdd
						cv, cHas := v[tKey]
						if cHas {
							cv = v[tKey] + tValue
							v[tKey] = cv
						} else {
							v[tKey] = tValue
						}
						dick[mapKey] = v
					} else {
						dick[mapKey] = map[string]string{
							tKey: tValue,
						}
					}
				}
			}
		}
	}

	return dick
}

// IsChineseCharacters check if the string is chinese character
func IsChineseCharacters(word string) bool {
	return CcReg.MatchString(word)
}

// GetLinesFromFile the line from file
func GetLinesFromFile(filename string) []string {
	fh, err := os.Open(filename)
	if err != nil {
		return nil
	}
	defer fh.Close()
	buf := bufio.NewReader(fh)
	var lines []string
	var linesCase []string
	for {
		line, err2 := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		// "#/;" 开头含忽略
		w := ""
		if line != "" {
			w = line[:1]
		}
		if w == "#" || w == ";" {
			line = ""
		}
		if line != "" {
			lines = append(lines, line)
			linesCase = append(linesCase, line)
		}
		// 错误
		if err2 != nil {
			break
		}
	}
	return lines
}

// GetLinesFromByte the line byte
func GetLinesFromByte(content []byte) []string {
	bf := bytes.NewBuffer(content)
	buf := bufio.NewReader(bf)
	var lines []string
	var linesCase []string
	for {
		line, err2 := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		// "#/;" 开头含忽略
		w := ""
		if line != "" {
			w = line[:1]
		}
		if w == "#" || w == ";" {
			line = ""
		}
		if line != "" {
			lines = append(lines, line)
			linesCase = append(linesCase, line)
		}
		// 错误
		if err2 != nil {
			break
		}
	}
	return lines
}
