package fs

import (
	"fmt"
	"os"
	"path"
	"strings"
)

// InsToFilename Insert the name after the file name and before the suffix，
//
// E.G `$s => /test/pathx/name-eg.docx` into `/test/pathx/name-eg{$s}.docx`
func InsToFilename(filename, ins string) string {
	vDir, vFile := path.Split(filename)
	name := vFile
	ext := path.Ext(vFile)
	if ext != "" {
		name = strings.ReplaceAll(name, ext, "")
	}

	name = fmt.Sprintf("%s%s%s", name, ins, ext)
	return path.Join(vDir, name)
}

// InsToFilenameDetect tentatively obtaining the additional name of the inserted file
//
// e.g.
//
// `-custom => /test/pathx/name-eg.docx` into `/test/pathx/name-eg-custom.docx`
//
// `custom.docx => /test/pathx/name-eg.docx` into `/test/pathx/custom.docx`
//
// `/new/x/custom.docx => /test/pathx/name-eg.docx` into `/new/x/custom.docx`
func InsToFilenameDetect(filename, ins string) string {
	sDir, sFile := path.Split(ins)
	if sDir != "" {
		return ins
	}

	vDir, vFile := path.Split(filename)
	ext := path.Ext(vFile)

	name := vFile
	if ext != "" {
		if path.Ext(sFile) != "" {
			ext = ""
			name = ""
		} else {
			name = strings.ReplaceAll(name, ext, "")
		}
	}

	name = fmt.Sprintf("%s%s%s", name, ins, ext)
	return path.Join(vDir, name)
}

// RemoveList batch remove File List
func RemoveList(files []string) string {
	var lines []string
	for _, fl := range files {
		err := os.Remove(fl)
		if err != nil {
			lines = append(lines, fmt.Sprintf("e:%s", fl))
		} else {
			lines = append(lines, fmt.Sprintf("o:%s", fl))
		}
	}
	return strings.Join(lines, ", ")
}
