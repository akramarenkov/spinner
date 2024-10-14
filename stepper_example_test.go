package spinner_test

import (
	"fmt"

	"github.com/akramarenkov/spinner"
)

func ExampleStepper() {
	spinner, err := spinner.NewStepper(-1, 2, 2)
	fmt.Println(err)
	fmt.Println(spinner.Actual())

	for range 7 {
		spinner.Spin()
		fmt.Println(spinner.Actual())
	}
	// Output:
	// <nil>
	// -1
	// 1
	// 2
	// -1
	// 1
	// 2
	// -1
	// 1
}

func ExampleStepper_Next() {
	spinner, err := spinner.NewStepper(-1, 2, 2)
	fmt.Println(err)
	fmt.Println(spinner.Actual())

	for range 7 {
		fmt.Println(spinner.Next())
	}
	// Output:
	// <nil>
	// -1
	// 1
	// 2
	// -1
	// 1
	// 2
	// -1
	// 1
}
