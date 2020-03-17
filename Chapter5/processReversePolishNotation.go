package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/golang-collections/collections/stack"
)

func solveRPN(rpn string) string {
	stck := stack.New()
	rpnstr := strings.Split(rpn, ",")
	for i := len(rpnstr) - 1; i >= 0; i-- {
		// fmt.Println((rpnstr[i]))
		stck.Push((rpnstr[i]))

	}
	val := reduceRPN(stck)
	return val
}

func reduceRPN(stck *stack.Stack) string {
	if stck.Len() >= 3 {
		numA := stck.Pop().(string)
		numB := stck.Pop().(string)
		operator := stck.Pop().(string)
		num1, err := strconv.Atoi(numA)
		chkError(err)
		num2, err := strconv.Atoi(numB)
		chkError(err)
		switch operator[0] {
		case '+':
			stck.Push(fmt.Sprintf("%d", num1+num2))
		case '-':
			stck.Push(fmt.Sprintf("%d", num1-num2))
		case '*':
			stck.Push(fmt.Sprintf("%d", num1*num2))
		case '/':
			stck.Push(fmt.Sprintf("%d", num1/num2))
		default:
			chkError(fmt.Errorf("Unknown Symbol %s", operator))
		}
		return reduceRPN(stck)
	}
	return stck.Pop().(string)
}

func chkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
