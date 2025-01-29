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
			expectedReturnValue: `{"yaml":"","err":"","warning":""}`,
		},

		// Values
		{
			templateYaml:        "name: {{ .Values.foobar }}",
			valuesYaml:          "foobar: hello",
			expectedReturnValue: `{"yaml":"name: hello","err":"","warning":""}`,
		},
		{
			templateYaml:        "name: {{ .Values.foobar }}",
			valuesYaml:          "",
			expectedReturnValue: `{"yaml":"name: \u003cno value\u003e","err":"","warning":""}`,
		},

		// Built-in objects
		{
			templateYaml:        "name: {{ .Release | toYaml | nindent 2 }}",
			valuesYaml:          "",
			expectedReturnValue: `{"yaml":"name: \n  IsInstall: false\n  IsUpgrade: false\n  Name: example\n  Namespace: example\n  Revision: 1\n  Service: Helm","err":"","warning":""}`,
		},
		{
			templateYaml:        "name: {{ .Chart | toYaml | nindent 2 }}",
			valuesYaml:          "",
			expectedReturnValue: `{"yaml":"name: \n  Annotations: {}\n  ApiVersion: v2\n  AppVersion: \"\"\n  Dependencies: []\n  Deprecated: false\n  Description: example\n  Home: \"\"\n  Icon: \"\"\n  Keywords: []\n  KubeVersion: \"\"\n  Maintainers: []\n  Name: example\n  Sources: []\n  Type: application\n  Version: 0.1.0","err":"","warning":""}`,
		},

		// Template functions
		{
			templateYaml:        "name: {{ .Values.foobar | default \"fallback\" }}",
			valuesYaml:          "",
			expectedReturnValue: `{"yaml":"name: fallback","err":"","warning":""}`,
		},

		// Helm built-in template functions
		{
			templateYaml:        "name: {{ .Values.foobar | toYaml | nindent 2 }}",
			valuesYaml:          "foobar:\n  first: 1\n  second: 2",
			expectedReturnValue: `{"yaml":"name: \n  first: 1\n  second: 2","err":"","warning":""}`,
		},
		{
			templateYaml:        "name: {{ .Values.foobar | required \".foobar must be set\" }}",
			valuesYaml:          "",
			expectedReturnValue: `{"yaml":"name: \u003cno value\u003e","err":"","warning":""}`,
		},
		{
			templateYaml:        "{{ fail \"reason\" }}",
			valuesYaml:          "",
			expectedReturnValue: `{"yaml":"","err":"","warning":""}`,
		},
		{
			templateYaml: `
				{{- define "foobar.name" -}}
					{{- . | trunc 5 -}}
				{{- end -}}
				fromTemplate: {{ include "foobar.name" .Values.name -}}
			`,
			valuesYaml:          "name: '123456789'",
			expectedReturnValue: `{"yaml":"fromTemplate: 12345","err":"","warning":""}`,
		},
		{
			templateYaml: `
				{{- define "foobar.name" -}}
				  {{- include "foobar.name" . -}}
				{{- end -}}
				fromTemplate: {{ include "foobar.name" .Values.name -}}
			`,
			valuesYaml:          "name: '123456789'",
			expectedReturnValue: `{"yaml":"","err":"template: template:5:21: executing \"template\" at \u003cinclude \"foobar.name\" .Values.name\u003e: error calling include: template: template:3:10: executing \"foobar.name\" at \u003cinclude \"foobar.name\" .\u003e: error calling include: template: template:3:10: executing \"foobar.name\" at \u003cinclude \"foobar.name\" .\u003e: error calling include: template: template:3:10: executing \"foobar.name\" at \u003cinclude \"foobar.name\" .\u003e: error calling include: template: template:3:10: executing \"foobar.name\" at \u003cinclude \"foobar.name\" .\u003e: error calling include: template: template:3:10: executing \"foobar.name\" at \u003cinclude \"foobar.name\" .\u003e: error calling include: template: template:3:10: executing \"foobar.name\" at \u003cinclude \"foobar.name\" .\u003e: error calling include: rendering template has a nested reference name: foobar.name","warning":""}`,
		},
		{
			templateYaml:        "tplname: {{ tpl .Values.tpltest . }}",
			valuesYaml:          "name: '123456789'\ntpltest: '{{ .Values.name }}'",
			expectedReturnValue: `{"yaml":"tplname: 123456789","err":"","warning":""}`,
		},
		{
			templateYaml:        "tplname: {{ tpl .Values.tpltest . }}",
			valuesYaml:          "tpltest: '{{ tpl .Values.tpltest . }}'",
			expectedReturnValue: `{"yaml":"","err":"template: template:1:12: executing \"template\" at \u003ctpl .Values.tpltest .\u003e: error calling tpl: template: tpl:1:3: executing \"tpl\" at \u003ctpl .Values.tpltest .\u003e: error calling tpl: template: tpl:1:3: executing \"tpl\" at \u003ctpl .Values.tpltest .\u003e: error calling tpl: tpl has been called 3 times, aborting to prevent infinite loops","warning":""}`,
		},
		{
			templateYaml:        "tplname: {{ tpl .Values.tpltest1 . }}",
			valuesYaml:          "tpltest1: '{{ tpl .Values.tpltest2 . }}'\ntpltest2: '{{ tpl .Values.tpltest1 . }}'",
			expectedReturnValue: `{"yaml":"","err":"template: template:1:12: executing \"template\" at \u003ctpl .Values.tpltest1 .\u003e: error calling tpl: template: tpl:1:3: executing \"tpl\" at \u003ctpl .Values.tpltest2 .\u003e: error calling tpl: template: tpl:1:3: executing \"tpl\" at \u003ctpl .Values.tpltest1 .\u003e: error calling tpl: tpl has been called 3 times, aborting to prevent infinite loops","warning":""}`,
		},
		// Template formatting error
		{
			templateYaml:        "name: {{ .Values. }}",
			valuesYaml:          "",
			expectedReturnValue: `{"yaml":"","err":"template: template:1: unexpected \u003c.\u003e in operand","warning":""}`,
		},
		{
			templateYaml:        "\n\nname: {{ .Values. }}",
			valuesYaml:          "",
			expectedReturnValue: `{"yaml":"","err":"template: template:3: unexpected \u003c.\u003e in operand","warning":""}`,
		},
		{
			templateYaml:        "\nname: {{ .Values.foobar | what }}",
			valuesYaml:          "",
			expectedReturnValue: `{"yaml":"","err":"template: template:2: function \"what\" not defined","warning":""}`,
		},

		// Warning about malformed yaml
		{
			templateYaml:        "first: truesecond:\n  third: true",
			valuesYaml:          "",
			expectedReturnValue: `{"yaml":"first: truesecond:\n  third: true","err":"","warning":"error converting YAML to JSON: yaml: mapping values are not allowed in this context"}`,
		},
	}

	for _, tc := range tests {
		returnValue := GetYaml(tc.templateYaml, tc.valuesYaml, GetYamlConfig{MaxTplRuns: 3})
		if returnValue != tc.expectedReturnValue {
			t.Errorf("\n\ntemplateYaml: %v\n\nvaluesYaml: %v\n\nexpected: %v\n\nactual: %v\n\n", tc.templateYaml, tc.valuesYaml, tc.expectedReturnValue, returnValue)
		}
	}
}
