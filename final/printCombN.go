package student

import(
	z01 "github.com/01-edu/z01"
)




func PrintCombN(r int) {
	a := '0'
	arr := make([]rune, 10)
	for i := 0; i < 10; i++ {
		arr[i] = a
		a++
	}
	data := make([]rune, r)
	var finalArr []string
	combinationUtil(&finalArr, arr, data, 0, 9, 0, r)
	for index, item := range finalArr {
		if index != 0 {
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
		for _, innerItm := range item {
			z01.PrintRune(innerItm)
		}
	}
	z01.PrintRune('\n')
}


func combinationUtil(finalArr *[]string, arr []rune, data []rune, start int, end int, index int, r int) {
	if index == r {
		str := ""
		for j := 0; j < r; j++ {
			str = str+string(data[j])
		}
		*finalArr = append(*finalArr, str)
		return
	}
	for i := start; i <=end && end-i+1 >= r-index; i++ {
		data[index] = arr[i]
		combinationUtil(finalArr, arr, data, i+1, end, index+1, r)
	}

}

///Testing

// package main

// import(
// 	student ".."
// 	"fmt"
// ) 

// func main() {
// 	for i:=1; i < 10; i++ {
// 		fmt.Print("n: ")
// 		fmt.Println(i)
// 		student.PrintCombN(i)
// 		fmt.Println()
// 	}
// }
