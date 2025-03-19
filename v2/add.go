// this is the first public module for add two number
package add

import "golang.org/x/exp/constraints"

// Number accept both Int and Float
type Number interface {
	constraints.Integer | constraints.Float
}

// An Add function is use for add a and b to an output
// see more https://www.mathsisfun.com/numbers/addition.html
func Add[N Number](a, b N) N {
	return a + b
}
