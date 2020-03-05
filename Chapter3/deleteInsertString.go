package main

import "fmt"

func deleteInsert(val string) string {
	if len(val) == 0 {
		return val
	}
	arr := []byte(val)
	finalCount := 0
	dummy := []byte{}
	for _, ele := range val {
		if ele == 'a' {
			finalCount += 2
			dummy = append(dummy, []byte{0, 0}...)
		} else if ele == 'b' {
			continue
		} else {
			dummy = append(dummy, 0)
			finalCount++
		}
	}
	if finalCount == 0 {
		return ""
	}
	origLen := len(arr)
	arr = append(arr, dummy...)
	origIndex := origLen - 1
	ele := arr[origIndex]
	fmt.Println(string(arr))
	for i := finalCount; i > 0; {
		if ele == 'a' {
			i--
			arr[i] = 'd'
			origIndex--
			if origIndex >= 0 {
				ele = arr[origIndex]
			}
			i--
			arr[i] = 'd'
		} else if ele == 'b' {
			origIndex--
			if origIndex >= 0 {
				ele = arr[origIndex]
			}
		} else {
			i--
			origIndex--
			var tmp byte
			if origIndex >= 0 {
				tmp = arr[origIndex]
			}
			arr[i] = ele
			ele = tmp
		}
		fmt.Println(string(arr), "---", string(ele))
	}

	return string(arr)
}
