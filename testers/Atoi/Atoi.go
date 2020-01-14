package student

func Atoi(s string) int {
	number := 0
	isNegative := false

	for index, str := range s {
		if index == 0 {
			if str == '-' {
				isNegative = true
				continue
			} else if str == '+' {
				isNegative = false
				continue
			}
		}
		if str >= '0' && str <= '9' {
			number = number*10 + int(rune(str)-'0')
		} else {
			number = 0
			break
		}
	}
	if isNegative {
		return number * -1
	}
	return number
}

///Testing--------------------------------------

// s := "12345"
// s2 := "0000000012345"
// s3 := "012 345"
// s4 := "Hello World!"
// s5 := "+1234"
// s6 := "-1234"
// s7 := "++1234"
// s8 := "--1234"

// n := student.Atoi(s)
// n2 := student.Atoi(s2)
// n3 := student.Atoi(s3)
// n4 := student.Atoi(s4)
// n5 := student.Atoi(s5)
// n6 := student.Atoi(s6)
// n7 := student.Atoi(s7)
// n8 := student.Atoi(s8)

// fmt.Println(n)
// fmt.Println(n2)
// fmt.Println(n3)
// fmt.Println(n4)
// fmt.Println(n5)
// fmt.Println(n6)
// fmt.Println(n7)
// fmt.Println(n8)