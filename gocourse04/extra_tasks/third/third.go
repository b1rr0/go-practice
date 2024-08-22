package main

import (
	"fmt"
)

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, fmt.Errorf("square root of negative number")
	}

	if x == 0 || x == 1 {
		return x, nil
	}

	start := float64(0)
	end := x

	for end-start > 0.00000000000001 {
		mid := (start + end) / 2
		sq := mid * mid

		if sq == x {
			return mid, nil
		}

		if sq < x {
			start = mid
		} else {
			end = mid
		}
	}

	return (start + end) / 2, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(9))
	fmt.Println(Sqrt(-2))
}
