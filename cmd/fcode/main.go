package main

import (
	"github.com/bilgehannal/foldercode-cli/internal/pkg/args"
	"github.com/bilgehannal/foldercode-cli/internal/pkg/operations"
	"os"
)

func main() {
	projectArgs := args.GetArgs(os.Args[1:])
	o := operations.GetOperation(projectArgs)
	o.Run()
}
