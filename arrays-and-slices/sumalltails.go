package main

func SumAllTails(numbersToSum ...[]int) (sums []int) {
	for _, num := range numbersToSum {
		if len(num) == 0 {
			num = []int{0}
		}
		sums = append(sums, Sum(num[1:])) // s[low:high]
	}

	return sums
}
