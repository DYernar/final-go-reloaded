package main

import (
	"os"
	"strconv"

	student ".."
	"github.com/01-edu/z01"
)

func main() {
	nbr, _ := strconv.Atoi(os.Args[1])
	base := os.Args[2]
	student.PrintNbrBase(nbr, base)
	z01.PrintRune('\n')
}
