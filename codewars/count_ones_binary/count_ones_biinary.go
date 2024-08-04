package countonesbinary

import "strconv"

func CountOnes(left, right uint64) uint64 {
	sum := uint64(0)

	for i := left; i <= right; i++ {
		bVal := strconv.FormatInt(int64(i), 2)
		for _, v := range bVal {
			if string(v) == "1" {
				sum++
			}
		}
	}

	return sum
}
