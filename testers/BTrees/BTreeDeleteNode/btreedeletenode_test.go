package student

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestBTreeDeleteNode(t *testing.T) {
	root1 := &TreeNode{Data: "100"}
	BTreeInsertData(root1, "90")
	BTreeInsertData(root1, "150")
	BTreeInsertData(root1, "140")
	BTreeInsertData(root1, "170")
	BTreeInsertData(root1, "70")
	BTreeInsertData(root1, "95")
	nodeDelete1 := BTreeSearchItem(root1, "100")
	nodeDelete2 := BTreeSearchItem(root1, "150")

	tables := []struct {
		root *TreeNode
		node *TreeNode
		res  string
	}{
		{root: root1, node: nodeDelete1, res: "140\n150\n170\n70\n90\n95\n"},
		{root: root1, node: nodeDelete2, res: "100\n140\n170\n70\n90\n95\n"},
	}

	for _, table := range tables {
		root := BTreeDeleteNode(table.root, table.node)
		rescueStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w
		BTreeApplyInorder(root, fmt.Println)
		w.Close()
		out, _ := ioutil.ReadAll(r)
		os.Stdout = rescueStdout
		total := string(out)
		if total != table.res {
			t.Errorf("BTreeDeleteNode of (%v, %v) was incorrect, got: %s, want: %s.", table.root, table.node, total, table.res)
		}
	}

}
