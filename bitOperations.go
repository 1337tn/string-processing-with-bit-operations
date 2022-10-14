package bitoperations

import (
	"strings"
)

const (
	UPPER = 1 << iota
	LOWER
	CAP
	REV
)

func ProcessString(str string, conf byte) string {
	// processString changes string str according to configuration conf:
	// UPPER - make str uppercase
	// LOWER - make str lowercase
	// CAP - capitalize first letter of each word in str
	// REV - reverse the letters order in str
	// conf can set multiple parameters as in conf = REV|CAP
	rev := func(s string) string {
		runes := []rune(s)
		n := len(runes)
		for i := 0; i < n/2; i++ {
			runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
		}
		return string(runes)
	}

	title := func(input string) string {
		words := strings.Fields(input)
		smallwords := " a an on the to "

		for index, word := range words {
			if strings.Contains(smallwords, " "+word+" ") {
				words[index] = word
			} else {
				words[index] = strings.Title(word)
			}
		}
		return strings.Join(words, " ")
	}

	if (conf & UPPER) != 0 {
		str = strings.ToUpper(str)
	}
	if (conf & LOWER) != 0 {
		str = strings.ToLower(str)
	}
	if (conf & CAP) != 0 {
		str = title(strings.ToLower(str))
	}
	if (conf & REV) != 0 {
		str = rev(str)
	}
	return str
}

func SameSign(a, b int) bool {
	// sameSign checks whether numbers a and b have the same sign
	return a^b >= 0
}

func SetNthBit(num, pos int) int {
	// setNthBit sets bit on position pos in num
	return num | (1 << (pos - 1))
}

func UnsetNthBit(num, pos int) int {
	// unsetNthBit unsets bit on position pos in num
	return num &^ (1 << (pos - 1))
}

func GetNthBit(num, pos int) int {
	// getNthBit returns bit on position pos in num
	return num & (1 << (pos - 1))
}
