package d2script

import (
	"fmt"
	"io/ioutil"

	"github.com/robertkrimen/otto"
	_ "github.com/robertkrimen/otto/underscore" // This causes the runtime to support underscore.js
)

type ScriptEngine struct {
	vm *otto.Otto
}

func CreateScriptEngine() *ScriptEngine {
	result := &ScriptEngine{
		vm: otto.New(),
	}

	result.vm.Set("debugPrint", func(call otto.FunctionCall) otto.Value {
		fmt.Printf("Script: %s\n", call.Argument(0).String())
		return otto.Value{}
	})

	return result
}

func (s *ScriptEngine) ToValue(source interface{}) (otto.Value, error) {
	return s.vm.ToValue(source)
}

func (s *ScriptEngine) AddFunction(name string, value interface{}) {
	s.vm.Set(name, value)
}

func (s *ScriptEngine) RunScript(fileName string) (*otto.Value, error) {
	fileData, _ := ioutil.ReadFile(fileName)
	val, err := s.vm.Run(string(fileData))
	if err != nil {
		fmt.Printf("Error running script: %s\n", err.Error())
		return nil, err
	}
	return &val, nil
}
