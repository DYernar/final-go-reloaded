package student


func printNbrBase(nbr int, base string) string {
		negative := false
		if nbr < 0 {
			negative = true
			nbr = -nbr
		}
	
		var retStr string
		for nbr >= len(base) {
			remainder := nbr%len(base)
			retStr = string(base[remainder]) + retStr
			nbr = nbr/len(base)
		}
		retStr = string(base[nbr]) + retStr
		if negative {
			retStr = "-" + retStr
		}
		return retStr
}

func ConvertBase(nbr, baseFrom, baseTo string) string {
	convertedNum := AtoiBase(nbr, baseFrom)
	finalConverted := printNbrBase(convertedNum, baseTo)
	return finalConverted
}

// result := student.ConvertBase("101011", "01", "0123456789")
// fmt.Println(result)
// result := student.ConvertBase("uuhoumo", "choumi", "Zone01")
// fmt.Println(result)