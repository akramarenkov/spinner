package spinner

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStepper(t *testing.T) {
	stepper, err := NewStepper(-2, 2, 3)
	require.NoError(t, err)
	require.Equal(t, -2, stepper.Actual())

	stepper.Spin()
	require.Equal(t, 1, stepper.Actual())

	stepper.Spin()
	require.Equal(t, 2, stepper.Actual())

	stepper.Spin()
	require.Equal(t, -2, stepper.Actual())

	stepper.Spin()
	require.Equal(t, 1, stepper.Actual())

	stepper.Spin()
	require.Equal(t, 2, stepper.Actual())

	stepper.Spin()
	require.Equal(t, -2, stepper.Actual())
}

func TestStepperError(t *testing.T) {
	stepper, err := NewStepper(-2, 2, 0)
	require.Error(t, err)
	require.Nil(t, stepper)

	stepper, err = NewStepper(-2, 2, -1)
	require.Error(t, err)
	require.Nil(t, stepper)
}

func TestStepperNext(t *testing.T) {
	stepper, err := NewStepper(-2, 2, 3)
	require.NoError(t, err)
	require.Equal(t, -2, stepper.Actual())
	require.Equal(t, 1, stepper.Next())
	require.Equal(t, 2, stepper.Next())
	require.Equal(t, -2, stepper.Next())
	require.Equal(t, 1, stepper.Next())
	require.Equal(t, 2, stepper.Next())
	require.Equal(t, -2, stepper.Next())
}

func TestStepperEndIsEqualBegin(t *testing.T) {
	stepper, err := NewStepper(-1, -1, 3)
	require.NoError(t, err)
	require.Equal(t, -1, stepper.Actual())

	stepper.Spin()
	require.Equal(t, -1, stepper.Actual())

	stepper.Spin()
	require.Equal(t, -1, stepper.Actual())

	stepper.Spin()
	require.Equal(t, -1, stepper.Actual())
}

func TestStepperEndIsLessBegin(t *testing.T) {
	stepper, err := NewStepper(2, -2, 3)
	require.NoError(t, err)
	require.Equal(t, 2, stepper.Actual())

	stepper.Spin()
	require.Equal(t, -1, stepper.Actual())

	stepper.Spin()
	require.Equal(t, -2, stepper.Actual())

	stepper.Spin()
	require.Equal(t, 2, stepper.Actual())

	stepper.Spin()
	require.Equal(t, -1, stepper.Actual())

	stepper.Spin()
	require.Equal(t, -2, stepper.Actual())

	stepper.Spin()
	require.Equal(t, 2, stepper.Actual())
}

func TestStepperIntegerBoundaries(t *testing.T) {
	stepper, err := NewStepper(math.MaxInt-2, math.MaxInt, 3)
	require.NoError(t, err)
	require.Equal(t, math.MaxInt-2, stepper.Actual())

	stepper.Spin()
	require.Equal(t, math.MaxInt, stepper.Actual())

	stepper.Spin()
	require.Equal(t, math.MaxInt-2, stepper.Actual())

	stepper.Spin()
	require.Equal(t, math.MaxInt, stepper.Actual())

	stepper.Spin()
	require.Equal(t, math.MaxInt-2, stepper.Actual())
}

func TestStepperIntegerBoundariesMin(t *testing.T) {
	stepper, err := NewStepper(math.MinInt+2, math.MinInt, 3)
	require.NoError(t, err)
	require.Equal(t, math.MinInt+2, stepper.Actual())

	stepper.Spin()
	require.Equal(t, math.MinInt, stepper.Actual())

	stepper.Spin()
	require.Equal(t, math.MinInt+2, stepper.Actual())

	stepper.Spin()
	require.Equal(t, math.MinInt, stepper.Actual())

	stepper.Spin()
	require.Equal(t, math.MinInt+2, stepper.Actual())
}

func BenchmarkStepper(b *testing.B) {
	stepper, err := NewStepper(0, 10, 1)
	require.NoError(b, err)

	actual := 0

	for range b.N {
		actual = stepper.Actual()
		stepper.Spin()
	}

	require.NotNil(b, actual)
}

func BenchmarkStepperNext(b *testing.B) {
	stepper, err := NewStepper(0, 10, 1)
	require.NoError(b, err)

	actual := 0

	for range b.N {
		actual = stepper.Next()
	}

	require.NotNil(b, actual)
}
