package student

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func PrintList2(l *List) {
	it := l.Head
	for it != nil {
		fmt.Print(it.Data, " -> ")
		it = it.Next
	}

	fmt.Print(nil, "\n")
}

func TestListRemoveIf(t *testing.T) {
	link1 := &List{}
	// fmt.Println("----normal state----")
	ListPushBack(link1, 1)
	ListPushBack(link1, "Hello")
	ListPushBack(link1, 1)
	ListPushBack(link1, "There")
	ListPushBack(link1, 1)
	ListPushBack(link1, 1)
	ListPushBack(link1, "How")
	ListPushBack(link1, 1)
	ListPushBack(link1, "are")
	ListPushBack(link1, "you")
	ListPushBack(link1, 1)

	tables := []struct {
		list     **List
		data_ref interface{}
		res      string
	}{
		{list: &link1, data_ref: 1, res: "Hello -> There -> How -> are -> you -> <nil>\n"},
	}

	for _, table := range tables {
		ListRemoveIf(*table.list, table.data_ref)
		rescueStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		PrintList2(*table.list)

		w.Close()
		out, _ := ioutil.ReadAll(r)
		os.Stdout = rescueStdout
		total := string(out)
		if total != table.res {
			t.Errorf("ListRemoveIf of (%v, %d) was incorrect, got: %s, want: %s.", table.list, table.data_ref, total, table.res)
		}
	}

}
