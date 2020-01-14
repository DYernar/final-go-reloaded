package student


func AtoiBase(s string, base string) int {
	for _, item := range s {
		contains := false
		for _, innerItem := range base {
			if innerItem == item {
				contains = true
			}
		}
		if contains == false {
			return 0
		}
	}
	for index, item := range base {
		for indec, innerItem := range base{
			if index == indec {
				continue
			}
			if innerItem == item {
				return 0
			}
			if item == '-' {
				return 0
			} 
		}
	}
		lenPower := len(s)-1
		baseLength := len(base)
		retVal := 0
		indec := 0
		for _, item := range s {
			for ind, innerItem := range base {
				if item == innerItem {
					indec = ind
				}
			}
			retVal = retVal + (indec*RecursivePower(baseLength, lenPower))
			lenPower--
		}
		return retVal
}


// package main

// import (
// 	"fmt"
// 	student ".."
// )

// func main() {
// 	fmt.Println(student.AtoiBase("125", "0123456789"))
// 	fmt.Println(student.AtoiBase("1111101", "01"))
// 	fmt.Println(student.AtoiBase("7D", "0123456789ABCDEF"))
// 	fmt.Println(student.AtoiBase("uoi", "choumi"))
// 	fmt.Println(student.AtoiBase("bbbbbab", "-ab"))
// 	fmt.Println()
// 	fmt.Println(student.AtoiBase("bcbbbbaab", "abc"))
// 	fmt.Println(student.AtoiBase("0001", "01"))
// 	fmt.Println(student.AtoiBase("00", "01"))
// 	fmt.Println(student.AtoiBase("saDt!I!sI", "CHOUMIisDAcat!"))
// 	fmt.Println(student.AtoiBase("AAho?Ao", "WhoAmI?"))
// 	fmt.Println(student.AtoiBase("thisinputshouldnotmatter", "abca"))
// 	fmt.Println(student.AtoiBase("125", "0123456789"))
// 	fmt.Println(student.AtoiBase("bbbbbab", "-ab"))
// }
