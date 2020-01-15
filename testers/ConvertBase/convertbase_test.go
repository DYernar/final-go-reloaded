package student

import (
	"testing"

)

func TestConvertBase(t *testing.T) {
	tables := []struct {
		nbr      string
		baseFrom string
		baseTo   string
		res      string
	}{
		{"101011", "01", "0123456789", "43"},
		{"31031310", "0123", "choumi", "hcimumc"},
		{"125", "0123456789", "choumi", "uoi"},
		{"125", "0123456789", "01", "1111101"},
	}
	for _, table := range tables {
		total := ConvertBase(table.nbr, table.baseFrom, table.baseTo)
		if total != table.res {
			t.Errorf("ConvertBase of (%s, %s, %s) was incorrect, got: %s, want: %s.", table.nbr, table.baseFrom, table.baseTo, total, table.res)
		}
	}
}
