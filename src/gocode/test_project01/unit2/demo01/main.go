package main

import "fmt"

func main() {
	var age int
	age = 18

	fmt.Println("My age is", age)

	var age2 int = 20
	fmt.Println("My age is", age2)

	var num2 float32
	fmt.Println("num2 = ", num2)

	var num3 = "hello"
	fmt.Println("num3 = ", num3)

	num4 := "男"
	fmt.Println("num4 = ", num4)

	var n1, n2, n3 int

	n4, n5, n6 := 1, 2, 3
	fmt.Println("n1 = ", n1, "n2 = ", n2, "n3 = ", n3)
	fmt.Println("n4 = ", n4, "n5 = ", n5, "n6 = ", n6)

	var (
		name   string = "Tom"
		age3   int    = 18
		isPass bool   = true
	)
	fmt.Println("name = ", name, "age3 = ", age3, "isPass = ", isPass)
}
