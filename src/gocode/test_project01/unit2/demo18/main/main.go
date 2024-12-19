package main

// import "fmt"
import (
	"fmt"
)

func main() {
	var age int
	fmt.Print("请输入用户年龄:")
	fmt.Scanln(&age)
	fmt.Printf("用户的年龄是:%d\n", age)
}
