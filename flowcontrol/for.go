package main

import "fmt"

func main() {
	sum := 0
	for i := 10; i > 0; i-- {
		sum += 1
	}
	fmt.Println(sum)
}
