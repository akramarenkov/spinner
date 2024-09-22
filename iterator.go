package spinner

import (
	"iter"

	"golang.org/x/exp/constraints"
)

// A range iterator for infinite iteration from begin to end inclusive with step one.
//
// If begin is greater than end, the return value will be decremented, otherwise it
// will be incremented.
func Iter[Type constraints.Integer](begin, end Type) iter.Seq[Type] {
	spinner := New(begin, end)

	iterator := func(yield func(Type) bool) {
		for ; ; spinner.Spin() {
			if !yield(spinner.Actual()) {
				return
			}
		}
	}

	return iterator
}

// A range iterator for infinite iteration from begin to end inclusive ​with the
// specified step.
//
// If begin is greater than end, the return value will be decremented, otherwise it
// will be incremented.
//
// If a zero or negative step is specified, the iterator will panic.
func IterStep[Type constraints.Integer](begin, end, step Type) iter.Seq[Type] {
	stepper, err := NewStepper(begin, end, step)
	if err != nil {
		panic(err)
	}

	iterator := func(yield func(Type) bool) {
		for ; ; stepper.Spin() {
			if !yield(stepper.Actual()) {
				return
			}
		}
	}

	return iterator
}
