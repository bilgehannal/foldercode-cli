package operations

import (
	"encoding/base64"
	"fmt"
	"github.com/bilgehannal/foldercode-cli/internal/env"
	"github.com/bilgehannal/foldercode-cli/internal/pkg/args"
	"github.com/bilgehannal/foldercode-cli/internal/pkg/errors"
)

type HelpOperationBuilder struct{}

type HelpOperation struct{}

func (e HelpOperation) Run() {
	decoded, err := base64.StdEncoding.DecodeString(env.HelpContent)
	if err != nil {
		errors.FatalPanic("Error decoding base64 string:", err)
	}
	fmt.Println(string(decoded))
}

func (e HelpOperationBuilder) CreateOperation(args args.Args) Operation {
	return HelpOperation{}
}
