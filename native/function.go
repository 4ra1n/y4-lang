package native

import (
	"errors"
	"reflect"

	"github.com/4ra1n/y4-lang/log"
)

type NativeFunction struct {
	name      string
	fn        reflect.Value
	numParams int
}

func NewNativeFunction(name string, fn interface{}) (*NativeFunction, error) {
	fnVal := reflect.ValueOf(fn)
	if fnVal.Kind() != reflect.Func {
		return nil, errors.New("not a function")
	}
	return &NativeFunction{
		name:      name,
		fn:        fnVal,
		numParams: fnVal.Type().NumIn(),
	}, nil
}

func (n *NativeFunction) GetName() string {
	return n.name
}

func (n *NativeFunction) String() string {
	return "<native:" + n.name + ">"
}

func (n *NativeFunction) GetNumParam() int {
	return n.numParams
}

func (n *NativeFunction) Invoke(args []interface{}) (interface{}, error) {
	if len(args) != n.numParams {
		return nil, errors.New("incorrect number of arguments")
	}
	in := make([]reflect.Value, len(args))
	for i, arg := range args {
		in[i] = reflect.ValueOf(arg)
	}
	defer func() {
		err := recover()
		if err != nil {
			log.Errorf("invoke panic: %s", err)
		}
	}()
	result := n.fn.Call(in)
	if len(result) > 0 {
		return result[0].Interface(), nil
	}
	return nil, nil
}
