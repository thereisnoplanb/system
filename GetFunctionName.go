package system

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

// Retrieves the name of the specified function or the calling function if no function is specified.
//
// Parameters:
//
//	function any - An optional parameter that accepts a variable number of interface{} arguments. Only first is accepted.
//
// Returns:
//
//	name string - Name of the function.
//	err error - An error if the provided argument is not a function, nil otherwise.
//
//	func exampleFunction() {}
//
//	main {
//	    name, err := system.GetFunctionName()
//	    if err != nil {
//	        fmt.Println("Error:", err)
//	    } else {
//	        fmt.Println("Function name:", name) // prints main
//	    }
//	    name, err := GetFunctionName(exampleFunction)
//	    if err != nil {
//	        fmt.Println("Error:", err)
//	    } else {
//	        fmt.Println("Function name:", name) // prints exampleFunction
//	    }
//	}
func GetFunctionName(function ...interface{}) (name string, err error) {
	var pc uintptr
	if len(function) == 0 {
		pc, _, _, _ = runtime.Caller(1)
	} else {
		value := reflect.ValueOf(function[0])
		if value.Kind() != reflect.Func {
			return "", fmt.Errorf("expected a function, but got %s", value.Kind())
		}
		pc = value.Pointer()
	}
	name = runtime.FuncForPC(pc).Name()
	name = strings.ReplaceAll(name, "[...]", "")
	name = name[strings.LastIndex(name, ".")+1:]
	return name, nil
}
