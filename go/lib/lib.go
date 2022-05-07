package lib

import (
	"bytes"
	"encoding/json"
	"text/template"

	"sigs.k8s.io/yaml"
)

type ValuesObj map[string]interface{}

type TemplateData struct {
	Values ValuesObj
}

type GetYamlReturnValue struct {
	Yaml string `json:"yaml"`
	Err  string `json:"err"`
}

func toJson(returnValue GetYamlReturnValue) string {
	bytes, err := json.Marshal(returnValue)
	if err != nil {
		return `{"err":"conversion to JSON failed"}`
	}
	return string(bytes)
}

func GetYaml(templateYaml string, valuesYaml string) string {
	valuesData := ValuesObj{}
	if err := yaml.Unmarshal([]byte(valuesYaml), &valuesData); err != nil {
		return toJson(GetYamlReturnValue{
			Err: err.Error(),
		})
	}

	templateData := TemplateData{valuesData}

	var output bytes.Buffer

	funcMap := funcMap()

	// If the template contains `required`, we don't want to return an error which
	// would prevent previewing the entire template. Simply pass the value through.
	funcMap["required"] = func(val interface{}) (interface{}, error) {
		return val, nil
	}

	// If the template contains `fail`, we don't want to return an error which
	// would prevent previewing the entire template. Return an empty string.
	funcMap["fail"] = func(val interface{}) (interface{}, error) {
		return "", nil
	}

	t, err := template.New("template").Funcs(funcMap).Parse(templateYaml)
	if err != nil {
		return toJson(GetYamlReturnValue{
			Err: err.Error(),
		})
	}

	if err := t.Execute(&output, templateData); err != nil {
		return toJson(GetYamlReturnValue{
			Err: err.Error(),
		})
	}

	return toJson(GetYamlReturnValue{
		Yaml: output.String(),
	})
}
