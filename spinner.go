package spinner

import (
	"golang.org/x/exp/constraints"
)

// Infinitely iterates over integer values ​​in a given range ​​in increments of one.
type Spinner[Type constraints.Integer] struct {
	actual Type
	begin  Type
	end    Type
}

// Creates Spinner instance.
//
// If begin is greater than end, the actual value will decrease, otherwise it will
// increase.
func New[Type constraints.Integer](begin Type, end Type) *Spinner[Type] {
	spn := &Spinner[Type]{
		actual: begin,
		begin:  begin,
		end:    end,
	}

	return spn
}

// Returns actual value.
func (spn *Spinner[Type]) Actual() Type {
	return spn.actual
}

// Increases/decreases the actual value of the counter. If the actual value is equal
// to the end value, it will be reset to the begin value.
func (spn *Spinner[Type]) Spin() {
	spn.actual = spn.spin()
}

// Increases/decreases the actual value of the counter and return it. If the actual
// value is equal to the end value, it will be reset to the begin value.
func (spn *Spinner[Type]) Next() Type {
	spn.Spin()
	return spn.Actual()
}

func (spn *Spinner[Type]) spin() Type {
	if spn.begin <= spn.end {
		return spn.forward()
	}

	return spn.backward()
}

func (spn *Spinner[Type]) forward() Type {
	if spn.actual == spn.end {
		return spn.begin
	}

	return spn.actual + 1
}

func (spn *Spinner[Type]) backward() Type {
	if spn.actual == spn.end {
		return spn.begin
	}

	return spn.actual - 1
}
