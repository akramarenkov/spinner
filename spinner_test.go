package spinner

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSpinner(t *testing.T) {
	spinner := New(-2, 2)
	require.Equal(t, -2, spinner.Actual())

	spinner.Spin()
	require.Equal(t, -1, spinner.Actual())

	spinner.Spin()
	require.Equal(t, 0, spinner.Actual())

	spinner.Spin()
	require.Equal(t, 1, spinner.Actual())

	spinner.Spin()
	require.Equal(t, 2, spinner.Actual())

	spinner.Spin()
	require.Equal(t, -2, spinner.Actual())

	spinner.Spin()
	require.Equal(t, -1, spinner.Actual())

	spinner.Spin()
	require.Equal(t, 0, spinner.Actual())

	spinner.Spin()
	require.Equal(t, 1, spinner.Actual())

	spinner.Spin()
	require.Equal(t, 2, spinner.Actual())

	spinner.Spin()
	require.Equal(t, -2, spinner.Actual())
}

func TestSpinnerEndIsEqualBegin(t *testing.T) {
	spinner := New(1, 1)
	require.Equal(t, 1, spinner.Actual())

	spinner.Spin()
	require.Equal(t, 1, spinner.Actual())

	spinner.Spin()
	require.Equal(t, 1, spinner.Actual())

	spinner.Spin()
	require.Equal(t, 1, spinner.Actual())
}

func TestSpinnerEndIsEqualBeginNegative(t *testing.T) {
	spinner := New(-1, -1)
	require.Equal(t, -1, spinner.Actual())

	spinner.Spin()
	require.Equal(t, -1, spinner.Actual())

	spinner.Spin()
	require.Equal(t, -1, spinner.Actual())

	spinner.Spin()
	require.Equal(t, -1, spinner.Actual())
}

func TestSpinnerEndIsLessBegin(t *testing.T) {
	spinner := New(2, -2)
	require.Equal(t, 2, spinner.Actual())

	spinner.Spin()
	require.Equal(t, 1, spinner.Actual())

	spinner.Spin()
	require.Equal(t, 0, spinner.Actual())

	spinner.Spin()
	require.Equal(t, -1, spinner.Actual())

	spinner.Spin()
	require.Equal(t, -2, spinner.Actual())

	spinner.Spin()
	require.Equal(t, 2, spinner.Actual())

	spinner.Spin()
	require.Equal(t, 1, spinner.Actual())

	spinner.Spin()
	require.Equal(t, 0, spinner.Actual())

	spinner.Spin()
	require.Equal(t, -1, spinner.Actual())

	spinner.Spin()
	require.Equal(t, -2, spinner.Actual())

	spinner.Spin()
	require.Equal(t, 2, spinner.Actual())
}

func TestSpinnerIntegerBoundaries(t *testing.T) {
	spinner := New(math.MaxInt-2, math.MaxInt)
	require.Equal(t, math.MaxInt-2, spinner.Actual())

	spinner.Spin()
	require.Equal(t, math.MaxInt-1, spinner.Actual())

	spinner.Spin()
	require.Equal(t, math.MaxInt, spinner.Actual())

	spinner.Spin()
	require.Equal(t, math.MaxInt-2, spinner.Actual())

	spinner.Spin()
	require.Equal(t, math.MaxInt-1, spinner.Actual())

	spinner.Spin()
	require.Equal(t, math.MaxInt, spinner.Actual())

	spinner.Spin()
	require.Equal(t, math.MaxInt-2, spinner.Actual())
}

func TestSpinnerIntegerBoundariesMin(t *testing.T) {
	spinner := New(math.MinInt+2, math.MinInt)
	require.Equal(t, math.MinInt+2, spinner.Actual())

	spinner.Spin()
	require.Equal(t, math.MinInt+1, spinner.Actual())

	spinner.Spin()
	require.Equal(t, math.MinInt, spinner.Actual())

	spinner.Spin()
	require.Equal(t, math.MinInt+2, spinner.Actual())

	spinner.Spin()
	require.Equal(t, math.MinInt+1, spinner.Actual())

	spinner.Spin()
	require.Equal(t, math.MinInt, spinner.Actual())

	spinner.Spin()
	require.Equal(t, math.MinInt+2, spinner.Actual())
}

func BenchmarkSpinner(b *testing.B) {
	spinner := New(0, 10)

	actual := 0

	for range b.N {
		actual = spinner.Actual()
		spinner.Spin()
	}

	require.NotNil(b, actual)
}
