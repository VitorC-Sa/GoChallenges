package move_zeros_to_the_end

func MoveZeros(arr []int) []int {
	zeroArr := []int{}

	for i := 0; i < len(arr); {
		if arr[i] == 0 && len(arr) >= i+1 {
			arr = append(arr[0:i], arr[i+1:]...)
			zeroArr = append(zeroArr, 0)
			continue
		}
		i++
	}

	return append(arr, zeroArr...)
}
