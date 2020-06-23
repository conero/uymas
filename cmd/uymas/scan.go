package main

import (
	"github.com/conero/uymas/fs"
	"io/ioutil"
)

//the scan the file system for to know.

type ScanDirData struct {
	AllSize      int64
	ChildrenDick map[string]int64
}

//@todo try use the struct to do this, and gite more data.
func ScanDir(vDir string) ScanDirData {
	if vDir == "" {
		vDir = "./"
	}
	var sdd ScanDirData = ScanDirData{}
	files, err := ioutil.ReadDir(vDir)
	if err != nil {
		return sdd
	}
	var sizeBytesCount int64
	var childrenSiteDick map[string]int64 = map[string]int64{}
	for _, fl := range files {
		name := fl.Name()
		if fl.IsDir() {
			childrenSiteDick[name] = readDirAllSize(fs.StdDir(vDir + "/" + fl.Name()))
			sizeBytesCount += childrenSiteDick[name]
		} else {
			sizeBytesCount += fl.Size()
			childrenSiteDick[name] = fl.Size()
		}
	}

	sdd.AllSize = sizeBytesCount
	sdd.ChildrenDick = childrenSiteDick
	return sdd
}

func readDirAllSize(vDir string) int64 {
	var sizeBytesCount int64
	files, err := ioutil.ReadDir(vDir)
	if err != nil {
		return sizeBytesCount
	}
	var childrenSiteDick map[string]int64 = map[string]int64{}
	for _, fl := range files {
		name := fl.Name()
		if fl.IsDir() {
			childrenSiteDick[name] = readDirAllSize(fs.StdDir(vDir + "/" + fl.Name()))
			sizeBytesCount += childrenSiteDick[name]
		} else {
			sizeBytesCount += fl.Size()
			childrenSiteDick[name] = fl.Size()
		}
	}
	return sizeBytesCount
}
