package palindrome

import (
	"strings"
)

//Return if input string is a palindrome

func textTreatment(rawText string) string {
	text := strings.ToLower(rawText)

	blankChars := []string{" ", ",", "-", ".", ";", ":", "!", "?", "\""}
	eChars := []string{"ê", "é"}
	aChars := []string{"ã", "â", "á"}
	iChars := []string{"í", "î"}
	oChars := []string{"õ", "ó", "ô"}
	uChars := []string{"ú", "û"}
	cChars := []string{"ç"}

	m := map[string][]string{
		"a": aChars,
		"e": eChars,
		"i": iChars,
		"o": oChars,
		"u": uChars,
		"":  blankChars,
		"c": cChars,
	}

	for k, v := range m {
		for _, char := range v {
			text = strings.ReplaceAll(text, char, k)
		}
	}

	return text
}

func swapString(s string) string {
	swap := []rune(s)
	var l int = len(s) - 1

	for i, j := 0, l; i < j; i, j = i+1, j-1 {
		swap[i], swap[j] = swap[j], swap[i]
	}

	return string(swap)
}

func MySolution(text string) bool {
	inputText := textTreatment(text)

	return strings.Compare(inputText, swapString(inputText)) == 0
}
