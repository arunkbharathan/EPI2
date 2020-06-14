package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	avl "github.com/emirpasic/gods/trees/avltree"
)

func mostVisitedPage(kMostVisited int) {
	// mostVisited := []*mostVisitedWithCount{}
	visitCountIDTree := avl.NewWithIntComparator()
	idVisitCountMap := map[string]int{}
	fs, err := os.Open("./log.txt")
	if err != nil {
		log.Fatalf("Can't open file %s", err)
	}
	lineReader := bufio.NewScanner(fs)

	for lineReader.Scan() {
		line := lineReader.Text()
		for _, chr := range strings.Split(line, " ") {
			if oldCount, ok := idVisitCountMap[chr]; ok {
				if idval, exists := visitCountIDTree.Get(oldCount); exists {
					visitCountIDTree.Remove(oldCount)
					ids := idval.([]string)
					newIds := []string{}
					for _, k := range ids {
						if k != chr {
							newIds = append(newIds, k)
						}
					}
					if len(newIds) != 0 {
						visitCountIDTree.Put(oldCount, newIds)
						// fmt.Println()
					}
					idVisitCountMap[chr]++
					newCount := idVisitCountMap[chr]
					if idval, exists := visitCountIDTree.Get(newCount); exists {
						visitCountIDTree.Remove(newCount)
						ids := idval.([]string)
						newIds := append(ids, chr)
						if len(newIds) != 0 {
							visitCountIDTree.Put(newCount, newIds)
							// fmt.Println()
						}
					} else {
						visitCountIDTree.Put(newCount, []string{chr})
					}

				}
			} else {

				if idval, exists := visitCountIDTree.Get(1); exists {
					visitCountIDTree.Remove(1)
					ids := idval.([]string)
					newIds := append(ids, chr)
					if len(newIds) != 0 {
						visitCountIDTree.Put(1, newIds)
						// fmt.Println()
					}
				} else {
					visitCountIDTree.Put(1, []string{chr})
				}
				idVisitCountMap[chr]++
			}
		}

	}

	for i := kMostVisited; i > 0; i-- {
		item := visitCountIDTree.Right()
		if item != nil {

			fmt.Println(item.Value.([]string), " : ", item.Key.(int))
			visitCountIDTree.Remove(item.Key)
		}
	}
	fmt.Println()
}
