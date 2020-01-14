package main

import (
	"io"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	argsCount := 0
	for range os.Args {
		argsCount++
	}
	if argsCount == 1 {
		_, err := io.Copy(os.Stdout, os.Stdin)
		if err != nil {
			str := err.Error()
			slice := []rune(str)
			for _, char := range slice {
				z01.PrintRune(char)
			}
			z01.PrintRune('\n')
			return
		}
	}
	if argsCount > 1 {
		for i := 1; i < argsCount; i++ {
			filename := os.Args[i]
			file, err := os.Open(filename)
			if err != nil {
				str := "cat: " + filename + ": No such file or directory"
				slice := []rune(str)
				for _, char := range slice {
					z01.PrintRune(char)
				}
				z01.PrintRune('\n')
				break
			} else {
				io.Copy(os.Stdout, file)
				z01.PrintRune('\n')
			}
			file.Close()
		}
	}
}