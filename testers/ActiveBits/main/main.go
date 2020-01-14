package main

import (
	"fmt"
	"os"
	"strconv"

	student ".."
)

func main() {
	n, _ := strconv.Atoi(os.Args[1])
	nbits := student.ActiveBits(n)
	fmt.Println(nbits)
}
