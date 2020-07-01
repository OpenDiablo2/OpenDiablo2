package d2script

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/robertkrimen/otto"
	_ "github.com/robertkrimen/otto/underscore" // This causes the runtime to support underscore.js
)

// ScriptEngine allows running JavaScript scripts
type ScriptEngine struct {
	vm *otto.Otto
}

// CreateScriptEngine creates the script engine and returns a pointer to it.
func CreateScriptEngine() *ScriptEngine {
	result := &ScriptEngine{
		vm: otto.New(),
	}

	err := result.vm.Set("debugPrint", func(call otto.FunctionCall) otto.Value {
		fmt.Printf("Script: %s\n", call.Argument(0).String())
		return otto.Value{}
	})
	if err != nil {
		fmt.Printf("could not bind the 'debugPrint' to the given function in script engine")
	}

	return result
}

// ToValue converts the given interface{} value to a otto.Value
func (s *ScriptEngine) ToValue(source interface{}) (otto.Value, error) {
	return s.vm.ToValue(source)
}

// AddFunction adds the given function to the script engine with the given name.
func (s *ScriptEngine) AddFunction(name string, value interface{}) {
	err := s.vm.Set(name, value)
	if err != nil {
		fmt.Printf("could not add the '%s' function to the script engine", name)
	}
}

// RunScript runs the script file within the given path.
func (s *ScriptEngine) RunScript(fileName string) (*otto.Value, error) {
	fileData, err := ioutil.ReadFile(filepath.Clean(fileName))
	if err != nil {
		fmt.Printf("could not read script file: %s\n", err.Error())
		return nil, err
	}

	val, err := s.vm.Run(string(fileData))
	if err != nil {
		fmt.Printf("Error running script: %s\n", err.Error())
		return nil, err
	}

	return &val, nil
}
