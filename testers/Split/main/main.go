package main

import (
	"fmt"
	"os"

	student ".."
)

func main() {

	str := os.Args[1]
	str2 := os.Args[2]
	fmt.Println(student.Split(str, str2))
}
