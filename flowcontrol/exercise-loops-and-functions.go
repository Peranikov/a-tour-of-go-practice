package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	var prevz float64
	for i := 0; i < 10; i++ {
		z = z - (z*z-x)/(2*z)
		if math.Abs(z - prevz) < 1e-10 {
			return z
		}

		prevz = z
	}

	return z
}

func main() {
	fmt.Println(Sqrt(1), "/", math.Sqrt(1))
	fmt.Println(Sqrt(2), "/", math.Sqrt(2))
	fmt.Println(Sqrt(3), "/", math.Sqrt(3))
	fmt.Println(Sqrt(4), "/", math.Sqrt(4))
	fmt.Println(Sqrt(5), "/", math.Sqrt(5))
	fmt.Println(Sqrt(9), "/", math.Sqrt(9))
}
