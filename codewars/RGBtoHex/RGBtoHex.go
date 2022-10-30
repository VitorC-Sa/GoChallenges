package rgbtohex

import (
	"fmt"
	"strings"
)

func MySolution(r, g, b int) string {
	fields := []int{r, g, b}
	resp := make([]string, 0, 3)

	for _, v := range fields {
		if v < 0 {
			v = 0
		}
		if v > 255 {
			v = 255
		}
		hexVal := fmt.Sprintf("%02X", v)
		resp = append(resp, hexVal)
	}

	return strings.Join(resp, "")
}

func AnotherSolution(r, g, b int) string {
	getValid := func(v int) int {
		switch {
		case v < 0:
			return 0
		case v > 255:
			return 255
		default:
			return v
		}
	}

	return fmt.Sprintf("%02X%02X%02X", getValid(r), getValid(g), getValid(b))
}
