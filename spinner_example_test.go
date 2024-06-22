package spinner_test

import (
	"fmt"

	"github.com/akramarenkov/spinner"
)

func ExampleSpinner() {
	spinner := spinner.New(0, 2)

	fmt.Println(spinner.Actual())

	for range 5 {
		spinner.Spin()
		fmt.Println(spinner.Actual())
	}

	// Output:
	// 0
	// 1
	// 2
	// 0
	// 1
	// 2
}
