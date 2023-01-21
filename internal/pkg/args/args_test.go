package args

import "testing"

func Test_validateVerb(t *testing.T) {
	testDataSuccess := []Verb{
		Verb{Value: "upload"},
		Verb{Value: "get"},
		Verb{Value: "--help"},
	}
	testDataErr := []Verb{
		Verb{Value: "uploadd"},
		Verb{},
	}
	for _, verb := range testDataSuccess {
		_, err := validateVerb(verb)
		if err != nil {
			t.Error(verb)
		}
	}
	for _, verb := range testDataErr {
		_, err := validateVerb(verb)
		if err == nil {
			t.Error(verb)
		}
	}
}

func Test_validateObject(t *testing.T) {
	testDataSuccess := []Object{
		Object{Value: "x"},
		Object{Value: "y"},
		Object{Value: "z"},
	}
	for _, object := range testDataSuccess {
		_, err := validateObject(object)
		if err != nil {
			t.Error(object)
		}
	}
}
