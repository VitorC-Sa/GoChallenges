package maximum_subarray_sum

func MySolution(numbers []int) int {
	fMax := func(x, y int) int {
		if x >= y {
			return x
		}
		return y
	}

	max := 0
	for i := range numbers {
		sum := 0
		for j := i; j < len(numbers); j++ {
			sum += numbers[j]
			max = fMax(max, sum)
		}
	}
	return max
}
