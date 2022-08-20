package validatebraces

import (
	"strings"
)

/* Write a function that takes a string of braces, and determines if the order of the braces is valid. It should return true if the string is valid, and false if it's invalid.

All input strings will be nonempty, and will only consist of parentheses, brackets and curly braces: ()[]{}.

What is considered Valid?
A string of braces is considered valid if all braces are matched with the correct brace. */

func MySolution(str string) bool {
	const sBracesIni = "([{"
	const sBracesEnd = ")]}"
	sIniIndexBackup := []int{}
	sentinel := 0

	if isFirstIndexValid := strings.IndexByte(sBracesIni, str[0]); isFirstIndexValid == -1 {
		return false
	}

	for _, v := range str {
		if index := strings.IndexRune(sBracesIni, v); index != -1 { //get index of v if match with Ini Braces
			sentinel++
			sIniIndexBackup = append(sIniIndexBackup, index)
		}
		if index := strings.IndexRune(sBracesEnd, v); index != -1 { //get index of v if match with End Braces
			sentinel--

			if sIniIndexBackup[len(sIniIndexBackup)-1] != index { //compare index of v with last index of Ini Braces
				return false
			}

			sIniIndexBackup = append(sIniIndexBackup, sIniIndexBackup[:len(sIniIndexBackup)-1]...) //delete last sIniIndexBackup value
		}
	}

	if sentinel != 0 {
		return false
	}

	return true
}

func BestSolution(str string) bool {
	var m = map[string]string{
		"{": "}",
		"(": ")",
		"[": "]",
	}

	s := make([]string, 0)
	for _, r := range str {
		r := string(r)
		if len(s) > 0 && m[s[len(s)-1]] == r {
			s = s[:len(s)-1]
		} else {
			s = append(s, r)
		}
	}
	return len(s) == 0
}
