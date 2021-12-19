package lib

import (
	"testing"
)

func TestGetYaml(t *testing.T) {
	tests := []struct {
		templateYaml        string
		valuesYaml          string
		expectedReturnValue string
	}{
		{
			templateYaml:        "",
			valuesYaml:          "",
			expectedReturnValue: `{"yaml":"","err":""}`,
		},
		{
			templateYaml:        "name: {{ .Values.foobar }}",
			valuesYaml:          "foobar: hello",
			expectedReturnValue: `{"yaml":"name: hello","err":""}`,
		},
		{
			templateYaml:        "name: {{ .Values.foobar | default \"fallback\" }}",
			valuesYaml:          "",
			expectedReturnValue: `{"yaml":"name: fallback","err":""}`,
		},
		{
			templateYaml:        "name: {{ .Values.foobar | toYaml }}",
			valuesYaml:          "foobar: ['first', 'second']",
			expectedReturnValue: `{"yaml":"name: - first\n- second","err":""}`,
		},
		{
			templateYaml:        "name: {{ .Values. }}",
			valuesYaml:          "",
			expectedReturnValue: `{"yaml":"","err":"template: template:1: unexpected \u003c.\u003e in operand"}`,
		},
		{
			templateYaml:        "\n\nname: {{ .Values. }}",
			valuesYaml:          "",
			expectedReturnValue: `{"yaml":"","err":"template: template:3: unexpected \u003c.\u003e in operand"}`,
		},
	}

	for _, tc := range tests {
		returnValue := GetYaml(tc.templateYaml, tc.valuesYaml)
		if returnValue != tc.expectedReturnValue {
			t.Errorf("\n\ntemplateYaml: %v\n\nvaluesYaml: %v\n\nexpected: %v\n\nactual: %v\n\n", tc.templateYaml, tc.valuesYaml, tc.expectedReturnValue, returnValue)
		}
	}
}
