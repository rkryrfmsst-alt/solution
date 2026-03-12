// Package v2 is major version of package solution
package v2

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

// Add b to a
func Add[T Number](a, b T) T {
	return a + b
}
