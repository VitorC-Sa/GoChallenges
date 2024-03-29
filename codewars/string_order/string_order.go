package stringorder

/*
* Your task is to sort a given string. Each word in the string will contain a single number.
* This number is the position the word should have in the result.
*
* Note: Numbers can be from 1 to 9. So 1 will be the first word (not 0).
*
* If the input string is empty, return an empty string. The words in the input String will only contain valid consecutive numbers.
*
* Examples:
* 		"is2 Thi1s T4est 3a"  -->  "Thi1s is2 3a T4est"
* 		"4of Fo1r pe6ople g3ood th5e the2"  -->  "Fo1r the2 g3ood 4of th5e pe6ople"
* 		""  -->  ""
 */

import (
	"strconv"
	"strings"
)

//Extract a number from string (1-9, ignore 0)
func ExtractIntFromString(s string) int {
	nOneAscii := byte('1')
	nNineAscii := byte('9')

	for _, v := range []byte(s) {
		if v >= nOneAscii && v <= nNineAscii {
			res, _ := strconv.Atoi(string(v))
			return res
		}
	}
	return -1
}

func MySolution(sentence string) string {
	fields := strings.Fields(sentence)
	orderedFields := make([]string, len(fields))

	for _, v := range fields {
		n := ExtractIntFromString(v)
		if n != -1 {
			orderedFields[n-1] = v
		}
	}

	return strings.Join(orderedFields, " ")
}
