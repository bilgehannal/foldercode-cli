package operations

import "github.com/bilgehannal/foldercode-cli/internal/pkg/args"

type OperationBuilder interface {
	CreateOperation(args.Args) Operation
}

type Operation interface {
	Run()
}

func GetOperationBuilderMap() map[string]OperationBuilder {
	return map[string]OperationBuilder{
		args.VerbUpload: UploadOperationBuilder{},
		args.VerbGet:    GettingOperationBuilder{},
		args.VerbHelp:   HelpOperationBuilder{},
	}
}

func GetOperationBuilder(args args.Args) OperationBuilder {
	operationBuilders := GetOperationBuilderMap()
	if val, ok := operationBuilders[args.Verb.Value]; ok {
		return val
	}
	return EmptyOperationBuilder{}
}

func GetOperation(args args.Args) Operation {
	return GetOperationBuilder(args).CreateOperation(args)
}
