package main

import "github.com/golang-collections/collections/stack"

type lifo struct {
	dataStack *stack.Stack
	maxStack  *stack.Stack
}

func new() *lifo {
	return &lifo{
		stack.New(),
		stack.New(),
	}
}

func (st *lifo) len() int {
	length := st.dataStack.Len()
	return length
}

func (st *lifo) peek() int {
	peekval := st.dataStack.Peek().(int)
	return peekval
}

func (st *lifo) pop() int {
	discard := st.maxStack.Pop().(int)
	_ = discard
	val := st.dataStack.Pop().(int)
	return val
}

func (st *lifo) max() int {
	max := st.maxStack.Peek()
	return max.(int)
}

func (st *lifo) push(val int) {
	currentMax := -1000000
	if st.maxStack.Len() > 0 {
		currentMax = st.maxStack.Peek().(int)
	}
	st.dataStack.Push(val)
	if val > currentMax {
		st.maxStack.Push(val)
	} else {
		st.maxStack.Push(currentMax)
	}
	return
}
