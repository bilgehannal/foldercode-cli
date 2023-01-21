package utils

import "testing"

type TestStructStringSliceContains struct {
	slice          []string
	element        string
	expectedResult bool
}

func Test_StringSliceContains(t *testing.T) {
	var testData []TestStructStringSliceContains
	testData = append(testData, TestStructStringSliceContains{
		slice:   []string{"test1", "test2", "test3"},
		element: "test1", expectedResult: true})
	testData = append(testData, TestStructStringSliceContains{
		slice:   []string{"test1", "test2", "test3"},
		element: "test4", expectedResult: false})

	for _, test := range testData {
		if StringSliceContains(test.slice, test.element) != test.expectedResult {
			t.Error(test)
		}
	}
}
