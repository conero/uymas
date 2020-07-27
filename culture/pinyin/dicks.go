package pinyin

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

//从 ini 文加读取字典
func ReadDickFromIni(filename string) map[string]map[string]string {
	dick := map[string]map[string]string{}

	fh, err := os.Open(filename)
	if err == nil {
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
	}

	return dick
}

// 读取ini文件行
//支持， [key] 等键值
//支持重复，行解析
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

// 是否为汉字
func IsChineseCharacters(word string) bool {
	return CcReg.MatchString(word)
}

//the line from file
func GetLinesFromFile(filename string) []string {
	fh, err := os.Open(filename)
	if err == nil {
		buf := bufio.NewReader(fh)
		lines := []string{}
		linesCase := []string{}
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
	return nil
}
