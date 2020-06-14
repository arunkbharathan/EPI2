package main

import (
	"fmt"

	"github.com/golang-collections/collections/stack"
)

func testWellFormedness(str string) bool {
	stk := stack.New()
	for _, c := range str {
		switch c {
		case '(', '[', '{':
			stk.Push(1)
		case ')', ']', '}':
			if stk.Len() > 0 {
				stk.Pop()
			} else {
				return false
			}
		default:
			fmt.Print()
		}
	}
	if stk.Len() > 0 {
		return false
	}
	return true
}
