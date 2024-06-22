package spinner

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSpinner(t *testing.T) {
	spinner := New(0, 2)

	require.Equal(t, 0, spinner.Actual())

	spinner.Spin()
	require.Equal(t, 1, spinner.Actual())

	spinner.Spin()
	require.Equal(t, 2, spinner.Actual())

	spinner.Spin()
	require.Equal(t, 0, spinner.Actual())

	spinner.Spin()
	require.Equal(t, 1, spinner.Actual())

	spinner.Spin()
	require.Equal(t, 2, spinner.Actual())

	spinner.Spin()
	require.Equal(t, 0, spinner.Actual())
}

func TestSpinnerEndIsEqualBegin(t *testing.T) {
	spinner := New(0, 0)

	require.Equal(t, 0, spinner.Actual())

	spinner.Spin()
	require.Equal(t, 0, spinner.Actual())

	spinner.Spin()
	require.Equal(t, 0, spinner.Actual())

	spinner.Spin()
	require.Equal(t, 0, spinner.Actual())

	spinner = New(1, 1)

	require.Equal(t, 1, spinner.Actual())

	spinner.Spin()
	require.Equal(t, 1, spinner.Actual())

	spinner.Spin()
	require.Equal(t, 1, spinner.Actual())

	spinner.Spin()
	require.Equal(t, 1, spinner.Actual())
}

func TestSpinnerEndIsLessBegin(t *testing.T) {
	spinner := New(0, -1)

	require.Equal(t, 0, spinner.Actual())

	spinner.Spin()
	require.Equal(t, 0, spinner.Actual())

	spinner.Spin()
	require.Equal(t, 0, spinner.Actual())

	spinner.Spin()
	require.Equal(t, 0, spinner.Actual())

	spinner = New(1, 0)

	require.Equal(t, 1, spinner.Actual())

	spinner.Spin()
	require.Equal(t, 1, spinner.Actual())

	spinner.Spin()
	require.Equal(t, 1, spinner.Actual())

	spinner.Spin()
	require.Equal(t, 1, spinner.Actual())
}

func BenchmarkSpinner(b *testing.B) {
	spinner := New(0, 10)

	b.ResetTimer()

	for run := 0; run < b.N; run++ {
		_ = spinner.Actual()
		spinner.Spin()
	}
}
