package main

import (
	"os"
	"strconv"

	student ".."
)

func main() {
	n, _ := strconv.Atoi(os.Args[1])
	student.PrintCombN(n)

}
