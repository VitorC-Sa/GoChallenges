package rgbtohex

/*
	The rgb function is incomplete. Complete it so that passing in RGB decimal values will result in a hexadecimal representation being returned. Valid decimal values for RGB are 0 - 255. Any values that fall out of that range must be rounded to the closest valid value.

	Note: Your answer should always be 6 characters long, the shorthand with 3 will not work here.

	The following are examples of expected output values:
		rgb(255, 255, 255) -- returns FFFFFF
		rgb(255, 255, 300) -- returns FFFFFF
		rgb(0, 0, 0) -- returns 000000
		rgb(148, 0, 211) -- returns 9400D3
*/

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
