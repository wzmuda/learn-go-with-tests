package main

func Sum(numbers [5]int) (sum int) {
	// for i := 0; i < 5; i++ {
	// 	sum += numbers[i]
	// }

	// for index, value := range numbers {
	for _, value := range numbers { // we can ignore index
		sum += value
	}

	return
}
