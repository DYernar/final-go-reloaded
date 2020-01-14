package main

import (
	"fmt"
	"os"

	student ".."
)

func main() {
	s := os.Args[1]
	n := student.Atoi(s)
	fmt.Println(n)
}
