package main

import "math"

func primesUpto(num int) []int {
	if num < 2 {
		return []int{}
	} else if num < 3 {
		return []int{2}
	} else if num < 5 {
		return []int{2, 3}
	}

	primes := primesUpto(int(math.Sqrt(float64(num))))
	retPrime := []int{}

	for i := 2; i <= num; i++ {
		retPrime = append(retPrime, i)
	}

	for _, val := range primes {
		for i := 4; i < len(retPrime); i++ {
			if val == retPrime[i] {
				continue
			}
			if retPrime[i]%val == 0 {
				retPrime[i] = 0
			}

		}
	}
	newList := []int{}
	for _, val := range retPrime {
		if val != 0 {
			newList = append(newList, val)
		}
	}
	return newList

}
