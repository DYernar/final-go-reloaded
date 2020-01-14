package main

import (
	"fmt"
	"os"

	student ".."
)

func main() {
	arg1 := os.Args[1]
	arg2 := os.Args[2]
	arg3 := os.Args[3]
	result := student.ConvertBase(arg1, arg2, arg3)
	fmt.Println(result)
}
