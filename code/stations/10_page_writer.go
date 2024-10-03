package stations

import (
	"bufio"
	"os"
	"path/filepath"

	"tobloggan/code/contracts"
)

type PageWriter struct {
	targetDirectory string
}

func (this *PageWriter) Do(input any, output func(any)) {
	//    TODO: given a contracts.Page, create a directory at contracts.Page.Path then
	//    write contracts.Page.Content to filepath.Join(this.targetDirectory, input.Path, "index.html")

	switch input := input.(type) {
	case contracts.Page:
		fPath := filepath.Join(this.targetDirectory, input.Path, "index.html")
		os.MkdirAll(fPath, os.ModePerm)
		outputFile, _ := os.Create(fPath)
		writer := bufio.NewWriter(outputFile)
		writer.WriteString(input.Content)
		writer.Flush()
		outputFile.Close()
		output(input)
	default:
		output(input)
	}

}
