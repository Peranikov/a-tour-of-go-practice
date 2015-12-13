package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else { // ifのステートメントで宣言した変数はelse内でも使える
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func main() {
//	if a := 1; true {
//		return a
//	}

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
