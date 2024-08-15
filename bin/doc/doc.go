// Package doc command line help information, [Experimental]
//
// todolist
//
// 1. Support direct document generation through registration
//
// 2.Support generating documents directly through embedded syntax
package doc

import (
	"gitee.com/conero/uymas/util/rock"
	"regexp"
	"strings"
)

type HelpDick struct {
	// command alias
	Alias []string `json:"alias"`
	// command title
	Title string `json:"title"`
	// command option, format like:
	//
	// [file,f => desc], support mutilOption
	Option map[string]string `json:"option"`
	// command help by all text
	AllText string `json:"allText"`
}

type Doc struct {
	// [command => help message] current doc text
	CommandHelp []HelpDick `json:"commandHelp"`
	// all text
	// [lang => HelpDick]
	MutilLang map[string][]HelpDick `json:"mutilLang"`
	// all text
	AllText string `json:"allText"`
	// support lang, if not then fist
	LangList []string `json:"langList"`
	// current lang
	curLang string
}

// FromLine parse doc from content by strig line
func FromLine(content string, langs ...string) *Doc {
	queue := strings.Split(content, "\n")
	doc := &Doc{
		curLang: rock.ExtractParam("", langs...),
	}

	var (
		curLang   string
		curCmd    string
		curHelp   HelpDick
		allText   string
		cmdText   string
		mutilLang map[string][]HelpDick
	)

	// flush curHelp
	flushCurHelpFn := func() {
		if len(curHelp.Alias) == 0 {
			return
		}

		// update
		lastHd := mutilLang[curLang]
		curHelp.AllText = cmdText
		lastHd = append(lastHd, curHelp)
		mutilLang[curLang] = append(mutilLang[curLang], curHelp)

		// flush
		curHelp.Alias = []string{}
		curCmd = ""
		cmdText = ""
	}

	cmdNameReg := regexp.MustCompile(`^\w+(,\w+)*`)
	optNameReg := regexp.MustCompile(`^(-{1,2}\w+)+`)
	for _, line := range queue {
		ln := strings.TrimSpace(line)
		if ln == "" {
			continue
		}
		first := ln[:1]
		// ";" or "#" is comment
		if first == ";" || first == "#" {
			continue
		}
		// ":key = value"
		if first == ":" {
			pk, pv := ParseKv(ln)
			// flush lang
			_, findLast := mutilLang[curLang]
			if curLang != pv && findLast {
				flushCurHelpFn()
			}
			if pk == "lang" {
				if rock.ListIndex(doc.LangList, pv) == -1 {
					doc.LangList = append(doc.LangList, pv)
				}
				curLang = pv
			}
			continue
		}

		if mutilLang == nil {
			mutilLang = map[string][]HelpDick{}
		}
		hd := mutilLang[curLang]
		if allText != "" {
			allText += "\n"
		}
		allText += line

		// command
		if cmdNameReg.MatchString(ln) {
			cmdPairs := cmdNameReg.FindAllString(ln, -1)
			cmdList := strings.Split(cmdPairs[0], ",")
			//fmt.Printf("cmdPairs: (%s, %s), %s, %#v\n",
			//	curCmd, cmdList[0], curLang, cmdList)
			vCmd := cmdList[0]
			if vCmd != curCmd && curCmd != "" {
				curHelp.AllText = cmdText
				//fmt.Printf("curHelp: %#v\n", curHelp)
				mutilLang[curLang] = append(mutilLang[curLang], curHelp)
				hd = append(hd, curHelp)
				cmdText = ""
			}

			// flush
			curCmd = vCmd
			curHelp = HelpDick{
				Alias: cmdList,
				Title: strings.TrimSpace(strings.ReplaceAll(ln, cmdPairs[0], "")),
			}
		} else if optNameReg.MatchString(ln) { // option
			optPairs := optNameReg.FindAllString(ln, -1)
			if curHelp.Option == nil {
				curHelp.Option = map[string]string{}
			}
			optName := strings.ReplaceAll(optPairs[0], "-", "")
			curHelp.Option[optName] = strings.ReplaceAll(ln, optPairs[0], "")
		}

		if cmdText != "" {
			cmdText += "\n"
		}
		cmdText += line
		mutilLang[curLang] = hd
	}

	// last todo
	flushCurHelpFn()

	//doc.AllText = allText
	doc.MutilLang = mutilLang
	return doc
}

// ParseKv parse kv by line that format matched ":key=value"
func ParseKv(line string) (key string, value string) {
	line = strings.TrimSpace(line)
	idx := strings.Index(line, "=")
	if idx > -1 && strings.Index(line, ":") == 0 {
		key = strings.TrimSpace(line[1:idx])
		value = strings.TrimSpace(line[idx+1:])
	}
	return
}

func (c *Doc) HelpList(langs ...string) ([]HelpDick, bool) {
	lang := rock.ExtractParam(c.curLang, langs...)
	if lang == "" && len(c.LangList) > 0 {
		lang = c.LangList[0]
	}

	if lang == "" {
		return nil, false
	}

	list, exist := c.MutilLang[lang]
	return list, exist
}

func (c *Doc) Help(langs ...string) string {
	list, exist := c.HelpList(langs...)
	if !exist {
		return ""
	}
	return c.listAsHelp(list)
}

func (c *Doc) listAsHelp(list []HelpDick) string {
	var queue []string
	for _, ls := range list {
		queue = append(queue, ls.AllText)
	}
	return strings.Join(queue, "\n")
}

// Support detect if the lang is supported
func (c *Doc) Support(lng string) bool {
	return rock.ListIndex(c.LangList, lng) > -1
}

// Search help by give command to search
func (c *Doc) Search(cs ...string) (string, bool) {
	help, exist := c.HelpList()
	if !exist {
		return "", false
	}
	for _, vc := range cs {
		for _, fd := range help {
			if rock.ListIndex(fd.Alias, vc) > -1 {
				return fd.AllText, true
			}
		}
	}
	return "", false
}
