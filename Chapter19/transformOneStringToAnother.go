package main

import "fmt"

type nodeData struct {
	word  string
	count int
}

func transformOneStringToAnother(dist map[string]struct{}, str1, str2 string) int {
	queue := []nodeData{{str1, 0}}

	for len(queue) > 0 {
		strData := queue[0]
		queue = queue[1:]
		strBytes := []byte(strData.word)
		for i := 0; i < len(strBytes); i++ {
			for j := 0; j < 26; j++ {
				strBytes[i] = 'a' + byte(j)
				fmt.Println(string(strBytes))
			}
			strBytes = []byte(strData.word)
		}
	}
	return -1
}
