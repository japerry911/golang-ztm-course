package main

import "strings"

func isPalindromeHelper(lhs, rhs string) bool {
	ri := len(lhs) - 1
	for li := 0; li < len(lhs); li++ {
		if !strings.EqualFold(string(lhs[li]), string(rhs[ri])) {
			return false
		}
		ri--
	}

	return true
}

func IsPalindrome(str string) bool {
	if len(str) < 2 {
		return true
	} else if len(str) == 2 {
		return str[0] == str[1]
	} else if len(str)%2 == 1 {
		lhs := len(str) / 2
		return isPalindromeHelper(str[:lhs], str[lhs+1:])
	} else {
		return isPalindromeHelper(str[:len(str)/2], str[len(str)/2:])
	}
}
