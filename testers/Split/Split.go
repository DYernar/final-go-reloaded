package student

func Split(str, charset string) []string {
	var arr []string
	var indices []int
	indices = append(indices, -1)
	for index, item := range str {
		if item == rune(charset[0]) {
			newStr := ""
			for i := index; i < index+len(charset); i++ {
				newStr = newStr + string(str[i]) 
			}
			if newStr == charset {
				for k := index; k < index+len(charset); k++ {
					indices = append(indices, k)
				}
			}
		}
	}
	indices = append(indices, len(str))
	for i := 0; i < len(indices)-1; i++ {
		tempStr := ""
		for j := indices[i]+1; j < indices[i+1]; j++ {
			tempStr = tempStr+string(str[j])
		}
		if i != 0 {
			if tempStr != "" {
				arr = append(arr, tempStr)
			}
		} else {
			arr = append(arr, tempStr)
			
		}

	}
	return arr
}

