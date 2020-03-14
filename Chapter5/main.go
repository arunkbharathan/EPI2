package main

import "fmt"

func main() {
	stackWithMax()
}

func stackWithMax() {

	stack := new()
	stack.push(2)
	fmt.Println(stack.max())
	stack.push(2)
	fmt.Println(stack.max())
	stack.push(1)
	fmt.Println(stack.max())
	stack.push(4)
	fmt.Println(stack.max())
	stack.push(5)
	fmt.Println(stack.max())
	stack.push(5)
	fmt.Println(stack.max())
	stack.push(3)
	fmt.Println(stack.max())

	fmt.Println(stack.pop())
	fmt.Println(stack.max())
	fmt.Println(stack.pop())
	fmt.Println(stack.max())
	fmt.Println(stack.pop())
	fmt.Println(stack.max())
	fmt.Println(stack.pop())
	fmt.Println(stack.max())

	stack.push(0)
	fmt.Println(stack.max())
	stack.push(3)
	fmt.Println(stack.max())

}
