package main
import "fmt"

func main() {
	sum := 1
	for sum < 1000 { // whileはforで表せる
		sum += sum
	}
	fmt.Println(sum)
}
