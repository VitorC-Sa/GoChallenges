package isbn_10_validation

/*
	ISBN-10 identifiers are ten digits long. The first nine characters are digits 0-9. The last digit can be 0-9 or X, to indicate a value of 10.

	An ISBN-10 number is valid if the sum of the digits multiplied by their position modulo 11 equals zero.
*/

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
