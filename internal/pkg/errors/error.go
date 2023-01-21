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

type CodeMissingError struct{}

func (e CodeMissingError) Error() string {
	return `Code could not be get, most probably code is missing!`
}

type FileMissingError struct{}

func (e FileMissingError) Error() string {
	return `Files could not be get, most probably Files parameters are missing!`
}

type DownloadError struct{}

func (e DownloadError) Error() string {
	return `Downloading is not succesfull!`
}

type VerbMissingArgumentsError struct{}

func (e VerbMissingArgumentsError) Error() string {
	return `There are some missing arguments, There should be a verb and object
		ex: fc upload a.txt b.txt
		ex: fc get ut53
		fc [verb] [object]`
}
