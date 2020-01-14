package main

import (
	"io/ioutil"
	"os"
	s ".."
	"github.com/01-edu/z01"
	"fmt"
)

func main() {
	arguments := os.Args
	length := 0
	for range arguments {
		length++
	}
	if length == 1 {
		return
	} else {
		// flagC := false
		// invalFlag := false
		length := 10
		contesx := ""
		nameErr := ""
		for index, item := range arguments {
			if item == "-c" {
				arguments = append(arguments[:index], arguments[index+1:]...)
				if len(arguments) > index {
					nameErr = arguments[index]
					length = s.Atoi(arguments[index])
					arguments = append(arguments[:index], arguments[index+1:]...)
				} else {
					contesx = "tail: option requires an argument -- 'c'\nTry 'tail --help' for more information."
				}
				break
			}
		}
		for index, item := range arguments {
				if index == 0{
					continue
				}
				if length == 0 {
					erro := "tail: invalid number of bytes: ‘" + nameErr+"’"
					for _, inde := range erro {
						z01.PrintRune(inde)
					}
					z01.PrintRune('\n')
					return
				}
				if contesx != "" {
					fmt.Print(contesx)
					z01.PrintRune('\n')
				}
				file1, err1 := os.Open(item)
				title := item
				b, err2 := ioutil.ReadAll(file1)
				if err1 != nil {
					for _, item := range string(err1.Error()) {
						z01.PrintRune(item)
					}
					continue
				}
				if err2 != nil {
					for _, item := range string(err2.Error()) {
						z01.PrintRune(item)
					}
					continue
				}
				from := len(string(b))-length
				if from < 0 {
					from = 0
				}
				str := "==> " + title + " <=="
				for _, it := range str {
					z01.PrintRune(it)
				}
				z01.PrintRune('\n')

				for i := from; i < len(string(b)) ; i++ {
					z01.PrintRune(rune(b[i]))
				}
			}


	}
}