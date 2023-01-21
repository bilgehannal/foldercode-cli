package main

import (
	"fmt"
	"github.com/bilgehannal/foldercode-cli/internal/pkg/args"
	"github.com/bilgehannal/foldercode-cli/internal/pkg/operations"
	"os"
)

func main() {
	projectArgs := args.GetArgs(os.Args[1:])
	fmt.Println(projectArgs)
	o := operations.GetOperation(projectArgs)
	o.Run()
}
