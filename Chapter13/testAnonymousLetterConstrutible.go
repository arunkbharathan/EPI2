package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func testForAnonymousLetterConstrutible(magazine, letter string) bool {

	file, err := os.Open(magazine)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	magazineMap := map[rune]int{}
	for fileScanner.Scan() {
		line := fileScanner.Text()
		line = strings.ToLower(line)
		for _, c := range line {
			if c >= 97 && c <= 122 {
				magazineMap[c]++

			}
		}
	}

	file2, err := os.Open(letter)
	if err != nil {
		log.Fatal(err)
	}
	defer file2.Close()

	fileScanner = bufio.NewScanner(file2)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		line = strings.ToLower(line)
		for _, c := range line {
			if c >= 97 && c <= 122 {

				if val, ok := magazineMap[c]; ok {
					if val == 0 {
						return false
					}
					magazineMap[c]--

				} else {
					return false
				}
			}
		}
	}

	return true
}
