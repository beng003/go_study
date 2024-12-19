package main

// import "fmt"
import (
	"fmt"
	"go_project/src/gocode/test_project01/unit2/demo17/test"
)

func main() {
	age := 19
	fmt.Printf("My age is %d\n", age)

	num2 := test.StudentAge2
	fmt.Println("num2 = ", num2)

	// num3:= test.studentName
	// fmt.Println("num3 = ", num3)
}
