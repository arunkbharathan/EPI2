package main

var keyMap = map[rune]string{
	'0': "0",
	'1': "1",
	'2': "ABC",
	'3': "DEF",
	'4': "GHI",
	'5': "JKL",
	'6': "MNO",
	'7': "PQRS",
	'8': "TUV",
	'9': "WXYZ",
}

func generateAllKeyCodes(num string, index int) []string {
	if index == len(num) {
		return []string{num}
	}
	ind := num[index]
	alphas := keyMap[rune(ind)]

	output := []string{}
	byteArr := []byte(num)
	for _, ele := range alphas {
		byteArr[index] = byte(ele)
		strSlices := generateAllKeyCodes(string(byteArr), index+1)
		output = append(output, strSlices...)
	}
	// fmt.Println((output))
	return output
}
