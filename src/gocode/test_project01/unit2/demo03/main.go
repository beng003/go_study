package main

// import "fmt"
import (
	"fmt"
	"unsafe"
)

func main() {
	age := 18
	fmt.Printf("My age is %d\n", age)
	age2 := 20
	fmt.Printf("My age is %T\n", age2)
	num2 := 3.1459265358979
	fmt.Println(unsafe.Sizeof(num2))
	var num3 float32 = 314e-2
	fmt.Println("num3 = ", num3)
	num4 := 314e2
	fmt.Println("num4 = ", num4)

}
