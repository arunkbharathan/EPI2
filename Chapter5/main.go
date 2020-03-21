package main

import "fmt"

func main() {
	// stackWithMax()
	// solveReversePolishNotation()
	// checkWellFormedness()
	// printBinaryTree()
	circularQImplementation()

}

func circularQImplementation() {
	cr := getNewQ(2)
	cr.Push("1")
	cr.Push("2")
	cr.Push("3")
	cr.Push("4")
	fmt.Println(cr.String())
	fmt.Println(cr.Pop())
	fmt.Println(cr.Pop())
	cr.Push("a")
	cr.Push("b")
	cr.Push("c")
	cr.Push("d")
	fmt.Println(cr.String())
	cr.Pop()
	cr.Pop()
	cr.Pop()
	cr.Pop()
	cr.Pop()
	cr.Pop()
	fmt.Println(cr.String())
	cr.Push("a")
	cr.Push("b")
	cr.Push("c")
	cr.Push("d")
	cr.Push("1")
	cr.Push("2")
	cr.Push("3")
	cr.Push("4")
	cr.Push("h")
	cr.Push("g")
	cr.Push("f")
	cr.Push("e")
	fmt.Println(cr.String())
	cr.Push("AKB")
	cr.Pop()
	cr.Pop()
	cr.Pop()
	cr.Push("cds")
	fmt.Println(cr.String())
	fmt.Println(cr.Cap())
}

// func printBinaryTree() {
// 	fmt.Println(bta)
// }

// func checkWellFormedness() {
// 	str := ")(({,},[,],(,),))"
// 	fmt.Println(testWellFormedness(str))

// }

// func solveReversePolishNotation() {
// 	notation1 := "3,4,+,2,*,1,+"
// 	fmt.Println(solveRPN(notation1))
// 	notation2 := "1,1,+,-2,*,"
// 	fmt.Println(solveRPN(notation2))
// 	notation3 := "-641,6,/,28,/"
// 	fmt.Println(solveRPN(notation3))

// }

// func stackWithMax() {

// 	stack := new()
// 	stack.push(2)
// 	fmt.Println(stack.max())
// 	stack.push(2)
// 	fmt.Println(stack.max())
// 	stack.push(1)
// 	fmt.Println(stack.max())
// 	stack.push(4)
// 	fmt.Println(stack.max())
// 	stack.push(5)
// 	fmt.Println(stack.max())
// 	stack.push(5)
// 	fmt.Println(stack.max())
// 	stack.push(3)
// 	fmt.Println(stack.max())

// 	fmt.Println(stack.pop())
// 	fmt.Println(stack.max())
// 	fmt.Println(stack.pop())
// 	fmt.Println(stack.max())
// 	fmt.Println(stack.pop())
// 	fmt.Println(stack.max())
// 	fmt.Println(stack.pop())
// 	fmt.Println(stack.max())

// 	stack.push(0)
// 	fmt.Println(stack.max())
// 	stack.push(3)
// 	fmt.Println(stack.max())

// }
