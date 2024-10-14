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

func TestStep(t *testing.T) {
	expected := []int{-2, 1, 2, -2, 1, 2, -2}
	actual := make([]int, 0, len(expected))

	for number := range Step(-2, 2, 3) {
		if len(actual) == len(expected) {
			break
		}

		actual = append(actual, number)
	}

	require.Equal(t, expected, actual)
}

func TestStepPanic(t *testing.T) {
	func() {
		defer func() {
			require.Equal(t, ErrStepNegative, recover())
		}()

		for number := range Step(-2, 2, -1) {
			_ = number
		}
	}()

	func() {
		defer func() {
			require.Equal(t, ErrStepZero, recover())
		}()

		for number := range Step(-2, 2, 0) {
			_ = number
		}
	}()
}

func BenchmarkIter(b *testing.B) {
	for number := range Iter(1, b.N) {
		if number == b.N {
			return
		}
	}
}

func BenchmarkIterTwoLevel(b *testing.B) {
	for range b.N {
		for number := range Iter(1, 1) {
			if number == 1 {
				break
			}
		}
	}
}

func BenchmarkStep(b *testing.B) {
	for number := range Step(1, b.N, 1) {
		if number == b.N {
			return
		}
	}
}

func BenchmarkStepTwoLevel(b *testing.B) {
	for range b.N {
		for number := range Step(1, 1, 1) {
			if number == 1 {
				break
			}
		}
	}
}
