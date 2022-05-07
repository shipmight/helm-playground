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

		// Values
		{
			templateYaml:        "name: {{ .Values.foobar }}",
			valuesYaml:          "foobar: hello",
			expectedReturnValue: `{"yaml":"name: hello","err":""}`,
		},
		{
			templateYaml:        "name: {{ .Values.foobar }}",
			valuesYaml:          "",
			expectedReturnValue: `{"yaml":"name: \u003cno value\u003e","err":""}`,
		},

		// Built-in objects
		{
			templateYaml:        "name: {{ .Release | toYaml }}",
			valuesYaml:          "",
			expectedReturnValue: `{"yaml":"name: IsInstall: false\nIsUpgrade: false\nName: example\nNamespace: example\nRevision: 1\nService: Helm","err":""}`,
		},
		{
			templateYaml:        "name: {{ .Chart | toYaml }}",
			valuesYaml:          "",
			expectedReturnValue: `{"yaml":"name: Annotations: {}\nApiVersion: v2\nAppVersion: \"\"\nDependencies: []\nDeprecated: false\nDescription: example\nHome: \"\"\nIcon: \"\"\nKeywords: []\nKubeVersion: \"\"\nMaintainers: []\nName: example\nSources: []\nType: application\nVersion: 0.1.0","err":""}`,
		},

		// Template functions
		{
			templateYaml:        "name: {{ .Values.foobar | default \"fallback\" }}",
			valuesYaml:          "",
			expectedReturnValue: `{"yaml":"name: fallback","err":""}`,
		},

		// Helm built-in template functions
		{
			templateYaml:        "name: {{ .Values.foobar | toYaml }}",
			valuesYaml:          "foobar: ['first', 'second']",
			expectedReturnValue: `{"yaml":"name: - first\n- second","err":""}`,
		},
		{
			templateYaml:        "name: {{ .Values.foobar | toYaml | nindent 2 }}",
			valuesYaml:          "foobar:\n  first: 1\n  second: 2",
			expectedReturnValue: `{"yaml":"name: \n  first: 1\n  second: 2","err":""}`,
		},
		{
			templateYaml:        "name: {{ .Values.foobar | required }}",
			valuesYaml:          "",
			expectedReturnValue: `{"yaml":"name: \u003cno value\u003e","err":""}`,
		},
		{
			templateYaml:        "{{ fail \"reason\" }}",
			valuesYaml:          "",
			expectedReturnValue: `{"yaml":"","err":""}`,
		},
		{
			templateYaml: `
				{{- define "foobar.name" -}}
					{{- . | trunc 5 -}}
				{{- end -}}
				fromTemplate: {{ include "foobar.name" .Values.name -}}
			`,
			valuesYaml:          "name: '123456789'",
			expectedReturnValue: `{"yaml":"fromTemplate: 12345","err":""}`,
		},
		{
			templateYaml: `
				{{- define "foobar.name" -}}
				  {{- include "foobar.name" . -}}
				{{- end -}}
				fromTemplate: {{ include "foobar.name" .Values.name -}}
			`,
			valuesYaml:          "name: '123456789'",
			expectedReturnValue: `{"yaml":"","err":"template: template:5:21: executing \"template\" at \u003cinclude \"foobar.name\" .Values.name\u003e: error calling include: template: template:3:10: executing \"foobar.name\" at \u003cinclude \"foobar.name\" .\u003e: error calling include: template: template:3:10: executing \"foobar.name\" at \u003cinclude \"foobar.name\" .\u003e: error calling include: template: template:3:10: executing \"foobar.name\" at \u003cinclude \"foobar.name\" .\u003e: error calling include: template: template:3:10: executing \"foobar.name\" at \u003cinclude \"foobar.name\" .\u003e: error calling include: template: template:3:10: executing \"foobar.name\" at \u003cinclude \"foobar.name\" .\u003e: error calling include: template: template:3:10: executing \"foobar.name\" at \u003cinclude \"foobar.name\" .\u003e: error calling include: rendering template has a nested reference name: foobar.name"}`,
		},

		// Template formatting error
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
		{
			templateYaml:        "\nname: {{ .Values.foobar | what }}",
			valuesYaml:          "",
			expectedReturnValue: `{"yaml":"","err":"template: template:2: function \"what\" not defined"}`,
		},
	}

	for _, tc := range tests {
		returnValue := GetYaml(tc.templateYaml, tc.valuesYaml)
		if returnValue != tc.expectedReturnValue {
			t.Errorf("\n\ntemplateYaml: %v\n\nvaluesYaml: %v\n\nexpected: %v\n\nactual: %v\n\n", tc.templateYaml, tc.valuesYaml, tc.expectedReturnValue, returnValue)
		}
	}
}
