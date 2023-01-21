package operations

import (
	"fmt"
	"github.com/bilgehannal/foldercode-cli/internal/pkg/args"
	"github.com/bilgehannal/foldercode-cli/internal/pkg/errors"
	"github.com/bilgehannal/foldercode-cli/internal/pkg/foldercode"
	"os"
)

type UploadOperationBuilder struct{}

type UploadOperation struct {
	files []args.Object
}

func validateFiles(files []args.Object) {
	for _, file := range files {
		if _, err := os.Stat(file.Value); os.IsNotExist(err) {
			errors.FatalPanic(file.Value, errors.FileNotExistError{})
		}
	}
}

func uploadFiles(files []args.Object) {
	fc := foldercode.FoldercodeClient{}
	session, err := fc.GenerateSession()
	if err != nil {
		errors.FatalPanic("Upload File Error", err)
	}
	for _, file := range files {
		err = fc.UploadFile(session, file.Value)
		if err != nil {
			errors.FatalPanic("Upload File Error", err)
		}
	}
	fmt.Println(session.FolderCode)
}

func (u UploadOperation) Run() {
	validateFiles(u.files)
	uploadFiles(u.files)
}

func (u UploadOperationBuilder) CreateOperation(args args.Args) Operation {
	if !(len(args.Objects) > 0) {
		errors.FatalPanic("Code Error", errors.FileMissingError{})
	}
	return UploadOperation{
		files: args.Objects,
	}
}
