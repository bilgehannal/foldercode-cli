package args

import "testing"

func Test_GetAvailableVerbValues(t *testing.T) {
	testValues := []string{"upload", "get", "--help"}
	values := GetAvailableVerbValues()

	if len(testValues) != len(values) {
		t.Errorf("length is not matched!")
	}
	for i := 0; i < len(values); i++ {
		if values[i] != testValues[i] {
			t.Errorf("elements are not matched!")
		}
	}
}
