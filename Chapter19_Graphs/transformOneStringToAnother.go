package main

type nodeData struct {
	word  string
	count int
}

func transformOneStringToAnother(dict map[string]struct{}, str1, str2 string) int {
	queue := []nodeData{{str1, 0}}
	delete(dict, str1)
	for len(queue) > 0 {
		strData := queue[0]
		if strData.word == str2 {
			return strData.count
		}
		queue = queue[1:]
		strBytes := []byte(strData.word)
		for i := 0; i < len(strBytes); i++ {
			for j := 0; j < 26; j++ {
				strBytes[i] = 'a' + byte(j)
				tmpStr := string(strBytes)
				if _, ok := dict[tmpStr]; ok {
					queue = append(queue, nodeData{tmpStr, strData.count + 1})
					delete(dict, tmpStr)
				}
			}
			strBytes = []byte(strData.word)
		}
	}
	return -1
}
