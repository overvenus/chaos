package core

import (
	"fmt"
)

// Model specifies the behavior of a data object.
type Model interface {
	fmt.Stringer

	// Initial state of the data object.
	Init() interface{}

	// Step function for the data object. Returns whether or not the system
	// could take this step with the given inputs and outputs and also
	// returns the new state. This should not mutate the existing state.
	Step(state interface{}, input interface{}, output interface{}) (bool, interface{})

	// Equality on states.
	Equal(state1, state2 interface{}) bool
}

// Operation action
const (
	InvokeOperation = "call"
	ReturnOperation = "return"
)

// Operation of a data object.
type Operation struct {
	Action string      `json:"action"`
	Proc   int64       `json:"proc"`
	Data   interface{} `json:"data"`
}
