package spinner

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIter(t *testing.T) {
	expected := []int{-2, -1, 0, 1, 2, -2, -1, 0, 1, 2, -2}
	actual := make([]int, 0, len(expected))

	for number := range Iter(-2, 2) {
		if len(actual) == len(expected) {
			break
		}

		actual = append(actual, number)
	}

	require.Equal(t, expected, actual)
}

func TestIterStep(t *testing.T) {
	expected := []int{-2, 1, 2, -2, 1, 2, -2}
	actual := make([]int, 0, len(expected))

	for number := range IterStep(-2, 2, 3) {
		if len(actual) == len(expected) {
			break
		}

		actual = append(actual, number)
	}

	require.Equal(t, expected, actual)
}

func TestIterStepPanic(t *testing.T) {
	func() {
		defer func() {
			require.Equal(t, ErrStepNegative, recover())
		}()

		for number := range IterStep(-2, 2, -1) {
			_ = number
		}
	}()

	func() {
		defer func() {
			require.Equal(t, ErrStepZero, recover())
		}()

		for number := range IterStep(-2, 2, 0) {
			_ = number
		}
	}()
}
