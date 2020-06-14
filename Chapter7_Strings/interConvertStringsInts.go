package main

import (
	"fmt"
	"math"
)

func interConvert(val interface{}) interface{} {
	switch num := val.(type) {
	case int:
		digits := int(math.Floor(math.Log10(float64(num))))
		result := ""
		for i := int(math.Pow10(digits)); i >= 1; i /= 10 {
			e := num / i
			result = result + fmt.Sprintf("%d", e)
			num -= e * i
		}
		return result
	case string:
		result := 0
		for _, val := range num {

			result = result*10 + int(val) - 48
		}
		return result
	default:
		fmt.Printf("I don't know to handle type %T!\n", num)
		return nil
	}
}
