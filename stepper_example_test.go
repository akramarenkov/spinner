package spinner_test

import (
	"fmt"

	"github.com/akramarenkov/spinner"
)

func ExampleStepper() {
	stp, err := spinner.NewStepper(-1, 2, 2)
	if err != nil {
		panic(err)
	}

	fmt.Println(stp.Actual())

	for range 7 {
		stp.Spin()
		fmt.Println(stp.Actual())
	}
	// Output:
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
	stp, err := spinner.NewStepper(-1, 2, 2)
	if err != nil {
		panic(err)
	}

	fmt.Println(stp.Actual())

	for range 7 {
		fmt.Println(stp.Next())
	}
	// Output:
	// -1
	// 1
	// 2
	// -1
	// 1
	// 2
	// -1
	// 1
}
