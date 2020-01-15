package student

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestBTreeTransplant(t *testing.T) {
	root := &TreeNode{Data: "4"}
	BTreeInsertData(root, "1")
	BTreeInsertData(root, "7")
	BTreeInsertData(root, "5")
	node := BTreeSearchItem(root, "1")
	repl := &TreeNode{Data: "3"}
	// root = BTreeTransplant(root, node, repl)
	tables := []struct {
		root *TreeNode
		node *TreeNode
		repl *TreeNode
		res  string
	}{
		{root: root, node: node, repl: repl, res: "3\n4\n5\n7\n"},
	}

	for _, table := range tables {
		root := BTreeTransplant(table.root, table.node, table.repl)
		rescueStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w
		BTreeApplyInorder(root, fmt.Println)
		w.Close()
		out, _ := ioutil.ReadAll(r)
		os.Stdout = rescueStdout
		total := string(out)
		if total != table.res {
			t.Errorf("BTreeTransplant of (%v) was incorrect, got: %s, want: %s.", table.root, total, table.res)
		}
	}

}
