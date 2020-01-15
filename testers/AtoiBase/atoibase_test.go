package student

import (
	"testing"
)

func TestAtoiBase(t *testing.T) {
	tables := []struct {
		nbr  string
		base string
		res  int
	}{
		{"125", "0123456789", 125},
		{"125", "01234567890", 0},
		{"1111101", "01", 125},
		{"125", "-01", 0},
		{"7D", "0123456789ABCDEF", 125},
		{"uoi", "choumi", 125},
		{"sfdf", "asdf", 123},
		{"125", "aa", 0},
		{"125", "a", 0},
		//from audit
		{"bcbbbbaab", "a", 0},
		{"0001", "01", 1},
		{"00", "01", 0},
		{"saDt!I!sI", "CHOUMIisDAcat!", 11557277891},
		{"AAho?Ao", "WhoAmI?", 406772},
		{"thisinputshouldnotmatter", "abca", 0},
		{"125", "0123456789", 125},
		{"uoi", "choumi", 125},
		{"bbbbbab", "-ab", 0},
	}
	for _, table := range tables {
		total := AtoiBase(table.nbr, table.base)
		if total != table.res {
			t.Errorf("AtoiBase of (%s, %s) was incorrect, got: %d, want: %d.", table.nbr, table.base, total, table.res)
		}
	}
}
