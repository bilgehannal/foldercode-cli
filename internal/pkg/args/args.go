package args

import (
	"github.com/bilgehannal/foldercode-cli/internal/pkg/errors"
	"github.com/bilgehannal/foldercode-cli/internal/pkg/utils"
)

type Args struct {
	Verb    Verb
	Objects []Object
}

type Verb struct {
	Value string
}

type Object struct {
	Value string
}

func validateVerb(verb Verb) (Verb, error) {
	if !utils.StringSliceContains(GetAvailableVerbValues(), verb.Value) {
		return verb, errors.VerbUnsoppertedError{}
	}
	return verb, nil
}

func validateObject(object Object) (Object, error) {
	return object, nil
}
