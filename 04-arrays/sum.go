package main

func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}

func SumAll(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return
}

func SumAllTails(numbersToSum ...[]int) (sumsOfTails []int) {
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sumsOfTails = append(sumsOfTails, 0)
		} else {
			// slice numbers[1:] to exclude first number (definition of a TAIL)
			tail := numbers[1:]
			sumsOfTails = append(sumsOfTails, Sum(tail))
		}

	}
	return
}
