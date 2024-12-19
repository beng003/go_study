package main

// import "fmt"
import (
	"fmt"
)

// func cal(a int, b int) int {
// 	return a + b
// }

func cal2(a float64, b float64) float64 {
	return a + b
}

func main() {
	// a := 10.2
	// b := 20.3
	// c := cal(a, b)
	// fmt.Println("c=", c)

	d := 10.0
	e := 20.0
	f := cal2(d, e)
	fmt.Println("f=", f)
}
