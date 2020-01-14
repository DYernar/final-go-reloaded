package student

import(
	"github.com/01-edu/z01"
)


func PrintNbrBase(num int, base string) {
	retVal := ""
	length := 0
	for range base {
		length++
	}
	mx_p := length
	if num < 0 {
		retVal = "-"
		mx_p *= -1
	}
	if length > 1 {

		for num/mx_p >= length {
			mx_p *= length
		}
		for mx_p != 0 {
			retVal = retVal + string(base[num / mx_p])
			num = num - num / mx_p * mx_p
			mx_p /= length
		}
		symbols := map[rune]bool{}
		for _, c := range base {
			if c == '+' || c == '-' {
				retVal = "NV"
				break
			}
			if symbols[c] == false {
				symbols[c] = true
			} else {
				retVal = "NV"
				break
			}
		}
	} else {
		retVal = "NV"
	}
	for _, c := range retVal {
		z01.PrintRune(c)
	}
}

// student.PrintNbrBase(125, "0123456789")
// z01.PrintRune('\n')
// student.PrintNbrBase(-125, "01")
// z01.PrintRune('\n')
// student.PrintNbrBase(125, "0123456789ABCDEF")
// z01.PrintRune('\n')
// student.PrintNbrBase(-125, "choumi")
// z01.PrintRune('\n')
// student.PrintNbrBase(125, "aa")
// z01.PrintRune('\n')