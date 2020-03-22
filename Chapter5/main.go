package main

import (
	"fmt"
	"math"
)

type point struct {
	x int
	y int
}
type rectangle struct {
	ll point
	ur point
}

func init() {
	revArray4bit := [16]uint{0, 8, 4, 12, 2, 10, 6, 14, 1, 9, 5, 13, 3, 11, 7, 15}
	revArray16bit := make([]uint, 65536)

	for i := uint(0); i < 65536; i++ {
		a := revArray4bit[i>>12&15]
		b := revArray4bit[i>>8&15] << 4
		c := revArray4bit[i>>4&15] << 8
		d := revArray4bit[i&15] << 12
		revArray16bit[i] = a | b | c | d
	}
	fmt.Println("Ready")
}

func main() {
	// x := 1.1
	// y := uint64(1000)
	// fmt.Println(pow1(x, y), pow2(x, y), pow3(x, y), pow4(x, y))
	// fmt.Println(parity(24031988))
	// fmt.Println(reverseDigit(24031988))

	rec1 := rectangle{point{1, 1}, point{4, 4}}
	rec2 := rectangle{point{3, 3}, point{5, 5}}

	fmt.Println(boundingRectangle(rec1, rec2))

	rec1 = rectangle{point{1, 2}, point{4, 6}}
	rec2 = rectangle{point{5, 3}, point{7, 7}}

	fmt.Println(boundingRectangle(rec1, rec2))

	return
}

func boundingRectangle(rec1 rectangle, rec2 rectangle) rectangle {
	var boundingRec rectangle
	if rec1.ll.x <= rec2.ll.x && rec2.ll.x <= rec1.ur.x {
		if rec1.ll.y <= rec2.ll.y && rec2.ll.y <= rec1.ur.y {
			boundingRec.ll.x = rec2.ll.x
			boundingRec.ll.y = rec2.ll.y
			boundingRec.ur.x = min(rec1.ur.x, rec2.ur.x)
			boundingRec.ur.y = min(rec1.ur.y, rec2.ur.y)
		}
		if rec2.ll.y <= rec1.ll.y && rec1.ll.y <= rec2.ur.y {
			boundingRec.ll.x = rec2.ll.x
			boundingRec.ll.y = rec1.ll.y
			boundingRec.ur.x = min(rec1.ur.x, rec2.ur.x)
			boundingRec.ur.y = min(rec1.ur.y, rec2.ur.y)
		}
	}
	if rec2.ll.x <= rec1.ll.x && rec1.ll.x <= rec2.ur.x {
		if rec1.ll.y <= rec2.ll.y && rec2.ll.y <= rec1.ur.y {
			boundingRec.ll.x = rec1.ll.x
			boundingRec.ll.y = rec2.ll.y
			boundingRec.ur.x = min(rec1.ur.x, rec2.ur.x)
			boundingRec.ur.y = min(rec1.ur.y, rec2.ur.y)
		}
		if rec2.ll.y <= rec1.ll.y && rec1.ll.y <= rec2.ur.y {
			boundingRec.ll.x = rec1.ll.x
			boundingRec.ll.y = rec1.ll.y
			boundingRec.ur.x = min(rec1.ur.x, rec2.ur.x)
			boundingRec.ur.y = min(rec1.ur.y, rec2.ur.y)
		}
	}
	return boundingRec
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func reverseDigit(num uint) uint {
	result := uint(0)
	for num > 0 {
		lsb := num % 10
		num = num / 10
		result = result*10 + lsb
	}
	return result
}

func parity(num uint) int {
	fmt.Printf("%b\n", num)
	result := 0
	for num > 0 {
		result ^= 1
		num = num & (num - 1)
	}

	return result
}

func pow1(x float64, y uint64) float64 {
	result := 1.0
	power := y
	for power > 0 {
		if (power & 1) == 1 {
			result *= x
		}
		x *= x
		power >>= 1
	}
	return result
}

func pow2(x float64, y uint64) float64 {
	if y == 0 {
		return 1
	}
	if y%2 == 0 {
		return pow2(x, y/2) * pow2(x, y/2)
	}
	return x * pow2(x, y/2) * pow2(x, y/2)

}

func pow3(x float64, y uint64) float64 {
	if y == 0 {
		return 1
	}
	temp := pow3(x, y/2)
	if y%2 == 0 {
		return temp * temp
	}
	return x * temp * temp

}

func pow4(x float64, y uint64) float64 {
	return math.Pow(x, float64(y))
}

func bitReverse(num uint64) {

}
