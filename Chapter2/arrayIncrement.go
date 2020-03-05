package main

func arrayIncrement(arr [4]int) [4]int {
	c := 0
	s := 0
	for i := len(arr) - 1; i >= 0; i-- {
		if i == len(arr)-1 {
			s = arr[i] + 1
		} else {
			s = arr[i] + c
			if c == 0 {
				break
			}
		}
		if s > 9 {
			arr[i] = s % 10
			c = 1
			continue
		}
		c = 0
		arr[i] = s
	}
	return arr
}
