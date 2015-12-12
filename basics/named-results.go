package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // 名前をつけた戻り値の変数を使うと、 return ステートメントに何も書かずに戻すことができます。これを "naked" return と呼びます。
}

func main() {
	fmt.Println(split(17))
}
