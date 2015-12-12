package main

import "fmt"

func main() {
	var i, j int = 1, 2
	k := 3 // var 宣言の代わりに、短い := の代入文を使い、暗黙的な型宣言ができます
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
