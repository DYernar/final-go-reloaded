package student
import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

)

func TestBTreeApplyByLevel(t *testing.T) {
	root1 := &TreeNode{Data: "4"}
	BTreeInsertData(root1, "1")
	BTreeInsertData(root1, "7")
	BTreeInsertData(root1, "5")
	tables := []struct {
		root *TreeNode
		f    func(...interface{})
		res  string
	}{
		{root: root1, res: "4\n1\n7\n5\n"},
	}

	for _, table := range tables {
		rescueStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w
		BTreeApplyByLevel(table.root, fmt.Println)
		w.Close()
		out, _ := ioutil.ReadAll(r)
		os.Stdout = rescueStdout
		total := string(out)
		if total != table.res {
			t.Errorf("BTreeApplyByLevel of (%v) was incorrect, got: %s, want: %s.", table.root, total, table.res)
		}
	}

}
