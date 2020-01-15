package student

import (
	"testing"

)

func TestActiveBits(task *testing.T) {
	tables := []struct {
		input  int
		output uint
	}{
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 1},
		{5, 2},
		{6, 2},
		{7, 3},
		{8, 1},
		{9, 2},
		{10, 2},
	}
	for _, table := range tables {
		total := ActiveBits(table.input)
		if total != table.output {
			task.Errorf("ActiveBits of (%d) was incorrect, got: %d, want: %d.", table.input, total, table.output)
		}
	}
}
