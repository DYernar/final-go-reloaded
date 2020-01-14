package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if length(os.Args) == 4 {
		Doop(os.Args)
	}
}

func Doop(args []string) {
	isNum1, positive1 := isNumeric(args[1])
	isNum2, positive2 := isNumeric(args[3])

	if isOperand(args[2]) && isNum1 && isNum2 {
		if !isOverflow(positive1, positive2, args[1], args[3]) {
			num1 := strToInt(args[1], positive1)
			num2 := strToInt(args[3], positive2)
			operand := args[2]
			if !isResultOverflow(num1, num2, operand) {
				res, err := makeOpeartion(num1, num2, operand)
				if err {
					checkErr(err, operand)
				} else {
					_, finalStr := intToStr(res, "")
					printStr(finalStr)
				}
			} else {
				z01.PrintRune('0')
				z01.PrintRune(10)
			}
		} else {
			z01.PrintRune('0')
			z01.PrintRune(10)
		}
	} else {
		z01.PrintRune('0')
		z01.PrintRune(10)
	}
}

func intToStr(res int, s string) (int, string) {
	if res < 0 {
		z01.PrintRune('-')
		res *= -1
	}
	if res/10 == 0 {
		s = string(res%10+48) + s
		return res, s
	}

	s = string(res%10+48) + s
	return intToStr(res/10, s)
}

func checkErr(err bool, operand string) {

	if operand == "/" {
		strErr := "No division by 0"
		printStr(strErr)
	} else {
		strErr := "No remainder of division by 0"
		printStr(strErr)
	}

}

func printStr(s string) {
	for _, w := range s {
		z01.PrintRune(w)
	}
	z01.PrintRune(10)
}

func makeOpeartion(n1, n2 int, operand string) (int, bool) {
	res := 0
	err := false

	if operand == "+" {
		res = n1 + n2
	} else if operand == "-" {
		res = n1 - n2
	} else if operand == "*" {
		res = n1 * n2
	} else if operand == "/" {
		if n2 == 0 {
			err = true
		} else {
			res = n1 / n2
		}
	} else if operand == "%" {
		if n2 == 0 {
			err = true
		} else {
			res = n1 % n2
		}
	}
	return res, err
}

func isResultOverflow(num1, num2 int, operand string) bool {
	max := 9223372036854775807
	min := -9223372036854775808

	if num1 > 0 && num2 > 0 {
		if operand == "+" {
			if num1 > max-num2 {
				return true
			}
		} else if operand == "*" {
			if num1 > max/num2 {
				return true
			}
		}
	} else if num1 < 0 && num2 < 0 {
		if operand == "+" {
			if num1 < min-num2 {
				return true
			}
		} else if operand == "*" {
			if num1 < min/(num2*-1) {
				return true
			}
		}
	} else if num1 > 0 && num2 < 0 {
		if operand == "-" {
			if num1 > max+num2 {
				return true
			}
		} else if operand == "*" {
			if num1*(-1) < min/(num2*-1) {
				return true
			}
		}
	} else if num1 < 0 && num2 > 0 {
		if operand == "-" {
			if num1 < min+num2 {
				return true
			}
		} else if operand == "*" {
			if num1 < min/num2 {
				return true
			}
		}
	}
	return false
}

func strToInt(s string, positive bool) int {
	if !positive {
		s = s[1:]
	}
	runeSlice := []rune(s)
	res := 0

	for _, w := range runeSlice {
		res = res*10 + int(w) - 48
	}
	if !positive {
		res *= -1
	}
	return res
}

func isOverflow(positive1, positive2 bool, num1, num2 string) bool {
	if positive1 {
		if maxOverflow(num1) {
			return true
		}
	} else {
		if minOverflow(num1) {
			return true
		}
	}
	if positive2 {
		if maxOverflow(num2) {
			return true
		}
	} else {
		if minOverflow(num2) {
			return true
		}
	}
	return false
}

func minOverflow(num string) bool {
	min := "-9223372036854775808"
	numR := []rune(num)
	numSize := runeSize(numR)

	if numSize > 20 {
		return true
	} else if numSize == 20 && num > min {
		return true
	}
	return false
}
func maxOverflow(num string) bool {
	max := "9223372036854775807"
	numR := []rune(num)
	numSize := runeSize(numR)

	if numSize > 19 {
		return true
	} else if numSize == 19 && num > max {
		return true
	}
	return false
}
func isNumeric(str string) (bool, bool) {
	runeSlice := []rune(str)
	positive := true

	rSize := runeSize(runeSlice)

	if rSize > 1 {
		if runeSlice[0] == '-' {
			runeSlice = runeSlice[1:]
			positive = false
		}
	}

	for _, w := range runeSlice {
		if !(w >= '0' && w <= '9') {
			return false, positive
		}
	}
	return true, positive

}

func isOperand(str string) bool {
	if !(str == "+" || str == "-" || str == "*" || str == "/" || str == "%") {
		return false
	}
	return true
}
func length(s []string) int {
	count := 0
	for range s {
		count++
	}
	return count
}

func runeSize(str []rune) int {
	length := 0

	for range str {
		length++
	}
	return length
}