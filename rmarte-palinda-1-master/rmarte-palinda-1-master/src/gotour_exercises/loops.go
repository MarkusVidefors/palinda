package main

import (
	"fmt"
)

const small = 0.000000000000001

func Sqrt(x float64) float64 {
	z := 1.0
	for {
		diff := (z*z - x) / (2 * z)
		z -= diff
		fmt.Println(z)
		if diff > -small && diff < small {
			break
		}
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}

