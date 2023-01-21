package operations

import (
	"fmt"
	"github.com/bilgehannal/foldercode-cli/internal/pkg/args"
)

type EmptyOperationBuilder struct{}

type EmptyOperation struct{}

func (e EmptyOperation) Run() {
	fmt.Println("Empty Operation")
}

func (e EmptyOperationBuilder) CreateOperation(args args.Args) Operation {
	return EmptyOperation{}
}
