package validatebraces

import (
	"strings"
)

/* Write a function that takes a string of braces, and determines if the order of the braces is valid. It should return true if the string is valid, and false if it's invalid.

All input strings will be nonempty, and will only consist of parentheses, brackets and curly braces: ()[]{}.

What is considered Valid?
A string of braces is considered valid if all braces are matched with the correct brace. */

func MySolution(str string) bool {
	var indexBackup []int
	s := 0
	sBracesIni := "([{"
	sBracesEnd := ")]}"

	if isFirstIndexValid := strings.IndexByte(sBracesIni, str[0]); isFirstIndexValid == -1 {
		return false
	}

	for _, v := range str {
		if index := strings.IndexRune(sBracesIni, v); index != -1 { //get index of v if match with Ini Braces
			indexBackup = append(indexBackup, index)
			s++
		}
		if index := strings.IndexRune(sBracesEnd, v); index != -1 { //get index of v if match with End Braces
			l := len(indexBackup)
			if index != indexBackup[l-1] { //compare last End Brace with last Ini Brace
				return false
			}
			indexBackup = append(indexBackup, indexBackup[:l-1]...) //delete last Ini Brace index from indexBackup
			s--
		}
	}

	if s != 0 {
		return false
	}

	return true
}
