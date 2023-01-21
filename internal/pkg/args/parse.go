package args

import (
	"github.com/bilgehannal/foldercode-cli/internal/pkg/errors"
	"log"
)

const (
	orderOfVerbValue   = 0
	orderOfObjectValue = 1
)

func getVerb(args []string) (Verb, error) {
	if len(args) < 2 {
		return Verb{}, errors.VerbMissingArgumentsError{}
	}
	currentVerb := Verb{Value: args[orderOfVerbValue]}
	return validateVerb(currentVerb)
}

func addObjectToList(objects []Object, element string) ([]Object, error) {
	object := Object{Value: element}
	object, err := validateObject(object)
	if err != nil {
		return objects, err
	}
	objects = append(objects, object)
	return objects, nil
}

func getObjects(args []string) ([]Object, error) {
	if len(args) < 2 {
		return []Object{}, errors.VerbMissingArgumentsError{}
	}
	var objects []Object
	var err error
	for i := 1; i < len(args); i++ {
		objects, err = addObjectToList(objects, args[i])
		if err != nil {
			return objects, err
		}
	}
	return objects, nil
}

func GetArgs(args []string) Args {
	verb, err := getVerb(args)
	if err != nil {
		log.Fatal("Error during getting verb: ", err)
	}
	objects, err := getObjects(args)
	if err != nil {
		log.Fatal("Error during getting files: ", err)
	}
	return Args{
		Verb:    verb,
		Objects: objects,
	}
}
