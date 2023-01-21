package operations

import (
	"github.com/bilgehannal/foldercode-cli/internal/pkg/args"
	"github.com/bilgehannal/foldercode-cli/internal/pkg/errors"
	"github.com/bilgehannal/foldercode-cli/internal/pkg/foldercode"
)

type GettingOperationBuilder struct{}

type GettingOperation struct {
	code args.Object
}

func getFiles(code string) []foldercode.File {
	fc := foldercode.FoldercodeClient{}
	files, err := fc.GetFiles(code)
	if err != nil {
		errors.FatalPanic("Getting files error", err)
	}
	return files
}

func downloadFiles(files []foldercode.File) {
	fc := foldercode.FoldercodeClient{}
	for _, file := range files {
		err := fc.DownloadFile(file)
		if err != nil {
			errors.FatalPanic("Downloading files error", err)
		}
	}
}

func (g GettingOperation) Run() {
	files := getFiles(g.code.Value)
	downloadFiles(files)
}

func (g GettingOperationBuilder) CreateOperation(args args.Args) Operation {
	if !(len(args.Objects) > 0) {
		errors.FatalPanic("Code Error", errors.CodeMissingError{})
	}
	return GettingOperation{
		code: args.Objects[0],
	}
}
