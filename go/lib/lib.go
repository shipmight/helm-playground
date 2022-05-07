package lib

import (
	"bytes"
	"encoding/json"
	"errors"
	"strings"
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

	t := template.New("template")

	var output bytes.Buffer

	funcMap := funcMap()

	/*
	 * Based on:
	 * https://github.com/helm/helm/blob/3d1bc72827e4edef273fb3d8d8ded2a25fa6f39d/pkg/engine/engine.go#L112-L125
	 */
	recursionMaxNums := 5
	includedNames := make(map[string]int)
	funcMap["include"] = func(name string, data interface{}) (interface{}, error) {
		var buf strings.Builder
		if v, ok := includedNames[name]; ok {
			if v > recursionMaxNums {
				return "", errors.New("rendering template has a nested reference name: " + name)
			}
			includedNames[name]++
		} else {
			includedNames[name] = 1
		}
		err := t.ExecuteTemplate(&buf, name, data)
		includedNames[name]--
		return buf.String(), err
	}

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

	t.Funcs(funcMap)

	t, err := t.Parse(templateYaml)
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
