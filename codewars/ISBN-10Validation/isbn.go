package isbn10validation

import (
	"strconv"
	"strings"
)

func MySolution(isbn string) bool {
	var res int
	for i, v := range strings.Split(isbn, "") {
		var n int
		var err error
		if (v == "X" || v == "x") && (i+1 == 10) {
			n = 10
		} else {
			n, err = strconv.Atoi(v)
			if (err != nil) || (n < 0 || n > 9) {
				return false
			}
		}
		res += n * (i + 1)
	}
	return (res%11 == 0) && (len(isbn) == 10)
}
