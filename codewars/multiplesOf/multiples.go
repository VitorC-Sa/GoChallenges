package multiplesof

func MySolution(number int) int {
	sum := 0

	if number >= 0 {
		for i := 0; i < number; i++ {
			if i%3 == 0 || i%5 == 0 {
				sum += i
			}
		}
	}

	return sum
}

func BestSolution(number int) int {
	sum := 0
	for i := 1; i < number; i++ {
		if (i%3 == 0) || (i%5 == 0) {
			sum += i
		}
	}
	return sum
}
