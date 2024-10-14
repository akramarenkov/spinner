package spinner_test

import (
	"fmt"

	"github.com/akramarenkov/spinner"
)

func ExampleIter() {
	iterations := 0

	for number := range spinner.Iter(-1, 2) {
		fmt.Println(number)

		iterations++

		if iterations == 8 {
			return
		}
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

func ExampleStep() {
	iterations := 0

	for number := range spinner.Step(-1, 2, 2) {
		fmt.Println(number)

		iterations++

		if iterations == 8 {
			return
		}
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
