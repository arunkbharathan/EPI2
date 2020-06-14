package main

func baseConversion(val string, fromBase, toBase int) string {
	inDecimal := 0
	for _, item := range val {

		switch {
		case item >= 'A' && item <= 'Z':
			item = item - 'A' + 10
		case item >= '0' && item <= '9':
			item = item - '0'
		}
		inDecimal = inDecimal*fromBase + int(item)
	}

	intoBase := toBaseB2(inDecimal, toBase)

	return intoBase
}

func toBaseB2(num int, base int) string {

	if num == 0 {
		return ""
	}

	return toBaseB2(num/base, base) + digitHash(num%base)

}

func digitHash(item int) string {

	switch {
	case item >= 0 && item <= 9:
		return string('0' + rune(item))
	case item > 9:
		return string('A' + rune(item))
	}
	return ""
}
