package main

import (
	"fmt"
	"os"
	student ".."
)

func main() {
	if os.Args[1] == "1" {
		test1()
	}
	if os.Args[1] == "2" {
		test2()
	}
	if os.Args[1] == "3" {
		test3()
	}
}

func test1() {
	root := &student.TreeNode{Data: "01"}
	student.BTreeInsertData(root, "07")
	student.BTreeInsertData(root, "05")
	student.BTreeInsertData(root, "12")
	student.BTreeInsertData(root, "02")
	student.BTreeInsertData(root, "10")
	student.BTreeInsertData(root, "03")
	node := student.BTreeSearchItem(root, "12")
	replacement := &student.TreeNode{Data: "55"}
	student.BTreeInsertData(replacement, "33")
	student.BTreeInsertData(replacement, "60")
	student.BTreeInsertData(replacement, "12")
	student.BTreeInsertData(replacement, "15")
	root = student.BTreeTransplant(root, node, replacement)
	student.BTreeApplyInorder(root, fmt.Println)
}

func test2() {
	root := &student.TreeNode{Data: "33"}
	student.BTreeInsertData(root, "20")
	student.BTreeInsertData(root, "5")
	student.BTreeInsertData(root, "13")
	student.BTreeInsertData(root, "31")
	student.BTreeInsertData(root, "52")
	student.BTreeInsertData(root, "11")
	node := student.BTreeSearchItem(root, "20")
	replacement := &student.TreeNode{Data: "55"}
	student.BTreeInsertData(replacement, "33")
	student.BTreeInsertData(replacement, "60")
	student.BTreeInsertData(replacement, "12")
	student.BTreeInsertData(replacement, "15")
	root = student.BTreeTransplant(root, node, replacement)
	student.BTreeApplyInorder(root, fmt.Println)
}

func test3() {
	root := &student.TreeNode{Data: "03"}
	student.BTreeInsertData(root, "39")
	student.BTreeInsertData(root, "11")
	student.BTreeInsertData(root, "99")
	student.BTreeInsertData(root, "14")
	student.BTreeInsertData(root, "44")
	student.BTreeInsertData(root, "11")
	node := student.BTreeSearchItem(root, "11")
	replacement := &student.TreeNode{Data: "55"}
	student.BTreeInsertData(replacement, "33")
	student.BTreeInsertData(replacement, "60")
	student.BTreeInsertData(replacement, "12")
	student.BTreeInsertData(replacement, "15")
	root = student.BTreeTransplant(root, node, replacement)
	student.BTreeApplyInorder(root, fmt.Println)
}
