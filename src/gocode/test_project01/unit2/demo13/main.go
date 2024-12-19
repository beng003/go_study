package main

// import "fmt"
import (
	"fmt"
	"strconv"
)

func main() {
	age := 19

	var b string = fmt.Sprintf("%d", age)
	fmt.Printf("My age is %s\n", b)
	fmt.Printf("My age is %v\n", b)

	num2 := 3.1459265358979
	num3 := strconv.FormatFloat(num2, 'f', -1, 64)
	fmt.Println("num3 = ", num3)

}
