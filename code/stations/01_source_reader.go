package stations

//type SourceReader struct {
//}

//func (this *SourceReader) Do(input any, output func(v any)) {
//    TODO: given a contracts.SourceFilePath, read its contents and emit contracts.SourceFile
//    input: contracts.SourceFilePath
//    output: contracts.SourceFile, or error from fs.ReadFile, or input w/ unrecognized type
//}

import (
	"io/fs"

	"tobloggan/code/contracts"
)

type SourceReader struct {
	filesystem fs.FS
}

func NewSourceReader(filesystem fs.FS) contracts.Station {
	return &SourceReader{filesystem: filesystem}
}

func (this *SourceReader) Do(input any, output func(any)) {
	switch typed := input.(type) {
	case contracts.SourceFilePath:
		this.readFile(typed, output)
	default:
		output(input)
	}
}

func (this *SourceReader) readFile(filename contracts.SourceFilePath, output func(any)) {
	contents, err := fs.ReadFile(this.filesystem, string(filename))
	if err != nil {
		output(err)
	} else {
		output(contents)
	}
}
