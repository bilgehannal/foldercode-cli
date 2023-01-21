package operations

import (
	"fmt"
	"github.com/bilgehannal/foldercode-cli/internal/pkg/args"
)

type UploadOperationBuilder struct{}

type UploadOperation struct {
	files []args.Object
}

func (u UploadOperation) Run() {
	fmt.Println("Upload Operation")
	fmt.Println(u.files)
}

func (u UploadOperationBuilder) CreateOperation(args args.Args) Operation {
	return UploadOperation{
		files: args.Objects,
	}
}
