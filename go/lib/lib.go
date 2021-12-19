package lib

import (
	"bytes"
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
	yaml string
	err  string
}

func GetYaml(templateYaml string, valuesYaml string) GetYamlReturnValue {
	valuesData := ValuesObj{}
	if err := yaml.Unmarshal([]byte(valuesYaml), &valuesData); err != nil {
		return "", err
	}

	asd := TemplateData{valuesData}

	var output bytes.Buffer

	t := template.Must(template.New("template").Funcs(sprig.FuncMap()).Parse(templateYaml))
	if err := t.Execute(&output, asd); err != nil {
		log.Println("executing template:", err)
	}

	return output.String(), nil
}
