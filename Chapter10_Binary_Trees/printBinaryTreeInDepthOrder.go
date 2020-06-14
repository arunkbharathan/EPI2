package main

import (
	"container/list"
	"fmt"
	"strings"
)

func (bt *binaryTree) String() string {
	li := list.New()
	li.PushBack(bt)
	var str = &strings.Builder{}
	retstr := printFromQ(li, str)
	return retstr
}

func printFromQ(li *list.List, str *strings.Builder) string {
	cnt := li.Len()
	if cnt == 0 {
		return str.String()
	}
	for i := 0; i < cnt; i++ {
		val := li.Front()
		bn := val.Value.(*binaryTree)
		li.Remove(val)
		if bn != nil {
			// str.WriteString(fmt.Sprintf("%d ", bn.data))
			str.WriteString(fmt.Sprintf("%s ", bn.str))
			li.PushBack(bn.left)
			li.PushBack(bn.right)
		}
	}
	str.WriteString(fmt.Sprintln())
	retval := printFromQ(li, str)
	return retval
}
