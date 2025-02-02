package spinner_test

import (
	"fmt"

	"github.com/akramarenkov/spinner"
)

func ExampleSpinner() {
	spn := spinner.New(-1, 2)
	fmt.Println(spn.Actual())

	for range 7 {
		spn.Spin()
		fmt.Println(spn.Actual())
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
	spn := spinner.New(-1, 2)
	fmt.Println(spn.Actual())

	for range 7 {
		fmt.Println(spn.Next())
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
