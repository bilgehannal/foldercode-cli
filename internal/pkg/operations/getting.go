package operations

import (
	"fmt"
	"github.com/bilgehannal/foldercode-cli/internal/pkg/args"
)

type GettingOperationBuilder struct{}

type GettingOperation struct {
	code args.Object
}

func (g GettingOperation) Run() {
	fmt.Println("Upload Operation")
	fmt.Println(g.code)
}

func (g GettingOperationBuilder) CreateOperation(args args.Args) Operation {
	return GettingOperation{
		code: args.Objects[0],
	}
}
