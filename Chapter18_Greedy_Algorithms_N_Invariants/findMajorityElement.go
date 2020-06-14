package main

func findMajorityElement(str string) string {

	count := 0
	var chr string

	for _, ele := range str {
		tmp := string(ele)
		if count == 0 {
			chr = tmp
			count = 1
		} else if chr == tmp {
			count++
		} else {
			count--
		}

	}
	return chr
}
