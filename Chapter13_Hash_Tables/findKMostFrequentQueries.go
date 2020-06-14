package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func findFrequentQueries(k int) []*string {
	strMap := map[string]int{}
	result := []*string{}
	h := &stringCounterHeap{}
	heap.Init(h)
	file, err := os.Open("./queriesFile")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		line = strings.ToLower(line)
		line = strings.TrimSpace(line)
		//remove common commands
		if !strings.Contains(line, "ls") {
			strMap[line]++
		}
		// strMap[line]++
	}

	for key, count := range strMap {
		// fmt.Println(key)
		if h.Len() >= k {
			minItem := heap.Pop(h).(*stringCounter)
			if count >= minItem.StrCount {
				val := key
				heap.Push(h, &stringCounter{&val, count})
			} else {
				heap.Push(h, minItem)
			}
		} else {
			val := key
			item := stringCounter{&val, count}
			heap.Push(h, &item)
		}
	}
	sort.Sort(h)
	for h.Len() != 0 {
		item := heap.Pop(h).(*stringCounter)
		vl := fmt.Sprintf("%d: %s", item.StrCount, *item.Str)
		result = append(result, &vl)
	}
	fmt.Println()
	return result
}
