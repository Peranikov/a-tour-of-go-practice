package main

import "fmt"

func main() {
	defer fmt.Println("world") // 呼び出し元の関数の終わり(returnする)まで遅延させる
	fmt.Println("hello")
}
