package main

func findMissingAndDuplicate(arr []uint) (uint, uint) {
	properSum := (0)
	arraySum := (0)
	properXor := uint(0)
	arrayXor := uint(0)
	possibleDuplicate := uint(0)
	possibleDelete := uint(0)
	possibleNum1 := uint(0)
	possibleNum2 := uint(0)
	for i, val := range arr {

		properSum += (i)
		properXor ^= uint(i)
		arraySum += int(val)
		arrayXor ^= uint(val)
	}
	tminusm := arraySum - properSum
	txorm := properXor ^ arrayXor

	/* Get the rightmost set bit in set_bit_no (consider only 1 bit not all)*/
	setBit := txorm & ^(txorm - 1)

	for i := uint(0); i < uint(len(arr)); i++ {
		//xor prpoperArr and givenArr where setBit is set.
		if arr[i]&setBit > 0 {
			possibleNum1 ^= arr[i]
		} else {
			possibleNum2 ^= arr[i]
		}
		if i&setBit > 0 {
			possibleNum1 ^= i
		} else {
			possibleNum2 ^= i
		}
	}
	if tminusm < 0 {
		possibleDuplicate = possibleNum1
		possibleDelete = possibleNum2
	} else {
		possibleDuplicate = possibleNum2
		possibleDelete = possibleNum1
	}
	return possibleDelete, possibleDuplicate
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
