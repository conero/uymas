package fs

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"
)

const (
	DirScannerChanMax = 5000
)

type ChildDirData struct {
	Name  string
	Size  int64
	IsDir bool
	Path  string
	Depth int // the max depth of dir. @todo need to do
}

// DirScanner the tool to scan the dirs.
type DirScanner struct {
	baseDir       string
	AllItem       int
	AllDirItem    int
	AllFileItem   int
	AllSize       int64
	TopChildDick  map[string]ChildDirData
	Runtime       time.Duration
	cddWait       sync.WaitGroup    // [Experimental]
	cddChan       chan ChildDirData // [Experimental]
	cddChanDone   chan bool         // [Experimental]
	CddChanMax    int               // [Experimental]
	iCddChanCount int               // [Experimental]

	//struct inner/private valuable
	//排除名称，以"*"匹配（不包含）
	excludeExp []string
	//过滤表达式（包含）
	includeExp   []string
	filterNameMK bool
}

// Exclude exclude exp for dir scan
func (ds *DirScanner) Exclude(excludes ...string) *DirScanner {
	var newExcludes []string
	for _, ecld := range excludes {
		if "" == strings.TrimSpace(ecld) {
			continue
		}
		newExcludes = append(newExcludes, ecld)
	}
	if len(newExcludes) == 0 {
		return ds
	}
	ds.excludeExp = append(ds.excludeExp, newExcludes...)
	if !ds.filterNameMK {
		ds.filterNameMK = len(ds.excludeExp) > 0
	}
	return ds
}

// Include exclude exp for dir scan
func (ds *DirScanner) Include(includes ...string) *DirScanner {
	var newInclude []string
	for _, icld := range includes {
		if "" == strings.TrimSpace(icld) {
			continue
		}
		newInclude = append(newInclude, icld)
	}
	if len(newInclude) == 0 {
		return ds
	}
	ds.includeExp = append(ds.includeExp, newInclude...)
	if !ds.filterNameMK {
		ds.filterNameMK = len(ds.includeExp) > 0
	}
	return ds
}

// Scan to star scan the dir.
func (ds *DirScanner) Scan() error {
	baseDir := ds.baseDir
	ds.Runtime = time.Duration(0)
	var err error = nil
	if IsDir(baseDir) {
		start := time.Now()
		ds.scanRecursion(baseDir, 0)
		ds.AllItem = ds.AllDirItem + ds.AllFileItem
		ds.Runtime = time.Since(start)
	} else {
		err = fmt.Errorf("%v is not a valid dir", baseDir)
	}
	return err
}

// recursion to scan dir, return the children count size.
func (ds *DirScanner) scanRecursion(vDir string, depth int) int64 {
	dirEntries, err := os.ReadDir(vDir)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	isTopClass := false
	if ds.TopChildDick == nil {
		ds.TopChildDick = map[string]ChildDirData{}
		isTopClass = true
	}
	var currentSize int64 = 0
	for _, entry := range dirEntries {
		name := entry.Name()
		vPath := StdPathName(fmt.Sprintf("%v/%v", vDir, name))
		var size int64
		if entry.IsDir() {
			ds.AllDirItem += 1
			depth += 1
			size = ds.scanRecursion(vPath, depth)
			currentSize += size
		} else {
			if ds.ignoreScan(name) {
				continue
			}
			if fi, fErr := entry.Info(); fErr == nil {
				size = fi.Size()
				ds.AllSize += fi.Size()
			}
			currentSize += size
			ds.AllFileItem += 1
		}
		if isTopClass {
			// filter zero top dir when have a filter rule
			if ds.filterNameMK && size == 0 {
				continue
			}
			ds.TopChildDick[name] = ChildDirData{
				Name:  name,
				Size:  size,
				IsDir: entry.IsDir(),
				Depth: depth,
			}
		}
	}
	return currentSize
}

// ScanParallel to star scan the dir.[Experimental]
func (ds *DirScanner) ScanParallel() error {
	baseDir := ds.baseDir
	ds.Runtime = time.Duration(0)
	var err error = nil
	if IsDir(baseDir) {
		start := time.Now()
		//default channel cache
		if ds.CddChanMax < 1 {
			ds.CddChanMax = DirScannerChanMax
		}

		ds.cddChan = make(chan ChildDirData, ds.CddChanMax)
		ds.cddChanDone = make(chan bool)

		// sync read the channel
		go func() {
			isDown := false
			var sizeCount int64
			for {
				select {
				case cc := <-ds.cddChan:
					sizeCount += cc.Size
				case <-ds.cddChanDone:
					isDown = true
				default:
				}
				if isDown {
					break
				}
			}

			//add the parallel scan
			ds.AllSize += sizeCount
		}()

		ds.scanRecursionParallel(baseDir, 0)
		// read the chan
		//for cc :=<-ds.cddChan; cc.Name != "";{
		//	fmt.Printf("%#v\r\n", cc)
		//}

		//wait all goroutine is done
		ds.cddWait.Wait()
		ds.cddChanDone <- true

		ds.AllItem = ds.AllDirItem + ds.AllFileItem
		ds.Runtime = time.Since(start)
	} else {
		err = fmt.Errorf("%v is not a valid dir", baseDir)
	}
	return err
}

// recursion to scan dir, return the children count size.[Experimental]
func (ds *DirScanner) scanRecursionParallel(vDir string, depth int) int64 {
	dirEntries, err := os.ReadDir(vDir)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	isTopClass := false
	if ds.TopChildDick == nil {
		ds.TopChildDick = map[string]ChildDirData{}
		isTopClass = true
	}
	var currentSize int64 = 0
	for _, entry := range dirEntries {
		name := entry.Name()
		vPath := StdPathName(fmt.Sprintf("%v/%v", vDir, name))
		var size int64
		if entry.IsDir() {
			ds.AllDirItem += 1
			depth += 1
			// if goroutine is bigger than `ds.CddChanMax` will stop to distributive it.
			if ds.iCddChanCount > ds.CddChanMax {
				size = ds.scanRecursionParallel(vPath, depth)
			} else {
				ds.cddWait.Add(1)
				ds.iCddChanCount += 1
				go func(vp string, cdd chan ChildDirData) {
					defer ds.cddWait.Done()
					chSize := ds.scanRecursionParallel(vPath, depth)
					cdd <- ChildDirData{
						Name:  name,
						Size:  chSize,
						IsDir: true,
						Path:  vPath,
						Depth: depth,
					}
				}(vPath, ds.cddChan)
			}
			currentSize += size
		} else {
			if ds.ignoreScan(name) {
				continue
			}

			if fi, fEr := entry.Info(); fEr == nil {
				size = fi.Size()
				ds.AllSize += fi.Size()
			}

			currentSize += size
			ds.AllFileItem += 1
		}
		if isTopClass {
			// filter zero top dir when have a filter rule
			if ds.filterNameMK && size == 0 {
				continue
			}
			ds.TopChildDick[name] = ChildDirData{
				Name:  name,
				Size:  size,
				IsDir: entry.IsDir(),
				Depth: depth,
			}
		}
	}
	return currentSize
}

// ignore scan target name
func (ds *DirScanner) ignoreScan(name string) bool {
	ignore := false
	allExp := "*"
	if len(ds.includeExp) > 0 {
		isFilter := false
		for _, filter := range ds.includeExp {
			if "" == strings.TrimSpace(filter) {
				continue
			}
			if idx := strings.Index(filter, allExp); idx > -1 {
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
			if "" == strings.TrimSpace(filter) {
				continue
			}
			if idx := strings.Index(filter, allExp); idx > -1 {
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

func (ds *DirScanner) BaseDir() string {
	return ds.baseDir
}

func (ds *DirScanner) ChanNumber() int {
	return ds.iCddChanCount
}

func NewDirScanner(vDir string) *DirScanner {
	ds := &DirScanner{}
	ds.baseDir = StdDir(vDir)
	return ds
}
