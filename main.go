package main

import (
	"fmt"
	"math"
)

func main() {
	numberCollection := []int{1, 2, 3, 4, 5}

	fmt.Println("Initial numberCollection")
	for _, item := range numberCollection {
		fmt.Println(item)
	}

	// Map operation
	squaredNumberCollection := numberMap(numberCollection, func(item int) int {
		return int(math.Pow(float64(item), 2))
	})

	fmt.Println("numberCollection squared")
	for _, item := range squaredNumberCollection {
		fmt.Println(item)
	}
	// End Map operation

	// Filter operation
	filteredNumberCollection := numberFilter(squaredNumberCollection, func(item int) bool {
		return item%2 == 0
	})

	fmt.Println("numberCollection filtered")
	for _, item := range filteredNumberCollection {
		fmt.Println(item)
	}
	// End Filter operation

	// Includes operation
	fmt.Println("numberCollection includes 4?")
	if numberIncludes(filteredNumberCollection, 4) {
		fmt.Println("Includes 4")
	}
	// End Includes operation

	// Each operation
	fmt.Println("numberCollection walked")
	numberEach(filteredNumberCollection, func(item int) {
		fmt.Println("The current value is", item)
	})
	// End Each operation
}
