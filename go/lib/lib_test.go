package lib

import (
	"testing"
)

func TestGetYaml(t *testing.T) {
	tests := []struct {
		templateYaml string
		valuesYaml   string
		expectedYaml string
	}{
		{
			templateYaml: "",
			valuesYaml:   "",
			expectedYaml: "",
		},
	}

	for _, tc := range tests {
		yaml, err := GetYaml(tc.templateYaml, tc.valuesYaml)
		if err != nil {
			t.Errorf("error: %w", err)
		}
		if yaml != tc.expectedYaml {
			t.Errorf("handler returned wrong status code: got %v want %v", yaml, tc.expectedYaml)
		}
	}
}
