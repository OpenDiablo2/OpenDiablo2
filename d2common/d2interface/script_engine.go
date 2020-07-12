package d2interface

import "github.com/robertkrimen/otto"

type ScriptEngine interface {
	AppComponent
	AllowEval()
	DisallowEval()
	ToValue(source interface{}) (otto.Value, error)
	AddFunction(name string, value interface{})
	RunScript(fileName string) (*otto.Value, error)
	Eval(code string) (string, error)
}
