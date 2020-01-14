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

// package main

// import (
// 	"fmt"
// 	student ".."
// )

// func main() {
// 	str := "HelloHAhowHAareHAyou?"
// 	str1 := "|=choumi=|which|=choumi=|itself|=choumi=|used|=choumi=|cards|=choumi=|and|=choumi=|a|=choumi=|central|=choumi=|computing|=choumi=|unit.|=choumi=|When|=choumi=|the|=choumi=|machine|=choumi=|was|=choumi=|finished,"
// 	str2 := "!==!which!==!was!==!making!==!all!==!kinds!==!of!==!punched!==!card!==!equipment!==!and!==!was!==!also!==!in!==!the!==!calculator!==!business[10]!==!to!==!develop!==!his!==!giant!==!programmable!==!calculator,"
// 	str3 := "AFJCharlesAFJBabbageAFJstartedAFJtheAFJdesignAFJofAFJtheAFJfirstAFJautomaticAFJmechanicalAFJcalculator,"
// 	str4 := "<<==123==>>In<<==123==>>1820,<<==123==>>Thomas<<==123==>>de<<==123==>>Colmar<<==123==>>launched<<==123==>>the<<==123==>>mechanical<<==123==>>calculator<<==123==>>industry[note<<==123==>>1]<<==123==>>when<<==123==>>he<<==123==>>released<<==123==>>his<<==123==>>simplified<<==123==>>arithmometer,"
// 	fmt.Println(student.Split(str, "HA"))
// 	fmt.Println(student.Split(str1, "|=choumi=|"))
// 	fmt.Println(student.Split(str2, "!==!"))
// 	fmt.Println(student.Split(str3, "AFJ"))
// 	fmt.Println(student.Split(str4, "<<==123==>>"))
// }
