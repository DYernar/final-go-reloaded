package student


func SplitWhiteSpaces(str string) []string {
	var arr []string
	var indices []int 
	if string(str[0]) != " " {
		indices = append(indices, -1)
	}
	for index, item := range str {
		if string(item) == " " {
			indices = append(indices, index-1)
		}
	}
	indices = append(indices, len(str)-1)
	for i := 0; i < len(indices)-1; i++ {
		tempStr := ""
		for j := indices[i]+1; j < indices[i+1]+1; j++ {
			if str[j] != ' ' {
				tempStr = tempStr+string(str[j])
			}
		}
		if tempStr != ""{
			arr = append(arr, tempStr)

		}
	}

	return arr
}


// package main

// import (
// 	"fmt"
// 	student ".."
// )

// func main() {
// 	str := "Hello how are you?"
// 	str1 := "The earliest foundations of what would become computer science predate the invention of the modern digital computer"
// 	str2 := "Machines for calculating fixed numerical tasks such as the abacus have existed since antiquity,"
// 	str3 := "aiding in computations such as multiplication and division ."
// 	str4 := "Algorithms for performing computations have existed since antiquity, even before the development of sophisticated computing equipment."
// 	str5 := "Wilhelm Schickard designed and constructed the first working mechanical calculator in 1623.[4]"
// 	str6 := "In 1673, Gottfried Leibniz demonstrated a digital mechanical calculator,"
// 	fmt.Println(student.SplitWhiteSpaces(str))
// 	fmt.Println(student.SplitWhiteSpaces(str1))
// 	fmt.Println(student.SplitWhiteSpaces(str2))
// 	fmt.Println(student.SplitWhiteSpaces(str3))
// 	fmt.Println(student.SplitWhiteSpaces(str4))
// 	fmt.Println(student.SplitWhiteSpaces(str5))
// 	fmt.Println(student.SplitWhiteSpaces(str6))
// }
