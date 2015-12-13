package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i) // defer へ渡した関数は LIFO(last-in-first-out) の順番で実行される
	}

	fmt.Println("done")
}
