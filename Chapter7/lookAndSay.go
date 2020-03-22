package main

import (
	"bytes"
	"fmt"
)

func lookAndSay(nth int) string {

	return coreLogic("1", nth)

}

func coreLogic(phrase string, nth int) string {
	if nth == 0 {
		return phrase
	}
	nth--
	var pre rune
	var count int
	var nextSay bytes.Buffer
	for ind, val := range phrase {
		if val != pre {
			str := ""
			if ind != 0 {
				str = fmt.Sprintf("%d%s", count, string(pre))
				nextSay.WriteString(str)
				count = 0
			}
		}
		count++
		pre = val
	}
	str := fmt.Sprintf("%d%s", count, string(pre))
	nextSay.WriteString(str)
	return coreLogic(nextSay.String(), nth)
}
