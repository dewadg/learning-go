package main

func numberMap(c []int, f func(int) int) []int {
	output := make([]int, 0)

	for _, item := range c {
		output = append(output, f(item))
	}

	return output
}

func numberFilter(c []int, f func(int) bool) []int {
	output := make([]int, 0)

	for _, item := range c {
		if f(item) {
			output = append(output, item)
		}
	}

	return output
}

func numberIncludes(c []int, val int) bool {
	for _, item := range c {
		if item == val {
			return true
		}
	}

	return false
}

func numberEach(c []int, f func(int)) {
	for _, item := range c {
		f(item)
	}
}

func numberReduce(c []int, f func(int, int) int, initial int) int {
	for _, item := range c {
		initial = f(initial, item)
	}

	return initial
}
