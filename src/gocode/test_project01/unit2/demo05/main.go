package main

// import "fmt"
import (
	"fmt"
)

func main() {
	age := 'a'
	fmt.Printf("My age is %d\n", age)
	age2 := 'b'
	fmt.Printf("My age is %T\n", age2)
	num2 := '中'
	fmt.Println(num2)
	fmt.Printf("num2 = %c\n", num2)
	fmt.Printf("num2 = %T\n", num2)
	num3 := "中国"
	fmt.Println(num3)
	fmt.Printf("num3 = %s\n", num3)
	fmt.Printf("num3 = %T\n", num3)

}
