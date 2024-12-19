package main

// import "fmt"
import (
	"fmt"
)

func main() {
	age := 19

	b := &age
	fmt.Printf("b = %v\n", b)
	fmt.Printf("b = %T\n", b)
	fmt.Printf("b = %v\n", *b)
	fmt.Printf("b = %T\n", &b)

}
