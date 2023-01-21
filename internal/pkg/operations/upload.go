package operations

import (
	"fmt"
	"github.com/bilgehannal/foldercode-cli/internal/pkg/args"
	"github.com/bilgehannal/foldercode-cli/internal/pkg/errors"
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

func (u UploadOperation) Run() {
	validateFiles(u.files)
	fmt.Println("Upload Operation")
	fmt.Println(u.files)
}

func (u UploadOperationBuilder) CreateOperation(args args.Args) Operation {
	return UploadOperation{
		files: args.Objects,
	}
}
