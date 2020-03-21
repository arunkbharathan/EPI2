package main

import (
	"fmt"
	"strings"
)

type circularQueue struct {
	data  []string
	len   int
	start int
	end   int
}

func getNewQ(cap int) *circularQueue {
	return &circularQueue{
		make([]string, cap, cap),
		0,
		0,
		0,
	}
}
func (cq *circularQueue) Len() int {
	return cq.len
}
func (cq *circularQueue) Cap() int {
	return cap(cq.data)
}
func (cq *circularQueue) Push(str string) {
	if cq.Len() < cq.Cap() {
		cq.data[cq.end] = str
		cq.end++
		cq.len++
		cq.end = cq.end % cq.Cap()
		return
	}
	newData := make([]string, cq.Cap()+2, cq.Cap()+2)
	i := 0
	for ; i < cq.Len(); i++ {
		ind := (i + cq.start) % cq.Cap()
		newData[i] = cq.data[ind]

	}
	cq.start = 0
	cq.end = i
	newData[i] = str
	cq.end++
	cq.len++
	cq.data = newData
	return
}
func (cq *circularQueue) Pop() string {
	if cq.Len() > 0 {
		val := cq.data[cq.start]
		cq.start++
		cq.len--
		return val
	}
	return ""
}

func (cq *circularQueue) String() string {
	str := strings.Builder{}
	for i := 0; i < cq.Len(); i++ {
		ind := (i + cq.start) % cq.Cap()
		str.WriteString(fmt.Sprintf("%s ", cq.data[ind]))

	}
	return str.String()
}
