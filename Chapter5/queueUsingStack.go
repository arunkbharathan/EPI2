package main

import "github.com/golang-collections/collections/stack"

type queueS struct {
	pushStack *stack.Stack
	popStack  *stack.Stack
}

func getQInstance() *queueS {
	return &queueS{
		stack.New(),
		stack.New(),
	}
}
func (cq *queueS) Len() int {
	return cq.pushStack.Len() + cq.popStack.Len()
}

func (cq *queueS) Push(str string) {
	cq.pushStack.Push(str)
}
func (cq *queueS) Pop() string {
	switch {
	case cq.popStack.Len() > 0:
		return cq.popStack.Pop().(string)
	case cq.pushStack.Len() > 0:
		slen := cq.pushStack.Len()
		for i := 0; i < slen; i++ {
			cq.popStack.Push(cq.pushStack.Pop())
		}
		return cq.popStack.Pop().(string)
	case cq.pushStack.Len() == 0:
		return ""
	}
	return ""
}
