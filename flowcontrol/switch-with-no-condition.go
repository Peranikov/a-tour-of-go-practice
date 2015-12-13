package main
import (
	"time"
	"fmt"
)

func main() {
	t := time.Now()
	switch { // 条件のないswitchは `switch true` と同様
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
