// Package super_duper_calculator is a super-duper calculator that I wrote while learning Go modules
// ain't it fun!
// You can learn about what it can do using incredible [mathisfun] website
// [mathisfun]: https://www.mathisfun.com/numbers/addition.html
package super_duper_calculator

import (
	"golang.org/x/exp/constraints"
)

type Calcable interface {
	constraints.Integer | constraints.Float
}

func Add[N Calcable](a, b N) N {
	return a + b
}
