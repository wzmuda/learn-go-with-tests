package main

// variadic functions!
func SumAll(numbersToSum ...[]int) []int {
	// make(t Type size ...IntegerType) Type creates a slice
	// sums := make([]int, len(numbersToSum))
	// for i, numbers := range numbersToSum {
	// 	sums[i] = Sum(numbers)
	// }

	var sums []int
	for _, numbers := range numbersToSum {
		// append(slice []Type, elems ...Type) []Type
		// creates a new slice of Type with contents of slice
		// plus contents of elems
		sums = append(sums, Sum(numbers))
	}

	return sums
}
