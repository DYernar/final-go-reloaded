package student

import (
	"testing"
)

func TestAdvancedSortWordArr(t *testing.T) {
	tables := []struct {
		array []string
		cons  []string
		res   []string
	}{
		{[]string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}, []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}, []string{"1", "2", "3", "A", "B", "C", "a", "b", "c"}},
		{[]string{"aa", "A2", "A1", "ab", "b", "B", "2", "c", "C", "3"}, []string{"aa", "A2", "A1", "ab", "b", "B", "2", "c", "C", "3"}, []string{"2", "3", "A1", "A2", "B", "C", "aa", "ab", "b", "c"}},
	}
	for _, table := range tables {
		AdvancedSortWordArr(table.array, Compare)
		for i, tt := range table.array {
			if tt != table.res[i] {
				t.Errorf("AdvancedSortWordArr of (%v) was incorrect, got: %v, want: %v.", table.cons, table.array, table.res)
				break
			}
		}
	}
}
