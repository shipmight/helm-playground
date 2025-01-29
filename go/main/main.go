package main

import (
	"syscall/js"

	"github.com/shipmight/helm-playground/go/lib"
)

func jsFuncWrapper() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 2 {
			return `{"err":""}`
		}
		templateYaml := args[0].String()
		valuesYaml := args[1].String()
		returnValue := lib.GetYaml(templateYaml, valuesYaml, lib.GetYamlConfig{MaxTplRuns: 100})
		return returnValue
	})
}

func main() {
	js.Global().Set("GetYaml", jsFuncWrapper())
	<-make(chan bool)
}
