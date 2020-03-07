package main

func nameReverse(sent string) string {
	a := []rune(sent)
	//b := make([]rune, len(a))
	j := len(a) - 1
	for i := range a {
		if i < j {
			a[i], a[j] = a[j], a[i]
		}
		j--
	}
	s := 0
	for i := range a {

		switch a[i] {
		case ' ':
			reverse(a, s, i-1)
			s = i + 1
		}
	}
	reverse(a, s, len(a)-1)
	return string(a)
}

func reverse(a []rune, i, j int) {
	d := (j - i) / 2
	for x := i; x <= d+i; x++ {
		a[x], a[j] = a[j], a[x]
		j--
	}
}
