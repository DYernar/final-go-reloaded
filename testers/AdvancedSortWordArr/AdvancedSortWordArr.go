package student


func AdvancedSortWordArr(array []string, f func(a, b string) int) {
	for i := 0; i < len(array)-1; i++ {
		for j := i; j < len(array); j++ {
			if f(array[i], array[j]) == 1 {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
}

// result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
// arrays := []string{"The", "earliest", "computing", "device", "undoubtedly", "consisted", "of", "the", "five", "fingers", "of", "each", "hand"}
// student.AdvancedSortWordArr(result, student.Compare)
// student.AdvancedSortWordArr(arrays, student.Compare)

// fmt.Println(result)
// fmt.Println(arrays)