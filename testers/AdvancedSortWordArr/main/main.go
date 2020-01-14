package main

import (
	"fmt"

	student ".."
)

func main() {

	result1 := []string{"The", "earliest", "computing", "device", "undoubtedly", "consisted", "of", "the", "five", "fingers", "of", "each", "hand"}
	result2 := []string{"The", "word", "digital", "comesfrom", "\"digits\"", "or", "fingers"}
	result3 := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}

	student.AdvancedSortWordArr(result1, student.Compare)
	student.AdvancedSortWordArr(result2, student.Compare)
	student.AdvancedSortWordArr(result3, student.Compare)

	fmt.Println("Expected result: [The computing consisted device each earliest fingers five hand of of the undoubtedly]")
	fmt.Println("Your result    :", result1)
	fmt.Println()
	fmt.Println("Expected result: [\"digits\" The comesfrom digital fingers or word]")
	fmt.Println("Your result    :", result2)
	fmt.Println()
	fmt.Println("Expected result: [1 2 3 A B C a b c]")
	fmt.Println("Your result    :", result3)
	fmt.Println()
}
