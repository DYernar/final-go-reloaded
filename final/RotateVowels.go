package main


import (

	z01 "github.com/01-edu/z01"
	"os"

)

func main()  {
	args := os.Args
	var str string
	length := 0
	for index, item := range args {
		if index != 0 {
			if index == 1{
				str = item
			} else {
				str = str + " " + item
			}
		}
	}

	for _, item := range str {
		if item == 'A' || item == 'E' || item == 'I' || item == 'O' || item == 'U' || item == 'a' || item == 'e' || item == 'i' || item == 'o' || item == 'u' {
			length++
		}
	}
	arr := make([]rune, length)
	i := 0
	for _, item := range str {
		if item == 'A' || item == 'E' || item == 'I' || item == 'O' || item == 'U' || item == 'a' || item == 'e' || item == 'i' || item == 'o' || item == 'u' {
			arr[i] = item
			i++
		}
	}
	j := length-1
	for _, item := range str {
		if item == 'A' || item == 'E' || item == 'I' || item == 'O' || item == 'U' || item == 'a' || item == 'e' || item == 'i' || item == 'o' || item == 'u' {
			z01.PrintRune(arr[j])
			j--
		} else {
	 		z01.PrintRune(item)
		}
	}
	z01.PrintRune('\n')

}

// ./rotatevowels "HEllO World" "problem solved"