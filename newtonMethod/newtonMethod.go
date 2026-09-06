package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {

	z := 1.0

	for {
		c := z
		z -= (z*z - x) / (2*z)
		
		fmt.Println(z)

		if math.Abs(c - z) < 0.0000000000001{
			break
		}
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
