package errors

import "log"

func FatalPanic(errorText string, err error) {
	log.Fatal(errorText, "\n", err)
	panic(err)
}

type FileNotExistError struct{}

func (e FileNotExistError) Error() string {
	return `File not exist!`
}

type VerbUnsoppertedError struct{}

func (e VerbUnsoppertedError) Error() string {
	return `The verb is not unsopperted. You can check the valid verbs using command below:
		$ fc --help`
}

type VerbMissingArgumentsError struct{}

func (e VerbMissingArgumentsError) Error() string {
	return `There are some missing arguments, There should be a verb and object
		ex: fc upload a.txt b.txt
		ex: fc get ut53
		fc [verb] [object]`
}
