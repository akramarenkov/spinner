// Iterates over integer values, while returning the begin value when trying to
// go beyond the end value.
package spinner

import (
	"golang.org/x/exp/constraints"
)

// Iterates over integer values, while returning the begin value when trying to
// go beyond the end value.
type Spinner[Type constraints.Integer] struct {
	actual Type
	begin  Type
	end    Type
}

// Creates Spinner instance.
func New[Type constraints.Integer](begin Type, end Type) *Spinner[Type] {
	spn := &Spinner[Type]{
		actual: begin,
		begin:  begin,
		end:    end,
	}

	return spn
}

// Returns actual value of counter.
func (spn *Spinner[Type]) Actual() Type {
	return spn.actual
}

// Increases/decreases the current value of the counter. If the next value is out of
// end value, it will be reset to the begin value.
func (spn *Spinner[Type]) Spin() {
	spn.actual = spn.spin()
}

func (spn *Spinner[Type]) spin() Type {
	if spn.begin <= spn.end {
		return spn.forward()
	}

	return spn.backward()
}

func (spn *Spinner[Type]) forward() Type {
	if spn.actual >= spn.end {
		return spn.begin
	}

	return spn.actual + 1
}

func (spn *Spinner[Type]) backward() Type {
	if spn.actual <= spn.end {
		return spn.begin
	}

	return spn.actual - 1
}
