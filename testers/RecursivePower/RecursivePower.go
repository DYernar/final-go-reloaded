package student

func RecursivePower(nb int, power int) int {
	if power < 0 {
		return 0
	}
	if power == 0 {
		return 1
	}
	return nb*RecursivePower(nb, power -1 )
}


////testing ---------------------------------------------
// package main

// import (
// 	"fmt"
// 	student ".."
// )

// func main() {
// 	a := -7
// 	ap := -2
// 	b := -8
// 	bp := -7
// 	c := 4
// 	cp := 8
// 	d := 1
// 	dp := 3
// 	e := -1
// 	ep := 1
// 	f := -6
// 	fp := 5
// 	fmt.Println(student.RecursivePower(a, ap))
// 	fmt.Println(student.RecursivePower(b, bp))
// 	fmt.Println(student.RecursivePower(c, cp))
// 	fmt.Println(student.RecursivePower(d, dp))
// 	fmt.Println(student.RecursivePower(e, ep))
// 	fmt.Println(student.RecursivePower(f, fp))

// }
