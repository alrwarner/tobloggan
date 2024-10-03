package stations

import (
	"os"
	"path/filepath"

	"tobloggan/code/contracts"
)

type PageWriter struct {
	targetDirectory string
	fileWriter      contracts.FSWriter
}

func NewPageWriter(directory string, writer contracts.FSWriter) *PageWriter {
	return &PageWriter{targetDirectory: directory, fileWriter: writer}
}

func (this *PageWriter) Do(input any, output func(any)) {
	//    TODO: given a contracts.Page, create a directory at contracts.Page.Path then
	//    write contracts.Page.Content to filepath.Join(this.targetDirectory, input.Path, "index.html")

	switch input := input.(type) {
	case contracts.Page:

		fPath := filepath.Join(this.targetDirectory, input.Path, "index.html")

		this.fileWriter.WriteFile(fPath, []byte(input.Content), os.ModePerm)

	default:
		output(input)
	}

}
