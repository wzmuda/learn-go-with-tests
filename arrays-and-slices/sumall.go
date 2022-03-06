package main

// variadic functions!
func SumAll(numbersToSum ...[]int) []int {
	// make(t Type size ...IntegerType) Type creates a slice
	sums := make([]int, len(numbersToSum))
	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}

	return sums
}
