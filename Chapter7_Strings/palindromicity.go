package main

import "strings"

func isPalindrome(str string) bool {
	str = strings.ToLower(str)
	j := len(str) - 1
	for i := 0; i < j; i++ {

		if !isAlphaNumeric(str[i]) {
			continue
		}
		if !isAlphaNumeric(str[j]) {
			j--
			i--
			continue
		}

		if str[i] != str[j] {
			return false
		}
		j--
	}
	return true
}

func isAlphaNumeric(chr byte) bool {
	if chr >= '0' && chr <= '9' || chr >= 'A' && chr <= 'Z' || chr >= 'a' && chr <= 'z' {
		return true
	}
	return false
}
