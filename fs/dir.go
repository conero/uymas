package fs

import (
	"errors"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
	"time"
)

type ChildDirData struct {
	Name  string
	Size  int64
	IsDir bool
}

//the tool to scan the dirs.
type DirScanner struct {
	baseDir      string
	AllItem      int
	AllDirItem   int
	AllFileItem  int
	AllSize      int64
	TopChildDick map[string]ChildDirData
	Runtime      time.Duration

	//内部变量
	//排除名称，以"*"匹配（不包含）
	excludeExp []string
	//过滤表达式（包含）
	includeExp []string
}

//exclude exp for dir scan
func (ds *DirScanner) Exclude(excludes ...string) *DirScanner {
	ds.excludeExp = append(ds.excludeExp, excludes...)
	return ds
}

//exclude exp for dir scan
func (ds *DirScanner) Include(includes ...string) *DirScanner {
	ds.includeExp = append(ds.includeExp, includes...)
	return ds
}

//to star scan the dir.
func (ds *DirScanner) Scan() error {
	baseDir := ds.baseDir
	ds.Runtime = time.Duration(0)
	var err error = nil
	if IsDir(baseDir) {
		start := time.Now()
		ds.scanRecursion(baseDir)
		ds.AllItem = ds.AllDirItem + ds.AllFileItem
		ds.Runtime = time.Since(start)
	} else {
		err = errors.New(fmt.Sprintf("%v is not a valid dir.", baseDir))
	}
	return err
}

//recursion to scan dir, return the children count size.
func (ds *DirScanner) scanRecursion(vDir string) int64 {
	files, err := ioutil.ReadDir(vDir)
	if err != nil {
		return 0
	}
	isTopClass := false
	if ds.TopChildDick == nil {
		ds.TopChildDick = map[string]ChildDirData{}
		isTopClass = true
	}
	var currentSize int64 = 0
	for _, fl := range files {
		name := fl.Name()
		if ds.ignoreScan(name) {
			continue
		}
		var size int64
		if fl.IsDir() {
			ds.AllDirItem += 1
			size = ds.scanRecursion(StdDir(vDir + "/" + name))
			currentSize += size
		} else {
			size = fl.Size()
			currentSize += size
			ds.AllSize += fl.Size()
			ds.AllFileItem += 1
		}
		if isTopClass {
			ds.TopChildDick[name] = ChildDirData{
				Name:  name,
				Size:  size,
				IsDir: fl.IsDir(),
			}
		}
	}
	return currentSize
}

//忽略扫描
func (ds *DirScanner) ignoreScan(name string) bool {
	ignore := false
	allExp := "*"
	if len(ds.includeExp) > 0 {
		isFilter := false
		for _, filter := range ds.includeExp {
			if idx := strings.Index(name, allExp); idx > -1 {
				filter = strings.ReplaceAll(filter, allExp, ".*")
				if isMatch, er := regexp.MatchString(filter, name); er == nil && isMatch {
					isFilter = true
					break
				}
			} else if name == filter {
				isFilter = true
				break
			}
		}
		return !isFilter
	}
	if len(ds.excludeExp) > 0 {
		isExclude := false
		for _, filter := range ds.includeExp {
			if idx := strings.Index(name, allExp); idx > -1 {
				filter = strings.ReplaceAll(filter, allExp, ".*")
				if isMatch, er := regexp.MatchString(filter, name); er == nil && isMatch {
					isExclude = true
					break
				}
			} else if name == filter {
				isExclude = true
				break
			}
		}

		ignore = isExclude
	}
	return ignore
}

func NewDirScanner(vDir string) *DirScanner {
	ds := &DirScanner{}
	ds.baseDir = vDir
	return ds
}
