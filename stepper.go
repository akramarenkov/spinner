package spinner

import (
	"golang.org/x/exp/constraints"
)

// Infinitely iterates over integer values ​​in a given range ​with the specified step.
type Stepper[Type constraints.Integer] struct {
	actual Type
	begin  Type
	end    Type
	step   Type
}

// Creates Stepper instance.
//
// If begin is greater than end, the actual value will decrease, otherwise it will
// increase.
//
// Step must be greater than zero.
func NewStepper[Type constraints.Integer](begin Type, end Type, step Type) (*Stepper[Type], error) {
	if step < 0 {
		return nil, ErrStepNegative
	}

	if step == 0 {
		return nil, ErrStepZero
	}

	spn := &Stepper[Type]{
		actual: begin,
		begin:  begin,
		end:    end,
		step:   step,
	}

	return spn, nil
}

// Returns actual value.
func (spn *Stepper[Type]) Actual() Type {
	return spn.actual
}

// Increases/decreases the actual value of the counter. If the actual value is equal
// to the end value, it will be reset to the begin value.
func (spn *Stepper[Type]) Spin() {
	spn.actual = spn.spin()
}

// Increases/decreases the actual value of the counter and return it. If the actual
// value is equal to the end value, it will be reset to the begin value.
func (spn *Stepper[Type]) Next() Type {
	spn.Spin()
	return spn.Actual()
}

func (spn *Stepper[Type]) spin() Type {
	if spn.begin <= spn.end {
		return spn.forward()
	}

	return spn.backward()
}

func (spn *Stepper[Type]) forward() Type {
	if spn.actual == spn.end {
		return spn.begin
	}

	actual := spn.actual + spn.step

	// integer overflow
	if actual < spn.actual {
		return spn.end
	}

	// step is not a multiple of the length between begin and end
	if actual > spn.end {
		return spn.end
	}

	return actual
}

func (spn *Stepper[Type]) backward() Type {
	if spn.actual == spn.end {
		return spn.begin
	}

	actual := spn.actual - spn.step

	// integer overflow
	if actual > spn.actual {
		return spn.end
	}

	// step is not a multiple of the length between begin and end
	if actual < spn.end {
		return spn.end
	}

	return actual
}
