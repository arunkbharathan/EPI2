package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	avl "github.com/emirpasic/gods/trees/avltree"
)

type mostVisitedWithCount struct {
	page  string
	count int
}

func (item *mostVisitedWithCount) String() string {
	return fmt.Sprintf("Site: %s ; Count: %d\n", item.page, item.count)
}

func mostVisitedPage(kMostVisited int) {
	// mostVisited := []*mostVisitedWithCount{}
	visitCountIDTree := avl.NewWithIntComparator()
	idVisitCountMap := map[string]int{}
	fs, err := os.Open("./log.txt")
	if err != nil {
		log.Fatalf("Can't open file %s", err)
	}
	lineReader := bufio.NewScanner(fs)
	// here solution is to use a hash map and bst to for quick pointing to bst for update and find max and min
	// But the bst I found is not allowing duplicate keys so doing a workaround
	for lineReader.Scan() {
		line := lineReader.Text()
		for _, chr := range strings.Split(line, " ") {
			idVisitCountMap[chr]++
		}

		for id, count := range idVisitCountMap {
			visitCountIDTree.Put(count, id)
		}
		for i := kMostVisited; i > 0; i-- {
			item := visitCountIDTree.Right()
			if item != nil {
				val := mostVisitedWithCount{item.Value.(string), item.Key.(int)}
				fmt.Print(&val)
				visitCountIDTree.Remove(item.Key)
			}
		}
		fmt.Println()
	}
}
