package stations

import (
	"io/fs"
	"strings"

	"tobloggan/code/contracts"
)

type SourceScanner struct {
	filesystem fs.FS
}

func NewSourceScanner(filesystem fs.FS) contracts.Station {
	return &SourceScanner{filesystem: filesystem}
}

func (this *SourceScanner) Do(input any, output func(any)) {
	switch typed := input.(type) {
	case contracts.SourceDirectory:
		this.discoverFiles(typed, output)
	default:
		output(input)
	}
}

func (this *SourceScanner) discoverFiles(typed contracts.SourceDirectory, output func(any)) {
	_ = fs.WalkDir(this.filesystem, string(typed), func(filename string, directory fs.DirEntry, err error) error {
		if err != nil {
			output(err)
		} else if isArticle(filename, directory) {
			output(contracts.SourceFilePath(filename))
		}
		return nil
	})
}

func isArticle(filename string, directory fs.DirEntry) bool {
	if directory.IsDir() {
		return false
	}
	if !strings.HasSuffix(filename, ".md") {
		return false
	}
	return true
}
