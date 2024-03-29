package duplicate_encoder

/* The goal of this exercise is to convert a string to a new string where each character in the new string is "(" if that character appears only once in the original string, or ")" if that character appears more than once in the original string. Ignore capitalization when determining if a character is a duplicate.

Examples
	"din"      =>  "((("
	"recede"   =>  "()()()"
	"Success"  =>  ")())())"
	"(( @"     =>  "))(("  		*/

func isDuplicate(n rune, word string) (bool, []int) {
	index := make([]int, 0)

	for i, v := range word {
		if v == n {
			index = append(index, i)
		}
	}

	if len(index) > 1 {
		return true, index
	}
	return false, index
}

func MySolution(word string) string {
	var resString string
	res := make([]bool, len(word))

	for i, v := range word {
		if cont, index := isDuplicate(v, word); cont {
			for _, v := range index {
				res[v] = true
			}
		} else {
			res[i] = false
		}
	}

	for _, v := range res {
		if v {
			resString += ")"
		} else {
			resString += "("
		}
	}
	return resString
}
