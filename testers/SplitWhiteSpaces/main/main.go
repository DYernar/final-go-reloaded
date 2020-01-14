package main

import (
	"fmt"
	"os"

	student ".."
)

func main() {
	str := os.Args[1]
	fmt.Println(student.SplitWhiteSpaces(str))
}
