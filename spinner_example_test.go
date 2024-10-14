package spinner_test

import (
	"fmt"

	"github.com/akramarenkov/spinner"
)

func ExampleSpinner() {
	spinner := spinner.New(-1, 2)
	fmt.Println(spinner.Actual())

	for range 7 {
		spinner.Spin()
		fmt.Println(spinner.Actual())
	}
	// Output:
	// -1
	// 0
	// 1
	// 2
	// -1
	// 0
	// 1
	// 2
}

func ExampleSpinner_Next() {
	spinner := spinner.New(-1, 2)
	fmt.Println(spinner.Actual())

	for range 7 {
		fmt.Println(spinner.Next())
	}
	// Output:
	// -1
	// 0
	// 1
	// 2
	// -1
	// 0
	// 1
	// 2
}
