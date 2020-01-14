package main

import (
	"fmt"
	"os"

	student ".."
)

func main() {
	if os.Args[1] == "1" {
		root := &student.TreeNode{Data: "01"}
		student.BTreeInsertData(root, "07")
		student.BTreeInsertData(root, "05")
		student.BTreeInsertData(root, "12")
		student.BTreeInsertData(root, "02")
		student.BTreeInsertData(root, "10")
		student.BTreeInsertData(root, "03")
		student.BTreeApplyByLevel(root, fmt.Println)
	}
	if os.Args[1] == "2" {
		root := &student.TreeNode{Data: "01"}
		student.BTreeInsertData(root, "07")
		student.BTreeInsertData(root, "05")
		student.BTreeInsertData(root, "12")
		student.BTreeInsertData(root, "02")
		student.BTreeInsertData(root, "10")
		student.BTreeInsertData(root, "03")
		student.BTreeApplyByLevel(root, fmt.Print)
		fmt.Println()
	}

}
