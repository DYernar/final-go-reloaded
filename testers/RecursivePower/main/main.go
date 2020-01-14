package main

import (
	"fmt"
	"os"
	"strconv"

	student ".."
)

func main() {
	arg1, _ := strconv.Atoi(os.Args[1])
	arg2, _ := strconv.Atoi(os.Args[2])
	fmt.Println(student.RecursivePower(arg1, arg2))
}
