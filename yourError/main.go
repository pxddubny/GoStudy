package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("%v пиздарики", float64(e))
} 

func Sqrt(x float64) (float64, error) {

	z := 1.0

	if x < 0 {

		return 0.0, ErrNegativeSqrt(x)

	}

	for {
		c := z
		z -= (z*z - x) / (2*z)
		
		fmt.Println(z)

		if math.Abs(c - z) < 0.0000000000001{
			break
		}
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
