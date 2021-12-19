package lib

import (
	"bytes"
	"encoding/json"
	"html/template"
	"log"

	"github.com/Masterminds/sprig"
	"gopkg.in/yaml.v2"
)

type ValuesObj map[string]interface{}

type TemplateData struct {
	Values ValuesObj
}

type GetYamlReturnValue struct {
	Yaml string `json:"yaml"`
	Err  error  `json:"err"`
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
			Err: err,
		})
	}

	asd := TemplateData{valuesData}

	var output bytes.Buffer

	t := template.Must(template.New("template").Funcs(sprig.FuncMap()).Parse(templateYaml))
	if err := t.Execute(&output, asd); err != nil {
		log.Println("executing template:", err)
	}

	return toJson(GetYamlReturnValue{
		Yaml: output.String(),
	})
}
