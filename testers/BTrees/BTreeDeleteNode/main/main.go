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
		node := student.BTreeSearchItem(root, "02")
		fmt.Println("Before delete:")
		student.BTreeApplyInorder(root, fmt.Println)
		root = student.BTreeDeleteNode(root, node)
		fmt.Println("\nTo delete: 02\n")
		fmt.Println("After delete:")
		student.BTreeApplyInorder(root, fmt.Println)
	}
	if os.Args[1] == "2" {
		root := &student.TreeNode{Data: "33"}
		student.BTreeInsertData(root, "20")
		student.BTreeInsertData(root, "5")
		student.BTreeInsertData(root, "13")
		student.BTreeInsertData(root, "31")
		student.BTreeInsertData(root, "52")
		student.BTreeInsertData(root, "11")
		node := student.BTreeSearchItem(root, "20")
		fmt.Println("Before delete:")
		student.BTreeApplyInorder(root, fmt.Println)
		root = student.BTreeDeleteNode(root, node)
		fmt.Println("\nTo delete: 20\n")
		fmt.Println("After delete:")
		student.BTreeApplyInorder(root, fmt.Println)
	}
	if os.Args[1] == "3" {
		root := &student.TreeNode{Data: "03"}
		student.BTreeInsertData(root, "39")
		student.BTreeInsertData(root, "11")
		student.BTreeInsertData(root, "99")
		student.BTreeInsertData(root, "14")
		student.BTreeInsertData(root, "44")
		student.BTreeInsertData(root, "11")
		node := student.BTreeSearchItem(root, "03")
		fmt.Println("Before delete:")
		student.BTreeApplyInorder(root, fmt.Println)
		root = student.BTreeDeleteNode(root, node)
		fmt.Println("\nTo delete: 03\n")
		fmt.Println("After delete:")
		student.BTreeApplyInorder(root, fmt.Println)
	}
	if os.Args[1] == "4" {
		root := &student.TreeNode{Data: "03"}
		student.BTreeInsertData(root, "01")
		student.BTreeInsertData(root, "03")
		student.BTreeInsertData(root, "94")
		student.BTreeInsertData(root, "19")
		student.BTreeInsertData(root, "111")
		student.BTreeInsertData(root, "24")
		node := student.BTreeSearchItem(root, "03")
		fmt.Println("Before delete:")
		student.BTreeApplyInorder(root, fmt.Println)
		root = student.BTreeDeleteNode(root, node)
		fmt.Println("\nTo delete: 03\n")
		fmt.Println("After delete:")
		student.BTreeApplyInorder(root, fmt.Println)
	}
}
