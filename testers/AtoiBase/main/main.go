package main

import (
	"fmt"
	"os"

	student ".."
)

func main() {
	word := os.Args[1]
	base := os.Args[2]
	fmt.Println(student.AtoiBase(word, base))
}
