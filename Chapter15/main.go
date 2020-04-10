package main

import "fmt"

func main() {
	checkBSTProperty()
}

func checkBSTProperty(){
	satisfies:=checkBTSatisfiesBSTProperty1(bsTree())
	fmt.Println(satisfies)
	satisfies=checkBTSatisfiesBSTProperty1(btUnbalanced())
	fmt.Println(satisfies)
	satisfies=checkBTSatisfiesBSTProperty2(bsTree())
	fmt.Println(satisfies)
	satisfies=checkBTSatisfiesBSTProperty2(btUnbalanced())
	fmt.Println(satisfies)
}