package multiplication_table

func MySolution(size int) [][]int {
	var m = make([][]int, 0, size)

	for i := 1; i <= size; i++ {
		var v = make([]int, 0, size)
		for j := 1; j <= size; j++ {
			v = append(v, i*j)
		}
		m = append(m, v)
	}

	return m
}
