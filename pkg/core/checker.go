package core

import (
	"fmt"
)

// Checker checks a history of operations.
type Checker interface {
	fmt.Stringer
	Check(m Model, ops []Operation) (bool, error)
}
