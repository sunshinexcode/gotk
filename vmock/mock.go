package vmock

import (
	"sync"
	"time"

	"github.com/agiledragon/gomonkey/v2"
)

type (
	// OutputCell is an alias for gomonkey.OutputCell
	OutputCell = gomonkey.OutputCell
	// Params is an alias for gomonkey.Params
	Params = gomonkey.Params
	// Patches is an alias for gomonkey.Patches
	Patches = gomonkey.Patches
)

var mockMutex sync.Mutex

// ApplyFunc applies a mock to a function
// target: the function to be mocked
// double: the mock implementation
// Returns: a Patches instance for cleanup
func ApplyFunc(target, double interface{}) *Patches {
	return gomonkey.ApplyFunc(target, double)
}

// ApplyFuncReturn applies a mock to a function with specified return values
// target: the function to be mocked
// output: the values to be returned
// Returns: a Patches instance for cleanup
func ApplyFuncReturn(target interface{}, output ...interface{}) *Patches {
	return gomonkey.ApplyFuncReturn(target, output...)
}

// ApplyFuncSeq applies a mock to a function with a sequence of outputs
// target: the function to be mocked
// outputs: sequence of output values
// Returns: a Patches instance for cleanup
func ApplyFuncSeq(target interface{}, outputs []OutputCell) *Patches {
	return gomonkey.ApplyFuncSeq(target, outputs)
}

// ApplyFuncVar applies a mock to a variable function
// target: the variable function to be mocked
// double: the mock implementation
// Returns: a Patches instance for cleanup
func ApplyFuncVar(target, double interface{}) *Patches {
	return gomonkey.ApplyFuncVar(target, double)
}

// ApplyFuncVarReturn applies a mock to a variable function with specified return values
// target: the variable function to be mocked
// output: the values to be returned
// Returns: a Patches instance for cleanup
func ApplyFuncVarReturn(target interface{}, output ...interface{}) *Patches {
	return gomonkey.ApplyFuncVarReturn(target, output...)
}

// ApplyFuncVarSeq applies a mock to a variable function with a sequence of outputs
// target: the variable function to be mocked
// outputs: sequence of output values
// Returns: a Patches instance for cleanup
func ApplyFuncVarSeq(target interface{}, outputs []OutputCell) *Patches {
	return gomonkey.ApplyFuncVarSeq(target, outputs)
}

// ApplyGlobalVar applies a mock to a global variable
// target: the global variable to be mocked
// double: the mock value
// Returns: a Patches instance for cleanup
func ApplyGlobalVar(target, double interface{}) *Patches {
	return gomonkey.ApplyGlobalVar(target, double)
}

// ApplyMethod applies a mock to a method
// target: the target type
// methodName: name of the method to be mocked
// double: the mock implementation
// Returns: a Patches instance for cleanup
func ApplyMethod(target interface{}, methodName string, double interface{}) *Patches {
	return gomonkey.ApplyMethod(target, methodName, double)
}

// ApplyMethodFunc applies a mock function to a method
// target: the target type
// methodName: name of the method to be mocked
// doubleFunc: the mock function
// Returns: a Patches instance for cleanup
func ApplyMethodFunc(target interface{}, methodName string, doubleFunc interface{}) *Patches {
	return gomonkey.ApplyMethodFunc(target, methodName, doubleFunc)
}

// ApplyMethodReturn applies a mock to a method with specified return values
// target: the target type
// methodName: name of the method to be mocked
// output: the values to be returned
// Returns: a Patches instance for cleanup
func ApplyMethodReturn(target interface{}, methodName string, output ...interface{}) *Patches {
	return gomonkey.ApplyMethodReturn(target, methodName, output...)
}

// ApplyMethodSeq applies a mock to a method with a sequence of outputs
// target: the target type
// methodName: name of the method to be mocked
// outputs: sequence of output values
// Returns: a Patches instance for cleanup
func ApplyMethodSeq(target interface{}, methodName string, outputs []OutputCell) *Patches {
	return gomonkey.ApplyMethodSeq(target, methodName, outputs)
}

// ApplyPrivateMethod applies a mock to a private method
// target: the target type
// methodName: name of the private method to be mocked
// double: the mock implementation
// Returns: a Patches instance for cleanup
func ApplyPrivateMethod(target interface{}, methodName string, double interface{}) *Patches {
	return gomonkey.ApplyPrivateMethod(target, methodName, double)
}

// Reset resets a single patch with thread safety
// patche: the patch to be reset
// Note: Includes a small delay after reset to ensure effectiveness
func Reset(patche *Patches) {
	// Lock the mutex to ensure thread safety
	mockMutex.Lock()
	defer mockMutex.Unlock()

	patche.Reset()
	// Avoid using the patch immediately after reset, which may cause it to be ineffective
	time.Sleep(1 * time.Millisecond)
}

// ResetAll resets multiple patches
// patches: array of patches to be reset
func ResetAll(patches []*Patches) {
	for _, patch := range patches {
		Reset(patch)
	}
}
