package wasm

import (
	"syscall/js"

	"github.com/codeclown/helm-playground/go/lib"
)

func jsFuncWrapper() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 2 {
			return "Invalid no of arguments passed"
		}
		templateYaml := args[0].String()
		valuesYaml := args[1].String()
		returned := lib.GetYaml(templateYaml, valuesYaml)
		return returned
	})
}

func main() {
	js.Global().Set("GetYaml", jsFuncWrapper())
	<-make(chan bool)
}
