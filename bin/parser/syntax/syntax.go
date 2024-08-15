// Package syntax is bin language syntax extend, like digital computing.
package syntax

import (
	"fmt"
	"gitee.com/conero/uymas/str"
	"regexp"
	"strconv"
	"strings"
)

// NumberOpera @TODO Need to do more optimize.
// NumberOpera parse string equation for number digital.
func NumberOpera(equation string) float64 {
	equation = str.ClearSpace(equation)
	var result float64
	// equation contain `()`
	reg := regexp.MustCompile(`^\([^(]+\)+`)
	if reg.MatchString(equation) {
		equation = equation[1 : len(equation)-1]
	}
	reg = regexp.MustCompile(`\([^(]+\)`)
	for {
		for _, brackets := range reg.FindAllString(equation, -1) {
			//fmt.Println(brackets)
			res := NumberOpera(brackets)
			equation = strings.ReplaceAll(equation, brackets, fmt.Sprintf("%v", res))
		}
		if !reg.MatchString(equation) {
			break
		}
	}
	equation = strings.TrimSpace(equation)
	equation = strings.ReplaceAll(equation, "--", "+") // `--` -> `+`
	//fmt.Println("  M ", equation)
	// add , subtract , multiply and divide
	// check -> [+,-,*,/]
	multiplyReg := regexp.MustCompile(`[\d.]+[-]*[*][-]*[\d.]+`)
	divideReg := regexp.MustCompile(`[\d.]+[-]*[/][-]*[\d.]+`)
	subtractReg := regexp.MustCompile(`[\d.]+[+-][\d.]+`)
	symCheckReg := regexp.MustCompile(`[*/]+`)
	for {
		// `*`
		for {
			for _, mul := range multiplyReg.FindAllString(equation, -1) {
				queue := strings.Split(mul, "*")
				tV := "1"
				if len(queue) == 2 {
					tV = fmt.Sprintf("%v", stringToF64(queue[0])*stringToF64(queue[1]))
				}
				equation = strings.ReplaceAll(equation, mul, tV)
			}
			if strings.Contains(equation, "*") {
				break
			}
		}
		// `/`
		for {
			for _, div := range divideReg.FindAllString(equation, -1) {
				queue := strings.Split(div, "/")
				tV := "1"
				if len(queue) == 2 {
					tV = fmt.Sprintf("%v", stringToF64(queue[0])/stringToF64(queue[1]))
				}
				equation = strings.ReplaceAll(equation, div, tV)
			}
			if strings.Contains(equation, "*") {
				break
			}
		}
		//fmt.Println("  L ", equation)
		for {
			equation = strings.TrimSpace(equation)
			// `+`/`-`
			for _, vAs := range subtractReg.FindAllString(equation, -1) {
				var v float64 = 0
				if strings.Contains(vAs, "+") {
					queue := strings.Split(vAs, "+")
					vqCk := len(queue)
					if vqCk == 2 {
						v = stringToF64(queue[0]) + stringToF64(queue[1])
					} else {
						v = stringToF64(queue[0])
					}
					equation = strings.ReplaceAll(equation, vAs, fmt.Sprintf("%v", v))
					//fmt.Println(" ~~ ", equation)
					break
				}
				if strings.Contains(vAs, "-") {
					queue := strings.Split(vAs, "-")
					vqCk := len(queue)
					if vqCk == 2 {
						v = stringToF64(queue[0]) - stringToF64(queue[1])
					} else {
						v = stringToF64(queue[0])
					}
					// report last
					result = v
					equation = strings.ReplaceAll(equation, vAs, fmt.Sprintf("%v", v))
					//fmt.Println(" ~~ ", equation)
					break
				}
			}
			// check `+/-`
			if !subtractReg.MatchString(equation) {
				break
			}
			//fmt.Println(" ~~ @", equation)
		}

		equation = strings.TrimSpace(equation)
		if !symCheckReg.MatchString(equation) {
			break
		}
	}
	result = stringToF64(equation)
	return result
}

// string cover float64
func stringToF64(vStr string) float64 {
	vStr = str.ClearSpace(vStr)
	if v, er := strconv.ParseFloat(vStr, 64); er == nil {
		return v
	}
	return 0
}
