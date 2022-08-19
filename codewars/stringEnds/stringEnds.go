package stringEnds

// challenge: Complete the solution so that it returns true if the first argument(string) passed in ends with the 2nd argument (also a string).

import "strings"

func mySolution(str, ending string) bool {
	if str == ending {
		return true
	}
	if len(str) == 0 || len(ending) > len(str) {
		return false
	}

	ref := len(str) - 1
	for i := len(ending) - 1; i >= 0; i-- {
		if ending[i] == str[ref] {
			ref--
			continue
		}
		return false
	}
	return true
}

func bestSolution(str, ending string) bool {
	return strings.HasSuffix(str, ending)
}
